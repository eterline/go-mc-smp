package jsonrpc

import (
	"encoding/json"
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
	Error  json.RawMessage   `json:"error,omitempty"`
}

func (r RPCResponse) ParamsNotEmpty() error {
	if len(r.Params) == 0 {
		return ErrResponseParamsEmpty
	}
	return nil
}

func (r RPCResponse) ResultNotEmpty() error {
	if len(r.Result) == 0 {
		return ErrResponseResultEmpty
	}
	return nil
}

func (r RPCResponse) Err() error {
	if r.Error != nil {
		embedErr := fmt.Errorf("%v", r.Error)
		return ErrResponseContains.Wrap(embedErr)
	}
	return nil
}

func DecodeRPCResult[T any](r *RPCResponse) (*T, error) {
	if r == nil {
		return nil, ErrResponseNil
	}

	if err := r.Err(); err != nil {
		return nil, err
	}

	if err := r.ResultNotEmpty(); err != nil {
		return nil, err
	}

	var data T

	if err := json.Unmarshal(r.Result, &data); err != nil {
		return nil, ErrDecodeResponse.Wrap(err)
	}

	return &data, nil
}

func DecodeRPCParams[T any](r *RPCResponse) (*T, error) {
	if r == nil {
		return nil, ErrResponseNil
	}

	if err := r.Err(); err != nil {
		return nil, err
	}

	if err := r.ParamsNotEmpty(); err != nil {
		return nil, err
	}

	var data T

	if err := json.Unmarshal(r.Params[0], &data); err != nil {
		return nil, ErrDecodeResponse.Wrap(err)
	}

	return &data, nil
}
