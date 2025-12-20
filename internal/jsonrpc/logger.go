package jsonrpc

import (
	"fmt"
	"log"
)

// Logger returns the client's logger.
// If no custom logger is set, it falls back to the default logger using log.Println.
func (c *JsonRPCClient) Log() Logger {
	if c.logger != nil {
		return c.logger
	}
	return defaultLogger{}
}

// defaultLogger implements Logger using the standard log package.
type defaultLogger struct{}

func (l defaultLogger) Info(v ...any) {
	log.Println("[INFO]", fmt.Sprint(v...))
}

func (l defaultLogger) Error(v ...any) {
	log.Println("[ERROR]", fmt.Sprint(v...))
}
