package gomcsmp

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
)

type ClientOption func(*clientConfig)

type clientConfig struct {
	path        string
	tls         bool
	callTimeout time.Duration
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		path:        "/",
		tls:         false,
		callTimeout: 5 * time.Second,
	}
}

// ================

func WithPath(path string) ClientOption {
	return func(cfg *clientConfig) {
		if path == "" {
			return
		}
		if path[0] != '/' {
			path = "/" + path
		}
		cfg.path = path
	}
}

func WithTLS() ClientOption {
	return func(cfg *clientConfig) {
		cfg.tls = true
	}
}

func WithCallTimeout(t time.Duration) ClientOption {
	return func(cfg *clientConfig) {
		cfg.callTimeout = t
	}
}

// ================

type RPCClient struct {
	core *jsonrpc.JsonRPCClient
}

func NewClient(host string, port uint16, token string, opts ...ClientOption) (*RPCClient, error) {
	joinedHost, err := joinHostPort(host, port)
	if err != nil {
		return nil, err
	}

	if err := validateToken(token); err != nil {
		return nil, err
	}

	cfg := defaultClientConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	u := &url.URL{
		Scheme: wsMode(cfg.tls),
		Host:   joinedHost,
		Path:   cfg.path,
	}

	core, err := jsonrpc.NewJsonRPCClient(u.String(), token, cfg.callTimeout)
	if err != nil {
		return nil, err
	}

	return &RPCClient{
		core: core,
	}, nil
}

func (rpc *RPCClient) Close() error {
	return rpc.core.Close()
}

// ================

func joinHostPort(host string, port uint16) (string, error) {
	if host == "" {
		return "", errors.New("invalid host")
	}
	if port == 0 {
		return "", errors.New("invalid port")
	}
	return fmt.Sprintf("%s:%d", host, port), nil
}

func wsMode(tls bool) string {
	if tls {
		return "wss"
	}
	return "ws"
}

func validateToken(token string) error {
	if len(token) != 40 {
		return fmt.Errorf("invalid token length: expected 40, got %d", len(token))
	}
	return nil
}
