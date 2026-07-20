package wazero

import (
	"context"
	"sync/atomic"

	"github.com/tetratelabs/wazero/api"
	experimentalapi "github.com/tetratelabs/wazero/experimental"
	"github.com/tetratelabs/wazero/internal/wasm"
)

type Runtime interface {
	Instantiate(ctx context.Context, source []byte) (api.Module, error)

	InstantiateWithConfig(ctx context.Context, source []byte, config ModuleConfig) (api.Module, error)

	NewHostModuleBuilder(moduleName string) HostModuleBuilder

	CompileModule(ctx context.Context, binary []byte) (CompiledModule, error)

	InstantiateModule(ctx context.Context, compiled CompiledModule, config ModuleConfig) (api.Module, error)

	CloseWithExitCode(ctx context.Context, exitCode uint32) error

	Module(moduleName string) api.Module

	api.Closer
}

func NewRuntime(ctx context.Context) Runtime { _ = "STUB: not implemented"; return *new(Runtime) }

func NewRuntimeWithConfig(ctx context.Context, rConfig RuntimeConfig) Runtime {
	_ = "STUB: not implemented"
	return *new(Runtime)
}

type runtime struct {
	store                 *wasm.Store
	cache                 *cache
	enabledFeatures       api.CoreFeatures
	memoryLimitPages      uint32
	memoryCapacityFromMax bool
	dwarfDisabled         bool
	storeCustomSections   bool

	closed atomic.Uint64

	ensureTermination bool
}

func (r *runtime) Module(moduleName string) api.Module {
	_ = "STUB: not implemented"
	return *new(api.Module)
}

func (r *runtime) CompileModule(ctx context.Context, binary []byte) (CompiledModule, error) {
	_ = "STUB: not implemented"
	return *new(CompiledModule), nil
}

func buildFunctionListeners(ctx context.Context, internal *wasm.Module) ([]experimentalapi.FunctionListener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runtime) failIfClosed() error { _ = "STUB: not implemented"; return nil }

func (r *runtime) Instantiate(ctx context.Context, binary []byte) (api.Module, error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}

func (r *runtime) InstantiateWithConfig(ctx context.Context, binary []byte, config ModuleConfig) (api.Module, error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}

func (r *runtime) InstantiateModule(
	ctx context.Context,
	compiled CompiledModule,
	mConfig ModuleConfig,
) (mod api.Module, err error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}

func (r *runtime) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *runtime) CloseWithExitCode(ctx context.Context, exitCode uint32) error {
	_ = "STUB: not implemented"
	return nil
}
