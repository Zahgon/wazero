package experimental

import (
	"context"
)

type Snapshot interface {
	Restore(ret []uint64)
}

type Snapshotter interface {
	Snapshot() Snapshot
}

func WithSnapshotter(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetSnapshotter(ctx context.Context) Snapshotter {
	_ = "STUB: not implemented"
	return *new(Snapshotter)
}
