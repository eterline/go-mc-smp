package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:server
//
// Path              | Description           | Parameters               | Result
// ------------------|-----------------------|--------------------------|------------------------
// /status           | Get server status     | None                     | status: ServerState
// /save             | Save server state     | flush: bool              | saving: bool
// /stop             | Stop server           | None                     | stopping: bool
// /system_message   | Send a system message | message: SystemMessage   | sent: bool

// ServerStatus - Get server status
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

// ServerSave - Save server state
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

// ServerStop - Stop server
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

// ServerSystemMessage - Send a system message
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
