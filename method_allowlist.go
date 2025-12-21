package gomcsmp

import (
	"context"
	"fmt"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:allowlist
//
// Path      | Description                                       | Parameters        | Result
// ----------|---------------------------------------------------|-------------------|------------------------
// /         | Get the allowlist                                 | None              | allowlist: []Player
// /set      | Set the allowlist to the provided list of players | players: []Player | allowlist: []Player
// /add      | Add players to the allowlist                      | add: []Player     | allowlist: []Player
// /remove   | Remove players from the allowlist                 | remove: []Player  | allowlist: []Player
// /clear    | Clear all players in the allowlist                | None              | allowlist: []Player

// AllowlistGet - Get the allowlist
func (rpc *RPCClient) AllowlistGet(ctx context.Context) ([]Player, error) {
	method := usage.NewMethod("allowlist").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[[]Player](r)
	if err != nil {
		return nil, err
	}

	return *data, nil
}

// AllowlistAdd - Add players to the allowlist
func (rpc *RPCClient) AllowlistAdd(ctx context.Context, p ...Player) error {
	method := usage.NewMethod("allowlist").Add("add").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

// AllowlistRemove - Remove players from allowlist
func (rpc *RPCClient) AllowlistRemove(ctx context.Context, p ...Player) error {
	method := usage.NewMethod("allowlist").Add("remove").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

// AllowlistClear - Clear all players in allowlist
func (rpc *RPCClient) AllowlistClear(ctx context.Context) error {
	method := usage.NewMethod("allowlist").Add("clear").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return err
	}

	return r.Err()
}

// AllowlistSet - Set the allowlist to the provided list of players
func (rpc *RPCClient) AllowlistSet(ctx context.Context, p ...Player) error {
	for _, pl := range p {
		if pl.ID == nil {
			return fmt.Errorf("player '%s' must have certain UUID", pl.Name)
		}
	}

	method := usage.NewMethod("allowlist").Add("add").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}
