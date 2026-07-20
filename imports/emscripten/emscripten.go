package emscripten

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/internal/wasm"
)

const i32 = wasm.ValueTypeI32

func MustInstantiate(ctx context.Context, r wazero.Runtime) { _ = "STUB: not implemented"; return }

func Instantiate(ctx context.Context, r wazero.Runtime) (api.Closer, error) {
	_ = "STUB: not implemented"
	return *new(api.Closer), nil
}

type FunctionExporter interface {
	ExportFunctions(wazero.HostModuleBuilder)
}

func NewFunctionExporter() FunctionExporter {
	_ = "STUB: not implemented"
	return *new(FunctionExporter)
}

type functionExporter struct{}

func (functionExporter) ExportFunctions(builder wazero.HostModuleBuilder) {
	_ = "STUB: not implemented"
	return
}

type emscriptenFns []*wasm.HostFunc

func InstantiateForModule(ctx context.Context, r wazero.Runtime, guest wazero.CompiledModule) (api.Closer, error) {
	_ = "STUB: not implemented"
	return *new(api.Closer), nil
}

func NewFunctionExporterForModule(guest wazero.CompiledModule) (FunctionExporter, error) {
	_ = "STUB: not implemented"
	return *new(FunctionExporter), nil
}

func (i emscriptenFns) ExportFunctions(builder wazero.HostModuleBuilder) {
	_ = "STUB: not implemented"
	return
}
