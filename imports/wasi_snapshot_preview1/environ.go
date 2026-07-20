package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var environGet = newHostFunc(wasip1.EnvironGetName, environGetFn, []wasm.ValueType{i32, i32}, "environ", "environ_buf")

func environGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

var environSizesGet = newHostFunc(wasip1.EnvironSizesGetName, environSizesGetFn, []wasm.ValueType{i32, i32}, "result.environc", "result.environv_len")

func environSizesGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}
