package assemblyscript

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	experimentalsys "github.com/tetratelabs/wazero/experimental/sys"
	. "github.com/tetratelabs/wazero/internal/assemblyscript"
	internalsys "github.com/tetratelabs/wazero/internal/sys"
	"github.com/tetratelabs/wazero/internal/wasm"
)

const (
	i32, f64 = wasm.ValueTypeI32, wasm.ValueTypeF64
)

func MustInstantiate(ctx context.Context, r wazero.Runtime) { _ = "STUB: not implemented"; return }

func Instantiate(ctx context.Context, r wazero.Runtime) (api.Closer, error) {
	_ = "STUB: not implemented"
	return *new(api.Closer), nil
}

type FunctionExporter interface {
	WithAbortMessageDisabled() FunctionExporter

	WithTraceToStdout() FunctionExporter

	WithTraceToStderr() FunctionExporter

	ExportFunctions(wazero.HostModuleBuilder)
}

func NewFunctionExporter() FunctionExporter {
	_ = "STUB: not implemented"
	return *new(FunctionExporter)
}

type functionExporter struct {
	abortFn, traceFn *wasm.HostFunc
}

func (e *functionExporter) WithAbortMessageDisabled() FunctionExporter {
	_ = "STUB: not implemented"
	return *new(FunctionExporter)
}

func (e *functionExporter) WithTraceToStdout() FunctionExporter {
	_ = "STUB: not implemented"
	return *new(FunctionExporter)
}

func (e *functionExporter) WithTraceToStderr() FunctionExporter {
	_ = "STUB: not implemented"
	return *new(FunctionExporter)
}

func (e *functionExporter) ExportFunctions(builder wazero.HostModuleBuilder) {
	_ = "STUB: not implemented"
	return
}

var abortMessageEnabled = &wasm.HostFunc{
	ExportName: AbortName,
	Name:       "~lib/builtins/abort",
	ParamTypes: []wasm.ValueType{i32, i32, i32, i32},
	ParamNames: []string{"message", "fileName", "lineNumber", "columnNumber"},
	Code:       wasm.Code{GoFunc: api.GoModuleFunc(abortWithMessage)},
}

var abortMessageDisabled = abortMessageEnabled.WithGoModuleFunc(abort)

func abortWithMessage(ctx context.Context, mod api.Module, stack []uint64) {
	_ = "STUB: not implemented"
	return
}

func abort(ctx context.Context, mod api.Module, _ []uint64) { _ = "STUB: not implemented"; return }

var traceDisabled = traceStdout.WithGoModuleFunc(func(context.Context, api.Module, []uint64) {})

var traceStdout = &wasm.HostFunc{
	ExportName: TraceName,
	Name:       "~lib/builtins/trace",
	ParamTypes: []wasm.ValueType{i32, i32, f64, f64, f64, f64, f64},
	ParamNames: []string{"message", "nArgs", "arg0", "arg1", "arg2", "arg3", "arg4"},
	Code: wasm.Code{
		GoFunc: api.GoModuleFunc(func(_ context.Context, mod api.Module, stack []uint64) {
			fsc := mod.(*wasm.ModuleInstance).Sys.FS()
			if stdout, ok := fsc.LookupFile(internalsys.FdStdout); ok {
				traceTo(mod, stack, stdout.File)
			}
		}),
	},
}

var traceStderr = traceStdout.WithGoModuleFunc(func(_ context.Context, mod api.Module, stack []uint64) {
	fsc := mod.(*wasm.ModuleInstance).Sys.FS()
	if stderr, ok := fsc.LookupFile(internalsys.FdStderr); ok {
		traceTo(mod, stack, stderr.File)
	}
})

func traceTo(mod api.Module, params []uint64, file experimentalsys.File) {
	_ = "STUB: not implemented"
	return
}

func formatFloat(f float64) string { _ = "STUB: not implemented"; return "" }

var seed = &wasm.HostFunc{
	ExportName:  SeedName,
	Name:        "~lib/builtins/seed",
	ResultTypes: []wasm.ValueType{f64},
	ResultNames: []string{"rand"},
	Code: wasm.Code{
		GoFunc: api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
			r := mod.(*wasm.ModuleInstance).Sys.RandSource()
			buf := make([]byte, 8)
			_, err := io.ReadFull(r, buf)
			if err != nil {
				panic(fmt.Errorf("error reading random seed: %w", err))
			}

			stack[0] = binary.LittleEndian.Uint64(buf)
		}),
	},
}

func readAssemblyScriptString(mem api.Memory, offset uint32) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func decodeUTF16(b []byte) string { _ = "STUB: not implemented"; return "" }
