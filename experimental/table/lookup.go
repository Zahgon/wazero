package table

import (
	"github.com/tetratelabs/wazero/api"
)

func LookupFunction(
	module api.Module, tableIndex uint32, tableOffset uint32,
	expectedParamTypes, expectedResultTypes []api.ValueType,
) api.Function {
	_ = "STUB: not implemented"
	return *new(api.Function)
}
