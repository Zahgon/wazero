package sock

import (
	"context"

	"github.com/tetratelabs/wazero/internal/sock"
)

type Config interface {
	WithTCPListener(host string, port int) Config
}

func NewConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type internalSockConfig struct {
	c *sock.Config
}

func (c *internalSockConfig) WithTCPListener(host string, port int) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}

func WithConfig(ctx context.Context, config Config) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
