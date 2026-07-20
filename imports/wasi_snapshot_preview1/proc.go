package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var procExit = &wasm.HostFunc{
	ExportName: wasip1.ProcExitName,
	Name:       wasip1.ProcExitName,
	ParamTypes: []wasm.ValueType{i32},
	ParamNames: []string{"rval"},
	Code:       wasm.Code{GoFunc: api.GoModuleFunc(procExitFn)},
}

func procExitFn(ctx context.Context, mod api.Module, params []uint64) {
	_ = "STUB: not implemented"
	return
}

var procRaise = stubFunction(wasip1.ProcRaiseName, []wasm.ValueType{i32}, "sig")
