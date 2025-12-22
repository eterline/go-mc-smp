package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:bans
//
// Path      | Description                      | Parameters            | Result
// ----------|----------------------------------|-----------------------|----------------------
// /         | Get the ban list                 | None                  | banlist: []UserBan
// /set      | Set the banlist                  | bans: []UserBan       | banlist: []UserBan
// /add      | Add players to the ban list      | add:  []UserBan       | banlist: []UserBan
// /remove   | Remove players from ban list     | remove: []Player      | banlist: []UserBan
// /clear    | Clear all players in ban list    | None                  | banlist: []UserBan

// BansGet - Get the ban list
func (rpc *RPCClient) BansGet(ctx context.Context) ([]UserBan, error) {
	method := usage.NewMethod("bans").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[[]UserBan](r)
	if err != nil {
		return nil, err
	}

	return *data, nil
}

// BansSet - Set the banlist
func (rpc *RPCClient) BansSet(ctx context.Context, ban ...UserBan) error {
	method := usage.NewMethod("bans").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, ban)
	if err != nil {
		return err
	}

	return r.Err()
}

// BansAdd - Add players to the ban list
func (rpc *RPCClient) BansAdd(ctx context.Context, ban ...UserBan) error {
	method := usage.NewMethod("bans").Add("add").String()
	r, err := rpc.core.CallWithContext(ctx, method, ban)
	if err != nil {
		return err
	}

	return r.Err()
}

// BansRemove - Remove players from ban list
func (rpc *RPCClient) BansRemove(ctx context.Context, player ...Player) error {
	method := usage.NewMethod("bans").Add("remove").String()
	r, err := rpc.core.CallWithContext(ctx, method, player)
	if err != nil {
		return err
	}

	return r.Err()
}

// BansClear - Clear all players in ban list
func (rpc *RPCClient) BansClear(ctx context.Context) error {
	method := usage.NewMethod("bans").Add("clear").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return err
	}

	return r.Err()
}
