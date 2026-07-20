package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
)

var schedYield = newHostFunc(wasip1.SchedYieldName, schedYieldFn, nil)

func schedYieldFn(_ context.Context, mod api.Module, _ []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}
