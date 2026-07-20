package wazerotest

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/internal/internalapi"
)

const (
	exitStatusMarker = 1 << 63
)

type Module struct {
	internalapi.WazeroOnlyType

	ModuleName string

	Functions []*Function

	Globals []*Global

	ExportMemory *Memory

	exitStatus atomic.Uint64

	once                        sync.Once
	exportedFunctions           map[string]api.Function
	exportedFunctionDefinitions map[string]api.FunctionDefinition
	exportedGlobals             map[string]api.Global
	exportedMemoryDefinitions   map[string]api.MemoryDefinition
}

func NewModule(memory *Memory, functions ...*Function) *Module {
	_ = "STUB: not implemented"
	return nil
}

func (m *Module) String() string { _ = "STUB: not implemented"; return "" }

func (m *Module) Name() string { _ = "STUB: not implemented"; return "" }

func (m *Module) Memory() api.Memory { _ = "STUB: not implemented"; return *new(api.Memory) }

func (m *Module) ExportedFunction(name string) api.Function {
	_ = "STUB: not implemented"
	return *new(api.Function)
}

func (m *Module) ExportedFunctionDefinitions() map[string]api.FunctionDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (m *Module) ExportedMemory(name string) api.Memory {
	_ = "STUB: not implemented"
	return *new(api.Memory)
}

func (m *Module) ExportedMemoryDefinitions() map[string]api.MemoryDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (m *Module) ExportedGlobal(name string) api.Global {
	_ = "STUB: not implemented"
	return *new(api.Global)
}

func (m *Module) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Module) CloseWithExitCode(ctx context.Context, exitCode uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Module) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (m *Module) NumGlobal() int { _ = "STUB: not implemented"; return 0 }

func (m *Module) Global(i int) api.Global { _ = "STUB: not implemented"; return *new(api.Global) }

func (m *Module) NumFunction() int { _ = "STUB: not implemented"; return 0 }

func (m *Module) Function(i int) api.Function { _ = "STUB: not implemented"; return *new(api.Function) }

func (m *Module) ExitStatus() (exitCode uint32, exited bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Module) initialize() { _ = "STUB: not implemented"; return }

type Global struct {
	internalapi.WazeroOnlyType

	ValueType api.ValueType

	Value uint64

	ExportNames []string
}

func (g *Global) String() string { _ = "STUB: not implemented"; return "" }

func (g *Global) Type() api.ValueType { _ = "STUB: not implemented"; return *new(api.ValueType) }

func (g *Global) Get() uint64 { _ = "STUB: not implemented"; return 0 }

func GlobalI32(value int32, export ...string) *Global { _ = "STUB: not implemented"; return nil }

func GlobalI64(value int64, export ...string) *Global { _ = "STUB: not implemented"; return nil }

func GlobalF32(value float32, export ...string) *Global { _ = "STUB: not implemented"; return nil }

func GlobalF64(value float64, export ...string) *Global { _ = "STUB: not implemented"; return nil }

type Function struct {
	internalapi.WazeroOnlyType

	GoModuleFunction api.GoModuleFunction

	ParamTypes  []api.ValueType
	ResultTypes []api.ValueType

	FunctionName string
	DebugName    string
	ParamNames   []string
	ResultNames  []string
	ExportNames  []string

	module *Module
	index  int
}

func NewFunction(fn any) *Function { _ = "STUB: not implemented"; return nil }

var (
	errMissingFunctionSignature      = errors.New("missing function signature")
	errMissingFunctionModule         = errors.New("missing function module")
	errMissingFunctionImplementation = errors.New("missing function implementation")
)

func (f *Function) Definition() api.FunctionDefinition {
	_ = "STUB: not implemented"
	return *new(api.FunctionDefinition)
}

func (f *Function) Call(ctx context.Context, params ...uint64) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Function) CallWithStack(ctx context.Context, stack []uint64) error {
	_ = "STUB: not implemented"
	return nil
}

type functionDefinition struct {
	internalapi.WazeroOnlyType
	function *Function
}

func (def functionDefinition) Name() string { _ = "STUB: not implemented"; return "" }

func (def functionDefinition) DebugName() string { _ = "STUB: not implemented"; return "" }

func (def functionDefinition) GoFunction() any { _ = "STUB: not implemented"; return *new(any) }

func (def functionDefinition) ParamTypes() []api.ValueType { _ = "STUB: not implemented"; return nil }

func (def functionDefinition) ParamNames() []string { _ = "STUB: not implemented"; return nil }

func (def functionDefinition) ResultTypes() []api.ValueType { _ = "STUB: not implemented"; return nil }

func (def functionDefinition) ResultNames() []string { _ = "STUB: not implemented"; return nil }

func (def functionDefinition) ModuleName() string { _ = "STUB: not implemented"; return "" }

func (def functionDefinition) Index() uint32 { _ = "STUB: not implemented"; return 0 }

func (def functionDefinition) Import() (moduleName, name string, isImport bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (def functionDefinition) ExportNames() []string { _ = "STUB: not implemented"; return nil }

type Memory struct {
	internalapi.WazeroOnlyType

	Bytes []byte

	Min uint32
	Max uint32

	module *Module
}

func NewMemory(size int) *Memory { _ = "STUB: not implemented"; return nil }

func NewFixedMemory(size int) *Memory { _ = "STUB: not implemented"; return nil }

const PageSize = 65536

func (m *Memory) Definition() api.MemoryDefinition {
	_ = "STUB: not implemented"
	return *new(api.MemoryDefinition)
}

func (m *Memory) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Memory) Pages() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Memory) Grow(deltaPages uint32) (previousPages uint32, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Memory) ReadByte(offset uint32) (byte, bool) { _ = "STUB: not implemented"; return 0, false }

func (m *Memory) ReadUint16Le(offset uint32) (uint16, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Memory) ReadUint32Le(offset uint32) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Memory) ReadUint64Le(offset uint32) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Memory) ReadFloat32Le(offset uint32) (float32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Memory) ReadFloat64Le(offset uint32) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *Memory) Read(offset, length uint32) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *Memory) WriteByte(offset uint32, value byte) bool { _ = "STUB: not implemented"; return false }

func (m *Memory) WriteUint16Le(offset uint32, value uint16) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Memory) WriteUint32Le(offset uint32, value uint32) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Memory) WriteUint64Le(offset uint32, value uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Memory) WriteFloat32Le(offset uint32, value float32) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Memory) WriteFloat64Le(offset uint32, value float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Memory) Write(offset uint32, value []byte) bool { _ = "STUB: not implemented"; return false }

func (m *Memory) WriteString(offset uint32, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Memory) isOutOfRange(offset, length uint32) bool { _ = "STUB: not implemented"; return false }

type memoryDefinition struct {
	internalapi.WazeroOnlyType
	memory *Memory
}

func (def memoryDefinition) ModuleName() string { _ = "STUB: not implemented"; return "" }

func (def memoryDefinition) Index() uint32 { _ = "STUB: not implemented"; return 0 }

func (def memoryDefinition) Import() (moduleName, name string, isImport bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (def memoryDefinition) ExportNames() []string { _ = "STUB: not implemented"; return nil }

func (def memoryDefinition) Min() uint32 { _ = "STUB: not implemented"; return 0 }

func (def memoryDefinition) Max() (uint32, bool) { _ = "STUB: not implemented"; return 0, false }

var (
	_ api.Module   = (*Module)(nil)
	_ api.Function = (*Function)(nil)
	_ api.Global   = (*Global)(nil)
)
