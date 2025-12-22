package gomcsmp

import (
	"context"
	"fmt"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:players
//
// Path    | Description               | Parameters                | Result
// --------|---------------------------|---------------------------|------------------------
// /       | Get all connected players | None                      | players: []Player
// /kick   | Kick players              | kick: []KickPlayer        | kicked: []Player

// PlayersGet - Get all connected players
func (rpc *RPCClient) PlayersGet(ctx context.Context) (*PlayerRegistry, error) {
	method := usage.NewMethod("players").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[PlayerRegistry](r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PlayersKick - Kick players
func (rpc *RPCClient) PlayersKick(ctx context.Context, kicks ...KickPlayer) (*PlayerRegistry, error) {

	/*
		NOTE: When kicking a player via RPC, an "Already retired" error may occur.
			* Reason:
			* 1. The player may already be in the process of disconnecting (e.g., left the server voluntarily).
			* 2. At the same time, the server attempts to remove the player from the world via PlayerList.remove.
			* 3. If the player's object is already "retired" in the EntityScheduler, an IllegalStateException is thrown.
			*
			* Why it is implemented this way:
			* - The RPC kick method does not check the player's state to avoid blocking the call and to prevent race conditions.
			* - The "Already retired" exception is safe to ignore because the goal of the RPC has been achieved: the player is already disconnected or kicked.
			* - After the RPC call, the online players list can be rechecked to confirm that the intended players have been removed.
			*
			* This pattern ensures:
			* - Idempotency of the kick call (calling it multiple times is safe),
			* - Correct behavior even with asynchronous player disconnections,
			* - No server crash when attempting to kick a player who has already disconnected.
	*/

	players, err := rpc.PlayersGet(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get players: %w", err)
	}

	toKick := make([]KickPlayer, 0, len(kicks))
	for _, k := range kicks {
		if players.IsOnline(k.Player.Name) {
			toKick = append(toKick, k)
		}
	}

	if len(toKick) == 0 {
		return players, nil
	}

	method := usage.NewMethod("players").Add("kick").String()
	_, _ = rpc.core.CallWithContext(ctx, method, []any{toKick})

	updatedPlayers, err := rpc.PlayersGet(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated players: %w", err)
	}

	return updatedPlayers, nil
}
