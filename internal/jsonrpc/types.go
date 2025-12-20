package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RPCRequest - a JSON-RPC request.
type RPCRequest struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

// RPCResponse - a JSON-RPC response.
type RPCResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  interface{}     `json:"error,omitempty"`
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

	var data T

	// Check if Result is empty; return zero value of T
	if len(rs.Result) == 0 {
		return &data, nil
	}

	// Decode JSON from Result into type T
	err := json.Unmarshal(rs.Result, &data)
	if err != nil {
		return nil, fmt.Errorf("rpc response decode error: %w", err)
	}

	return &data, nil
}
