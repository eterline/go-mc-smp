package jsonrpc

import "fmt"

type jsonrpcError struct {
	embed error
	data  string
}

// ==================

func newJsonrpcError(data string) *jsonrpcError {
	return &jsonrpcError{
		data: data,
	}
}

func (e *jsonrpcError) Error() string {
	if e.embed != nil {
		return fmt.Sprintf("jsonrpc error: %s: %s", e.data, e.embed.Error())
	}
	return fmt.Sprintf("jsonrpc error: %s", e.data)
}

func (e *jsonrpcError) Wrap(err error) error {
	e.embed = err
	return e
}

func (jerr *jsonrpcError) Unwrap() error {
	return jerr.embed
}

// ==================

var (
	ErrReadResponse = newJsonrpcError("read response error")
	ErrWriteRequest = newJsonrpcError("write request error")

	ErrDecodeResponse = newJsonrpcError("decode response error")
	ErrEncodeRequest  = newJsonrpcError("encode request error")

	ErrContext  = newJsonrpcError("context error")
	ErrRpcClose = newJsonrpcError("close rpc error")

	ErrResponseChannelClosed = newJsonrpcError("rpc response channel closed")
	ErrRequestChannelClosed  = newJsonrpcError("rpc request channel closed")

	ErrNotifyChannelOverflow = newJsonrpcError("notification channel full")

	ErrResponseNil         = newJsonrpcError("rpc response is nil")
	ErrResponseResultEmpty = newJsonrpcError("rpc response result empty")
	ErrResponseParamsEmpty = newJsonrpcError("rpc response params empty")
	ErrResponseContains    = newJsonrpcError("rpc response contains error")
)
