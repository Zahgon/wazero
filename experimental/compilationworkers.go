package experimental

import (
	"context"
)

func WithCompilationWorkers(ctx context.Context, workers int) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetCompilationWorkers(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }
