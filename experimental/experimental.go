package experimental

import (
	"github.com/tetratelabs/wazero/api"
)

type InternalModule interface {
	api.Module

	NumGlobal() int

	Global(i int) api.Global
}

type ProgramCounter uint64

type InternalFunction interface {
	Definition() api.FunctionDefinition

	SourceOffsetForPC(pc ProgramCounter) uint64
}
