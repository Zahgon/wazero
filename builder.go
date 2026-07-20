package wazero

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/internal/wasm"
)

type HostFunctionBuilder interface {
	WithGoFunction(fn api.GoFunction, params, results []api.ValueType) HostFunctionBuilder

	WithGoModuleFunction(fn api.GoModuleFunction, params, results []api.ValueType) HostFunctionBuilder

	WithFunc(interface{}) HostFunctionBuilder

	WithName(name string) HostFunctionBuilder

	WithParameterNames(names ...string) HostFunctionBuilder

	WithResultNames(names ...string) HostFunctionBuilder

	Export(name string) HostModuleBuilder
}

type HostModuleBuilder interface {
	NewFunctionBuilder() HostFunctionBuilder

	Compile(context.Context) (CompiledModule, error)

	Instantiate(context.Context) (api.Module, error)
}

type hostModuleBuilder struct {
	r              *runtime
	moduleName     string
	exportNames    []string
	nameToHostFunc map[string]*wasm.HostFunc
}

func (r *runtime) NewHostModuleBuilder(moduleName string) HostModuleBuilder {
	_ = "STUB: not implemented"
	return *new(HostModuleBuilder)
}

type hostFunctionBuilder struct {
	b           *hostModuleBuilder
	fn          interface{}
	name        string
	paramNames  []string
	resultNames []string
}

func (h *hostFunctionBuilder) WithGoFunction(fn api.GoFunction, params, results []api.ValueType) HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (h *hostFunctionBuilder) WithGoModuleFunction(fn api.GoModuleFunction, params, results []api.ValueType) HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (h *hostFunctionBuilder) WithFunc(fn interface{}) HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (h *hostFunctionBuilder) WithName(name string) HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (h *hostFunctionBuilder) WithParameterNames(names ...string) HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (h *hostFunctionBuilder) WithResultNames(names ...string) HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (h *hostFunctionBuilder) Export(exportName string) HostModuleBuilder {
	_ = "STUB: not implemented"
	return *new(HostModuleBuilder)
}

func (b *hostModuleBuilder) ExportHostFunc(fn *wasm.HostFunc) { _ = "STUB: not implemented"; return }

func (b *hostModuleBuilder) NewFunctionBuilder() HostFunctionBuilder {
	_ = "STUB: not implemented"
	return *new(HostFunctionBuilder)
}

func (b *hostModuleBuilder) Compile(ctx context.Context) (CompiledModule, error) {
	_ = "STUB: not implemented"
	return *new(CompiledModule), nil
}

type hostModuleInstance struct{ api.Module }

func (h hostModuleInstance) ExportedFunction(name string) api.Function {
	_ = "STUB: not implemented"
	return *new(api.Function)
}

func (b *hostModuleBuilder) Instantiate(ctx context.Context) (api.Module, error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}
