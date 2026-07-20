package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var clockResGet = newHostFunc(wasip1.ClockResGetName, clockResGetFn, []wasm.ValueType{i32, i32}, "id", "result.resolution")

func clockResGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

var clockTimeGet = newHostFunc(wasip1.ClockTimeGetName, clockTimeGetFn, []wasm.ValueType{i32, i64, i32}, "id", "precision", "result.timestamp")

func clockTimeGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}
