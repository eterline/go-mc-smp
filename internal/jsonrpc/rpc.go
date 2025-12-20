package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

type JsonRPCClient struct {
	conn       *websocket.Conn
	reqID      int
	reqMutex   sync.Mutex
	requests   chan *RPCRequest
	responses  map[int]chan *RPCResponse
	resMutex   sync.Mutex
	reqTimeout *atomic.Pointer[time.Duration]
}

func NewJsonRPCClient(url, token string) (*JsonRPCClient, error) {
	return NewJsonRPCClientWithContext(context.Background(), url, token)
}

func NewJsonRPCClientWithContext(ctx context.Context, url, token string) (*JsonRPCClient, error) {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, url, authHeader(token))
	if err != nil {
		return nil, err
	}

	callTm := &atomic.Pointer[time.Duration]{}
	d := 5 * time.Second
	callTm.Store(&d)

	client := &JsonRPCClient{
		conn:       conn,
		requests:   make(chan *RPCRequest, 100),
		responses:  make(map[int]chan *RPCResponse),
		reqTimeout: callTm,
	}

	go client.writer()
	go client.reader()

	return client, nil
}

func (c *JsonRPCClient) CallTimeout(d time.Duration) {
	c.reqTimeout.Store(&d)
}

func (c *JsonRPCClient) nextID() int {
	c.reqMutex.Lock()
	defer c.reqMutex.Unlock()
	c.reqID++
	return c.reqID
}

func (c *JsonRPCClient) writer() {
	for req := range c.requests {
		data, _ := json.Marshal(req)
		c.conn.WriteMessage(websocket.TextMessage, data)
	}
}

func (c *JsonRPCClient) reader() {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		var resp RPCResponse
		if err := json.Unmarshal(msg, &resp); err != nil {
			log.Println("unmarshal error:", err)
			continue
		}

		c.resMutex.Lock()
		if ch, ok := c.responses[resp.ID]; ok {
			ch <- &resp
			close(ch)
			delete(c.responses, resp.ID)
		}
		c.resMutex.Unlock()
	}
}

func (c *JsonRPCClient) Call(method string, params interface{}) (*RPCResponse, error) {
	id := c.nextID()
	req := &RPCRequest{
		ID:     id,
		Method: method,
		Params: params,
	}

	respCh := make(chan *RPCResponse, 1)
	c.resMutex.Lock()
	c.responses[id] = respCh
	c.resMutex.Unlock()

	c.requests <- req

	select {
	case resp := <-respCh:
		return resp, nil
	case <-time.After(*c.reqTimeout.Load()):
		return nil, fmt.Errorf("timeout rpc response")
	}
}

// ====

func authHeader(token string) http.Header {
	h := http.Header{}
	h.Set("Authorization", fmt.Sprintf("Bearer: %s", token))
	return h
}
