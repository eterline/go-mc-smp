package gomcsmp

import (
	"context"
	"fmt"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

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

func (rpc *RPCClient) AllowlistAdd(ctx context.Context, p ...Player) error {
	method := usage.NewMethod("allowlist").Add("add").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

func (rpc *RPCClient) AllowlistRemove(ctx context.Context, p ...Player) error {
	method := usage.NewMethod("allowlist").Add("remove").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

func (rpc *RPCClient) AllowlistClear(ctx context.Context) error {
	method := usage.NewMethod("allowlist").Add("clear").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return err
	}

	return r.Err()
}

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
