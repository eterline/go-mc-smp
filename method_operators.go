package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:operators
//
// Path      | Description            | Parameters            | Result
// ----------|------------------------|-----------------------|----------------------
// /         | Get all oped players   | None                  | operators: []Operator
// /set      | Set all oped players   | operators: []Operator | operators: []Operator
// /add      | Op players             | add: []Operator       | operators: []Operator
// /remove   | Deop players           | remove: []Player      | operators: []Operator
// /clear    | Deop all players       | None                  | operators: []Operator

// OperatorsGet - Get all oped players
func (rpc *RPCClient) OperatorsGet(ctx context.Context) ([]Operator, error) {
	method := usage.NewMethod("operators").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonrpc.DecodeRPCResult[[]Operator](r)
	if err != nil {
		return nil, err
	}

	return *data, nil
}

// OperatorsSet - Set all oped players
func (rpc *RPCClient) OperatorsSet(ctx context.Context, p ...Operator) error {
	method := usage.NewMethod("operators").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

// OperatorsAdd - Op players
func (rpc *RPCClient) OperatorsAdd(ctx context.Context, p ...Operator) error {
	method := usage.NewMethod("operators").Add("add").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

// OperatorsRemove - Deop players
func (rpc *RPCClient) OperatorsRemove(ctx context.Context, p ...Player) error {
	method := usage.NewMethod("operators").Add("remove").String()
	r, err := rpc.core.CallWithContext(ctx, method, []any{p})
	if err != nil {
		return err
	}

	return r.Err()
}

// OperatorsClear - Deop all players
func (rpc *RPCClient) OperatorsClear(ctx context.Context) error {
	method := usage.NewMethod("operators").Add("clear").String()
	r, err := rpc.core.CallWithContext(ctx, method, nil)
	if err != nil {
		return err
	}

	return r.Err()
}
