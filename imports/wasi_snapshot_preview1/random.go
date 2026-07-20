package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var randomGet = newHostFunc(wasip1.RandomGetName, randomGetFn, []wasm.ValueType{i32, i32}, "buf", "buf_len")

func randomGetFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}
