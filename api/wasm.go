package api

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero/internal/internalapi"
)

type ExternType = byte

const (
	ExternTypeFunc   ExternType = 0x00
	ExternTypeTable  ExternType = 0x01
	ExternTypeMemory ExternType = 0x02
	ExternTypeGlobal ExternType = 0x03
)

const (
	ExternTypeFuncName = "func"

	ExternTypeTableName = "table"

	ExternTypeMemoryName = "memory"

	ExternTypeGlobalName = "global"
)

func ExternTypeName(et ExternType) string { _ = "STUB: not implemented"; return "" }

type ValueType = byte

const (
	ValueTypeI32 ValueType = 0x7f

	ValueTypeI64 ValueType = 0x7e

	ValueTypeF32 ValueType = 0x7d

	ValueTypeF64 ValueType = 0x7c

	ValueTypeExternref ValueType = 0x6f
)

func ValueTypeName(t ValueType) string { _ = "STUB: not implemented"; return "" }

type Module interface {
	fmt.Stringer

	Name() string

	Memory() Memory

	ExportedFunction(name string) Function

	ExportedFunctionDefinitions() map[string]FunctionDefinition

	ExportedMemory(name string) Memory

	ExportedMemoryDefinitions() map[string]MemoryDefinition

	ExportedGlobal(name string) Global

	CloseWithExitCode(ctx context.Context, exitCode uint32) error

	Closer

	IsClosed() bool

	internalapi.WazeroOnly
}

type Closer interface {
	Close(context.Context) error
}

type ExportDefinition interface {
	ModuleName() string

	Index() uint32

	Import() (moduleName, name string, isImport bool)

	ExportNames() []string

	internalapi.WazeroOnly
}

type MemoryDefinition interface {
	ExportDefinition

	Min() uint32

	Max() (uint32, bool)

	internalapi.WazeroOnly
}

type FunctionDefinition interface {
	ExportDefinition

	Name() string

	DebugName() string

	GoFunction() interface{}

	ParamTypes() []ValueType

	ParamNames() []string

	ResultTypes() []ValueType

	ResultNames() []string

	internalapi.WazeroOnly
}

type Function interface {
	Definition() FunctionDefinition

	Call(ctx context.Context, params ...uint64) ([]uint64, error)

	CallWithStack(ctx context.Context, stack []uint64) error

	internalapi.WazeroOnly
}

type GoModuleFunction interface {
	Call(ctx context.Context, mod Module, stack []uint64)
}

type GoModuleFunc func(ctx context.Context, mod Module, stack []uint64)

func (f GoModuleFunc) Call(ctx context.Context, mod Module, stack []uint64) {
	_ = "STUB: not implemented"
	return
}

type GoFunction interface {
	Call(ctx context.Context, stack []uint64)
}

type GoFunc func(ctx context.Context, stack []uint64)

func (f GoFunc) Call(ctx context.Context, stack []uint64) { _ = "STUB: not implemented"; return }

type Global interface {
	fmt.Stringer

	Type() ValueType

	Get() uint64
}

type MutableGlobal interface {
	Global

	Set(v uint64)

	internalapi.WazeroOnly
}

type Memory interface {
	Definition() MemoryDefinition

	Size() uint32

	Grow(deltaPages uint32) (previousPages uint32, ok bool)

	ReadByte(offset uint32) (byte, bool)

	ReadUint16Le(offset uint32) (uint16, bool)

	ReadUint32Le(offset uint32) (uint32, bool)

	ReadFloat32Le(offset uint32) (float32, bool)

	ReadUint64Le(offset uint32) (uint64, bool)

	ReadFloat64Le(offset uint32) (float64, bool)

	Read(offset, byteCount uint32) ([]byte, bool)

	WriteByte(offset uint32, v byte) bool

	WriteUint16Le(offset uint32, v uint16) bool

	WriteUint32Le(offset, v uint32) bool

	WriteFloat32Le(offset uint32, v float32) bool

	WriteUint64Le(offset uint32, v uint64) bool

	WriteFloat64Le(offset uint32, v float64) bool

	Write(offset uint32, v []byte) bool

	WriteString(offset uint32, v string) bool

	internalapi.WazeroOnly
}

type CustomSection interface {
	Name() string

	Data() []byte

	internalapi.WazeroOnly
}

func EncodeExternref(input uintptr) uint64 { _ = "STUB: not implemented"; return 0 }

func DecodeExternref(input uint64) uintptr { _ = "STUB: not implemented"; return 0 }

func EncodeI32(input int32) uint64 { _ = "STUB: not implemented"; return 0 }

func DecodeI32(input uint64) int32 { _ = "STUB: not implemented"; return 0 }

func EncodeU32(input uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func DecodeU32(input uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func EncodeI64(input int64) uint64 { _ = "STUB: not implemented"; return 0 }

func EncodeF32(input float32) uint64 { _ = "STUB: not implemented"; return 0 }

func DecodeF32(input uint64) float32 { _ = "STUB: not implemented"; return 0 }

func EncodeF64(input float64) uint64 { _ = "STUB: not implemented"; return 0 }

func DecodeF64(input uint64) float64 { _ = "STUB: not implemented"; return 0 }
