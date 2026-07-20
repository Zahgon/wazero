package experimental

import (
	"context"
)

type CloseNotifier interface {
	CloseNotify(ctx context.Context, exitCode uint32)
}

type CloseNotifyFunc func(ctx context.Context, exitCode uint32)

func (f CloseNotifyFunc) CloseNotify(ctx context.Context, exitCode uint32) {
	_ = "STUB: not implemented"
	return
}

func WithCloseNotifier(ctx context.Context, notifier CloseNotifier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
