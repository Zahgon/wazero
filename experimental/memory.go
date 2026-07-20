package experimental

import (
	"context"
)

type MemoryAllocator interface {
	Allocate(cap, max uint64) LinearMemory
}

type MemoryAllocatorFunc func(cap, max uint64) LinearMemory

func (f MemoryAllocatorFunc) Allocate(cap, max uint64) LinearMemory {
	_ = "STUB: not implemented"
	return *new(LinearMemory)
}

type LinearMemory interface {
	Reallocate(size uint64) []byte

	Free()
}

func WithMemoryAllocator(ctx context.Context, allocator MemoryAllocator) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
