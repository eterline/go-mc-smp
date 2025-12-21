package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RPCRequest - a JSON-RPC request.
type RPCRequest struct {
	ID     int               `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params,omitempty"`
}

func NewRPCRequest(id int, method string, params []any) (*RPCRequest, error) {
	req := &RPCRequest{
		ID:     id,
		Method: method,
	}

	if len(params) == 0 {
		return req, nil
	}

	raw := make([]json.RawMessage, 0, len(params))
	for i, p := range params {
		b, err := json.Marshal(p)
		if err != nil {
			return nil, fmt.Errorf("marshal param %d: %w", i, err)
		}
		raw = append(raw, json.RawMessage(b))
	}

	req.Params = raw
	return req, nil
}

// RPCResponse - a JSON-RPC response.
type RPCResponse struct {
	ID     int               `json:"id,omitempty"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params,omitempty"`
	Result json.RawMessage   `json:"result,omitempty"`
	Error  interface{}       `json:"error,omitempty"`
}

func (r RPCResponse) Err() error {
	if r.Error != nil {
		return fmt.Errorf("rpc response contains error: %v", r.Error)
	}
	return nil
}

func DecodeRPCResult[T any](rs *RPCResponse) (*T, error) {
	if rs == nil {
		return nil, errors.New("rpc response is nil")
	}

	if rs.Error != nil {
		return nil, rs.Err()
	}

	if len(rs.Result) == 0 {
		return nil, errors.New("rpc result is empty")
	}

	var data T
	if err := json.Unmarshal(rs.Result, &data); err != nil {
		return nil, fmt.Errorf("rpc result decode failed: %w", err)
	}

	return &data, nil
}

func DecodeRPCParams[T any](rs *RPCResponse) (*T, error) {
	if rs == nil {
		return nil, errors.New("rpc response is nil")
	}

	if len(rs.Params) != 1 {
		return nil, fmt.Errorf("expected 1 param, got %d", len(rs.Params))
	}

	var data T
	if err := json.Unmarshal(rs.Params[0], &data); err != nil {
		return nil, fmt.Errorf("rpc params decode failed: %w", err)
	}

	return &data, nil
}

func DecodeRPCParamsInto[T any](rs *RPCResponse) (*T, error) {
	if rs == nil {
		return nil, errors.New("rpc response is nil")
	}

	if len(rs.Params) == 0 {
		return nil, errors.New("rpc params is empty")
	}

	raw, err := json.Marshal(rs.Params)
	if err != nil {
		return nil, err
	}

	var data T
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("rpc params decode failed: %w", err)
	}

	return &data, nil
}
