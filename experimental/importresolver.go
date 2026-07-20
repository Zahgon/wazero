package experimental

import (
	"context"

	"github.com/tetratelabs/wazero/api"
)

type ImportResolver func(name string) api.Module

func WithImportResolver(ctx context.Context, resolver ImportResolver) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
