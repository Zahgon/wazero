package wazero

import (
	"context"
	"io"
	"io/fs"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/internal/filecache"
	"github.com/tetratelabs/wazero/internal/internalapi"
	internalsock "github.com/tetratelabs/wazero/internal/sock"
	internalsys "github.com/tetratelabs/wazero/internal/sys"
	"github.com/tetratelabs/wazero/internal/wasm"
	"github.com/tetratelabs/wazero/sys"
)

type RuntimeConfig interface {
	WithCoreFeatures(api.CoreFeatures) RuntimeConfig

	WithMemoryLimitPages(memoryLimitPages uint32) RuntimeConfig

	WithMemoryCapacityFromMax(memoryCapacityFromMax bool) RuntimeConfig

	WithDebugInfoEnabled(bool) RuntimeConfig

	WithCompilationCache(CompilationCache) RuntimeConfig

	WithCustomSections(bool) RuntimeConfig

	WithCloseOnContextDone(bool) RuntimeConfig
}

func NewRuntimeConfig() RuntimeConfig { _ = "STUB: not implemented"; return *new(RuntimeConfig) }

type newEngine func(context.Context, api.CoreFeatures, filecache.Cache) wasm.Engine

type runtimeConfig struct {
	enabledFeatures       api.CoreFeatures
	memoryLimitPages      uint32
	memoryCapacityFromMax bool
	engineKind            engineKind
	dwarfDisabled         bool
	newEngine             newEngine
	cache                 CompilationCache
	storeCustomSections   bool
	ensureTermination     bool
}

var engineLessConfig = &runtimeConfig{
	enabledFeatures:       api.CoreFeaturesV2,
	memoryLimitPages:      wasm.MemoryLimitPages,
	memoryCapacityFromMax: false,
	dwarfDisabled:         false,
}

type engineKind int

const (
	engineKindAuto engineKind = iota - 1
	engineKindCompiler
	engineKindInterpreter
	engineKindCount
)

func NewRuntimeConfigCompiler() RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func NewRuntimeConfigInterpreter() RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) clone() *runtimeConfig { _ = "STUB: not implemented"; return nil }

func (c *runtimeConfig) WithCoreFeatures(features api.CoreFeatures) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) WithCloseOnContextDone(ensure bool) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) WithMemoryLimitPages(memoryLimitPages uint32) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) WithCompilationCache(ca CompilationCache) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) WithMemoryCapacityFromMax(memoryCapacityFromMax bool) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) WithDebugInfoEnabled(dwarfEnabled bool) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

func (c *runtimeConfig) WithCustomSections(storeCustomSections bool) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

type CompiledModule interface {
	Name() string

	ImportedFunctions() []api.FunctionDefinition

	ExportedFunctions() map[string]api.FunctionDefinition

	ImportedMemories() []api.MemoryDefinition

	ExportedMemories() map[string]api.MemoryDefinition

	CustomSections() []api.CustomSection

	Close(context.Context) error
}

var _ CompiledModule = &compiledModule{}

type compiledModule struct {
	module *wasm.Module

	compiledEngine wasm.Engine

	closeWithModule bool
	typeIDs         []wasm.FunctionTypeID
}

func (c *compiledModule) Name() (moduleName string) { _ = "STUB: not implemented"; return "" }

func (c *compiledModule) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *compiledModule) ImportedFunctions() []api.FunctionDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (c *compiledModule) ExportedFunctions() map[string]api.FunctionDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (c *compiledModule) ImportedMemories() []api.MemoryDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (c *compiledModule) ExportedMemories() map[string]api.MemoryDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (c *compiledModule) CustomSections() []api.CustomSection {
	_ = "STUB: not implemented"
	return nil
}

type customSection struct {
	internalapi.WazeroOnlyType
	name string
	data []byte
}

func (c *customSection) Name() string { _ = "STUB: not implemented"; return "" }

func (c *customSection) Data() []byte { _ = "STUB: not implemented"; return nil }

type ModuleConfig interface {
	WithArgs(...string) ModuleConfig

	WithEnv(key, value string) ModuleConfig

	WithFS(fs.FS) ModuleConfig

	WithFSConfig(FSConfig) ModuleConfig

	WithName(string) ModuleConfig

	WithStartFunctions(...string) ModuleConfig

	WithStderr(io.Writer) ModuleConfig

	WithStdin(io.Reader) ModuleConfig

	WithStdout(io.Writer) ModuleConfig

	WithWalltime(sys.Walltime, sys.ClockResolution) ModuleConfig

	WithSysWalltime() ModuleConfig

	WithNanotime(sys.Nanotime, sys.ClockResolution) ModuleConfig

	WithSysNanotime() ModuleConfig

	WithNanosleep(sys.Nanosleep) ModuleConfig

	WithOsyield(sys.Osyield) ModuleConfig

	WithSysNanosleep() ModuleConfig

	WithRandSource(io.Reader) ModuleConfig
}

type moduleConfig struct {
	name               string
	nameSet            bool
	startFunctions     []string
	stdin              io.Reader
	stdout             io.Writer
	stderr             io.Writer
	randSource         io.Reader
	walltime           sys.Walltime
	walltimeResolution sys.ClockResolution
	nanotime           sys.Nanotime
	nanotimeResolution sys.ClockResolution
	nanosleep          sys.Nanosleep
	osyield            sys.Osyield
	args               [][]byte

	environ [][]byte

	environKeys map[string]int

	fsConfig FSConfig

	sockConfig *internalsock.Config
}

func NewModuleConfig() ModuleConfig { _ = "STUB: not implemented"; return *new(ModuleConfig) }

func (c *moduleConfig) clone() *moduleConfig { _ = "STUB: not implemented"; return nil }

func (c *moduleConfig) WithArgs(args ...string) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func toByteSlices(strings []string) (result [][]byte) { _ = "STUB: not implemented"; return nil }

func (c *moduleConfig) WithEnv(key, value string) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithFS(fs fs.FS) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithFSConfig(config FSConfig) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithName(name string) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithStartFunctions(startFunctions ...string) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithStderr(stderr io.Writer) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithStdin(stdin io.Reader) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithStdout(stdout io.Writer) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithWalltime(walltime sys.Walltime, resolution sys.ClockResolution) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithSysWalltime() ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithNanotime(nanotime sys.Nanotime, resolution sys.ClockResolution) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithSysNanotime() ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithNanosleep(nanosleep sys.Nanosleep) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithOsyield(osyield sys.Osyield) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithSysNanosleep() ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) WithRandSource(source io.Reader) ModuleConfig {
	_ = "STUB: not implemented"
	return *new(ModuleConfig)
}

func (c *moduleConfig) toSysContext() (sysCtx *internalsys.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
