package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var argsGet = newHostFunc(wasip1.ArgsGetName, argsGetFn, []wasm.ValueType{i32, i32}, "argv", "argv_buf")

func argsGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

var argsSizesGet = newHostFunc(wasip1.ArgsSizesGetName, argsSizesGetFn, []wasm.ValueType{i32, i32}, "result.argc", "result.argv_len")

func argsSizesGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}
