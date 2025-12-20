package jsonrpc

import (
	"errors"
	"fmt"
)

// RPCRequest - a JSON-RPC request.
type RPCRequest struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Params any    `json:"params,omitempty"`
}

// RPCResponse - a JSON-RPC response.
type RPCResponse struct {
	ID     int         `json:"id,omitempty"`
	Method string      `json:"method"`
	Params any         `json:"params,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

// DecodeRPCResponse - decodes a JSON-RPC response into a specific type T.
func DecodeRPCResponse[T any](rs *RPCResponse) (*T, error) {
	// Check for nil response
	if rs == nil {
		return nil, errors.New("rpc response is nil")
	}

	// Check for server-side error
	if rs.Error != nil {
		return nil, fmt.Errorf("rpc response contains error: %v", rs.Error)
	}

	data, ok := rs.Params.(T)
	if !ok {
		return nil, fmt.Errorf("rpc response invalid type to decode: %v", rs.Error)
	}

	return &data, nil
}
