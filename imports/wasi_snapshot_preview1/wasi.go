package wasi_snapshot_preview1

import (
	"context"
	"encoding/binary"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

const ModuleName = wasip1.InternalModuleName

const i32, i64 = wasm.ValueTypeI32, wasm.ValueTypeI64

var le = binary.LittleEndian

func MustInstantiate(ctx context.Context, r wazero.Runtime) { _ = "STUB: not implemented"; return }

func Instantiate(ctx context.Context, r wazero.Runtime) (api.Closer, error) {
	_ = "STUB: not implemented"
	return *new(api.Closer), nil
}

type Builder interface {
	Compile(context.Context) (wazero.CompiledModule, error)

	Instantiate(context.Context) (api.Closer, error)
}

func NewBuilder(r wazero.Runtime) Builder { _ = "STUB: not implemented"; return *new(Builder) }

type builder struct{ r wazero.Runtime }

func (b *builder) hostModuleBuilder() wazero.HostModuleBuilder {
	_ = "STUB: not implemented"
	return *new(wazero.HostModuleBuilder)
}

func (b *builder) Compile(ctx context.Context) (wazero.CompiledModule, error) {
	_ = "STUB: not implemented"
	return *new(wazero.CompiledModule), nil
}

func (b *builder) Instantiate(ctx context.Context) (api.Closer, error) {
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

func exportFunctions(builder wazero.HostModuleBuilder) { _ = "STUB: not implemented"; return }

func writeOffsetsAndNullTerminatedValues(mem api.Memory, values [][]byte, offsets, bytes, bytesLen uint32) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

func newHostFunc(
	name string,
	goFunc wasiFunc,
	paramTypes []wasm.ValueType,
	paramNames ...string,
) *wasm.HostFunc {
	_ = "STUB: not implemented"
	return nil
}

type wasiFunc func(ctx context.Context, mod api.Module, params []uint64) sys.Errno

func (f wasiFunc) Call(ctx context.Context, mod api.Module, stack []uint64) {
	_ = "STUB: not implemented"
	return
}

func stubFunction(name string, paramTypes []wasm.ValueType, paramNames ...string) *wasm.HostFunc {
	_ = "STUB: not implemented"
	return nil
}
