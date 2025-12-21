package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

func (rpc *RPCClient) ServerStatus(ctx context.Context) (*ServerState, error) {
	method := usage.NewMethod("server").Add("status").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[ServerState](r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (rpc *RPCClient) ServerSave(ctx context.Context, flush bool) (bool, error) {
	method := usage.NewMethod("server").Add("save").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{flush})
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

func (rpc *RPCClient) ServerStop(ctx context.Context) (bool, error) {
	method := usage.NewMethod("server").Add("stop").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

func (rpc *RPCClient) ServerSystemMessage(ctx context.Context, message SystemMessage) (bool, error) {
	method := usage.NewMethod("server").Add("system_message").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{message})
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}
