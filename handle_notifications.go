package gomcsmp

import (
	"context"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

func notifyStream[T any](
	ctx context.Context,
	notify *notificationPipe,
	method string,
	okOnly bool,
) <-chan T {
	out := make(chan T, 1)

	in := notify.Register(method)

	go func() {
		defer close(out)
		defer notify.Unregister(method)

		var zero T

		for {
			select {
			case <-ctx.Done():
				return

			case n, ok := <-in:
				if !ok {
					return
				}

				if okOnly {
					if n.Err() != nil {
						continue
					}

					out <- zero
					continue
				}

				data, err := jsonrpc.DecodeRPCParams[T](n)
				if err != nil {
					continue
				}

				out <- *data
			}
		}
	}()

	return out
}

// ===========

func (rpc *RPCClient) NotifyPlayersJoined(ctx context.Context) <-chan Player {
	method := usage.NewMethod("notification").
		Add("players").
		Add("joined").
		String()

	return notifyStream[Player](ctx, rpc.notify, method, false)
}

func (rpc *RPCClient) NotifyPlayersLeft(ctx context.Context) <-chan Player {
	method := usage.NewMethod("notification").
		Add("players").
		Add("left").
		String()

	return notifyStream[Player](ctx, rpc.notify, method, false)
}

// ===========

func (rpc *RPCClient) NotifyServerStarted(ctx context.Context) <-chan struct{} {
	method := usage.NewMethod("notification").
		Add("server").
		Add("started").
		String()

	return notifyStream[struct{}](ctx, rpc.notify, method, true)
}

func (rpc *RPCClient) NotifyServerStopping(ctx context.Context) <-chan struct{} {
	method := usage.NewMethod("notification").
		Add("server").
		Add("stopping").
		String()

	return notifyStream[struct{}](ctx, rpc.notify, method, true)
}

func (rpc *RPCClient) NotifyServerSaving(ctx context.Context) <-chan struct{} {
	method := usage.NewMethod("notification").
		Add("server").
		Add("saving").
		String()

	return notifyStream[struct{}](ctx, rpc.notify, method, true)
}

func (rpc *RPCClient) NotifyServerSaved(ctx context.Context) <-chan struct{} {
	method := usage.NewMethod("notification").
		Add("server").
		Add("saved").
		String()

	return notifyStream[struct{}](ctx, rpc.notify, method, true)
}

func (rpc *RPCClient) NotifyServerStatus(ctx context.Context) <-chan ServerState {
	method := usage.NewMethod("notification").
		Add("server").
		Add("status").
		String()

	return notifyStream[ServerState](ctx, rpc.notify, method, false)
}

// ===========

func (rpc *RPCClient) NotifyGamerulesUpdates(ctx context.Context) <-chan GameRule {
	method := usage.NewMethod("notification").
		Add("gamerules").
		Add("updated").
		String()

	return notifyStream[GameRule](ctx, rpc.notify, method, false)
}
