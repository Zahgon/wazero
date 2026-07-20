package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/tetratelabs/wazero"
)

//go:embed testdata/add.wasm
var addWasm []byte

func main() {

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	compiledWasm, err := r.CompileModule(ctx, addWasm)
	if err != nil {
		log.Panicf("failed to compile Wasm binary: %v", err)
	}

	var wg sync.WaitGroup
	const goroutines = 50
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()

			instance, err := r.InstantiateModule(ctx, compiledWasm, wazero.NewModuleConfig().WithName(""))
			if err != nil {
				log.Panicf("[%d] failed to instantiate %v", i, err)
			}
			defer instance.Close(ctx)

			result, err := instance.ExportedFunction("add").Call(ctx, uint64(i), uint64(i))
			if err != nil {
				log.Panicf("[%d] failed to invoke \"add\": %v", i, err)
			}

			expected := uint64(i * 2)
			if result[0] != expected {
				log.Panicf("expected %d, but got %d", expected, result[0])
			}

			fmt.Println(expected)
		}(i)
	}

	wg.Wait()
}
