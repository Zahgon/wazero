package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/assemblyscript"
)

//go:embed testdata/index.wasm
var asWasm []byte

func main() {

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	_, err := assemblyscript.Instantiate(ctx, r)
	if err != nil {
		log.Panicln(err)
	}

	mod, err := r.InstantiateWithConfig(ctx, asWasm,

		wazero.NewModuleConfig().WithStdout(os.Stdout).WithStderr(os.Stderr))
	if err != nil {
		log.Panicln(err)
	}

	helloWorld := mod.ExportedFunction("hello_world")
	goodbyeWorld := mod.ExportedFunction("goodbye_world")

	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Panicln(err)
	}

	results, err := helloWorld.Call(ctx, uint64(num))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("hello_world returned: %v", results[0])

	if _, err = goodbyeWorld.Call(ctx); err == nil {
		log.Panicln("goodbye_world did not fail")
	}
}
