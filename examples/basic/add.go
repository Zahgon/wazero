package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed testdata/add.wasm
var addWasm []byte

func main() {

	flag.Parse()

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.InstantiateWithConfig(ctx, addWasm, wazero.NewModuleConfig().WithStartFunctions("_initialize"))
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	x, y, err := readTwoArgs(flag.Arg(0), flag.Arg(1))
	if err != nil {
		log.Panicf("failed to read arguments: %v", err)
	}

	add := mod.ExportedFunction("add")
	results, err := add.Call(ctx, x, y)
	if err != nil {
		log.Panicf("failed to call add: %v", err)
	}

	fmt.Printf("%d + %d = %d\n", x, y, results[0])
}

func readTwoArgs(xs, ys string) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
