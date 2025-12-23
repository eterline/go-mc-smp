package gomcsmp

import (
	"sync"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

type notificationPipe struct {
	mu     sync.RWMutex
	pipe   map[string]chan *jsonrpc.RPCResponse
	closed bool
}

func newNotificationPipe() *notificationPipe {
	return &notificationPipe{
		pipe: make(map[string]chan *jsonrpc.RPCResponse),
	}
}

func (np *notificationPipe) Close() {
	np.mu.Lock()
	defer np.mu.Unlock()

	if np.closed {
		return
	}

	for _, ch := range np.pipe {
		close(ch)
	}
	np.pipe = make(map[string]chan *jsonrpc.RPCResponse)
	np.closed = true
}

func (np *notificationPipe) Register(method string) <-chan *jsonrpc.RPCResponse {
	np.mu.Lock()
	defer np.mu.Unlock()

	if _, ok := np.pipe[method]; ok {
		usage.Panicf("notification handler for '%s' already registered", method)
	}

	ch := make(chan *jsonrpc.RPCResponse, 8)
	np.pipe[method] = ch
	return ch
}

func (np *notificationPipe) Unregister(method string) {
	np.mu.Lock()
	defer np.mu.Unlock()

	if ch, ok := np.pipe[method]; ok {
		close(ch)
		delete(np.pipe, method)
	}
}

func (np *notificationPipe) Push(method string, res *jsonrpc.RPCResponse) {
	np.mu.RLock()

	if np.closed {
		np.mu.RUnlock()
		return
	}

	ch, ok := np.pipe[method]
	np.mu.RUnlock()

	if !ok {
		return
	}

	go func() {
		ch <- res
	}()
}
