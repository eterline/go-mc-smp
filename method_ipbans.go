package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:ip_bans
//
// Path      | Description                      | Parameters            | Result
// ----------|----------------------------------|-----------------------|----------------------
// /         | Get the ip ban list              | None                  | banlist: []IPBan
// /set      | Set the ip ban list              | bans: []IPBan         | banlist: []IPBan
// /add      | Add players to the ip ban list   | add:  []IPBan         | banlist: []IPBan
// /remove   | Remove players from ip ban list  | remove: []Player      | banlist: []IPBan
// /clear    | Clear all players in ip ban list | None                  | banlist: []IPBan

// IPBansGet - Get the ip ban list
func (rpc *RPCClient) IPBansGet(ctx context.Context) ([]IPBan, error) {
	method := usage.NewMethod("ip_bans").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[[]IPBan](r)
	if err != nil {
		return nil, err
	}

	return *data, nil
}

// IPBansSet - Set the ip ban list
func (rpc *RPCClient) IPBansSet(ctx context.Context, ban ...IPBan) error {
	method := usage.NewMethod("ip_bans").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, ban)
	if err != nil {
		return err
	}

	return r.Err()
}

// IPBansAdd - Add players to the ip ban list
func (rpc *RPCClient) IPBansAdd(ctx context.Context, ban ...IPBan) error {
	method := usage.NewMethod("ip_bans").Add("add").String()
	r, err := rpc.core.CallWithContext(ctx, method, ban)
	if err != nil {
		return err
	}

	return r.Err()
}

// IPBansRemove - Remove players from ip ban list
func (rpc *RPCClient) IPBansRemove(ctx context.Context, player ...IPBan) error {
	method := usage.NewMethod("ip_bans").Add("remove").String()
	r, err := rpc.core.CallWithContext(ctx, method, player)
	if err != nil {
		return err
	}

	return r.Err()
}

// IPBansClear - Clear all players in ip ban list
func (rpc *RPCClient) IPBansClear(ctx context.Context) error {
	method := usage.NewMethod("ip_bans").Add("clear").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return err
	}

	return r.Err()
}
