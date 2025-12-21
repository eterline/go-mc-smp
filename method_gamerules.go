package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:gamerules
//
// Path     | Description                                               | Parameters                | Result
// ---------|-----------------------------------------------------------|---------------------------|----------------------------
// /        | Get the available game rule keys and their current values | None                      | gamerules: []TypedGameRule
// /update  | Update game rule value                                    | gamerule: UntypedGameRule | gamerule: TypedGameRule

// GamerulesGet - Get the available game rule keys and their current values
func (rpc *RPCClient) GamerulesGet(ctx context.Context) ([]GameRule, error) {
	method := usage.NewMethod("gamerules").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[[]GameRule](r)
	if err != nil {
		return nil, err
	}

	return *data, nil
}

// GamerulesUpdate - Update game rule value
func (rpc *RPCClient) GamerulesUpdate(ctx context.Context, rule GameRule) (*GameRule, error) {
	rule.Type = UntypedGameRule // for correct api usage

	method := usage.NewMethod("gamerules").Add("update").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{rule})
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[GameRule](r)
	if err != nil {
		return nil, err
	}

	return data, nil
}
