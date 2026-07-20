package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func main() {

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasm, err := resultOffsetWasmFunctions(ctx, r)
	if err != nil {
		log.Panicln(err)
	}

	runtimeWithMultiValue := wazero.NewRuntime(ctx)

	wasmWithMultiValue, err := multiValueWasmFunctions(ctx, runtimeWithMultiValue)
	if err != nil {
		log.Panicln(err)
	}

	multiValueFromImportedHost, err := multiValueFromImportedHostWasmFunctions(ctx, runtimeWithMultiValue)
	if err != nil {
		log.Panicln(err)
	}

	for _, mod := range []api.Module{wasm, wasmWithMultiValue, multiValueFromImportedHost} {
		getAge := mod.ExportedFunction("call_get_age")
		results, err := getAge.Call(ctx)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("%s: age=%d\n", mod.Name(), results[0])
	}
}

//go:embed testdata/result_offset.wasm
var resultOffsetWasm []byte

func resultOffsetWasmFunctions(ctx context.Context, r wazero.Runtime) (api.Module, error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}

//go:embed testdata/multi_value.wasm
var multiValueWasm []byte

func multiValueWasmFunctions(ctx context.Context, r wazero.Runtime) (api.Module, error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}

//go:embed testdata/multi_value_imported.wasm
var multiValueFromImportedHostWasm []byte

func multiValueFromImportedHostWasmFunctions(ctx context.Context, r wazero.Runtime) (api.Module, error) {
	_ = "STUB: not implemented"
	return *new(api.Module), nil
}
