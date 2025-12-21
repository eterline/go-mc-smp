package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Logger interface {
	Info(v ...any)
	Error(v ...any)
}

// ==========

type JsonRPCClient struct {
	conn *websocket.Conn

	reqID      int
	reqMutex   sync.Mutex
	requests   chan *RPCRequest
	reqTimeout time.Duration

	responses     map[int]chan *RPCResponse
	resMutex      sync.Mutex
	notifications chan *RPCResponse

	logger Logger
}

func NewJsonRPCClient(url, token string, callTimeout time.Duration) (*JsonRPCClient, error) {
	return NewJsonRPCClientWithContext(context.Background(), url, token, callTimeout)
}

func NewJsonRPCClientWithContext(ctx context.Context, url, token string, callTimeout time.Duration) (*JsonRPCClient, error) {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, url, authHeader(token))
	if err != nil {
		return nil, err
	}

	client := &JsonRPCClient{
		conn:          conn,
		requests:      make(chan *RPCRequest, 100),
		responses:     make(map[int]chan *RPCResponse),
		notifications: make(chan *RPCResponse, 100),
		reqTimeout:    callTimeout,
	}

	go client.writer()
	go client.reader()

	return client, nil
}

func (c *JsonRPCClient) Logger(l Logger) {
	c.logger = l
}

func (c *JsonRPCClient) nextID() int {
	c.reqMutex.Lock()
	defer c.reqMutex.Unlock()
	c.reqID++
	return c.reqID
}

func (c *JsonRPCClient) writer() {
	for req := range c.requests {
		if err := c.conn.WriteJSON(req); err != nil {
			c.Log().Error("write request error:", err)
		}
	}
}

func (c *JsonRPCClient) reader() {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			c.Log().Error("read response error:", err)
			return
		}

		var resp RPCResponse

		if err := json.Unmarshal(msg, &resp); err != nil {
			c.Log().Error("read response error:", err)
			return
		}

		if resp.ID != 0 {
			c.resMutex.Lock()
			if ch, ok := c.responses[resp.ID]; ok {
				ch <- &resp
				close(ch)
			}
			c.resMutex.Unlock()
			continue
		}

		select {
		case c.notifications <- &resp:
		default:
			c.Log().Info("notification channel full: dropping message")
		}
	}
}

func (c *JsonRPCClient) Call(method string, params []any) (*RPCResponse, error) {
	return c.CallWithContext(context.Background(), method, params)
}

func (c *JsonRPCClient) CallWithContext(ctx context.Context, method string, params []any) (*RPCResponse, error) {

	id := c.nextID()

	req, err := NewRPCRequest(id, method, params)
	if err != nil {
		return nil, fmt.Errorf("rpc request encode error: %w", err)
	}

	respCh := make(chan *RPCResponse, 1)

	c.resMutex.Lock()
	c.responses[id] = respCh
	c.resMutex.Unlock()

	defer func() {
		c.resMutex.Lock()
		delete(c.responses, id)
		c.resMutex.Unlock()
	}()

	select {
	case c.requests <- req:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, c.reqTimeout)
	defer cancel()

	select {
	case resp, ok := <-respCh:
		if !ok {
			return nil, fmt.Errorf("rpc response channel closed")
		}
		return resp, nil

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *JsonRPCClient) Notifications() <-chan *RPCResponse {
	return c.notifications
}

func (c *JsonRPCClient) Close() error {
	close(c.requests)
	return c.conn.Close()
}

// ====

func authHeader(token string) http.Header {
	h := http.Header{}
	h.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return h
}
