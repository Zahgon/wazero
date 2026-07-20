package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

//go:embed testdata/greet.wasm
var greetWasm []byte

func main() {

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	_, err := r.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(logString).Export("log").
		Instantiate(ctx)
	if err != nil {
		log.Panicln(err)
	}

	mod, err := r.Instantiate(ctx, greetWasm)
	if err != nil {
		log.Panicln(err)
	}

	greet := mod.ExportedFunction("greet")
	greeting := mod.ExportedFunction("greeting")
	allocate := mod.ExportedFunction("allocate")
	deallocate := mod.ExportedFunction("deallocate")

	name := os.Args[1]
	nameSize := uint64(len(name))

	results, err := allocate.Call(ctx, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	namePtr := results[0]

	defer deallocate.Call(ctx, namePtr, nameSize)

	if !mod.Memory().Write(uint32(namePtr), []byte(name)) {
		log.Panicf("Memory.Write(%d, %d) out of range of memory size %d",
			namePtr, nameSize, mod.Memory().Size())
	}

	_, err = greet.Call(ctx, namePtr, nameSize)
	if err != nil {
		log.Panicln(err)
	}

	ptrSize, err := greeting.Call(ctx, namePtr, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	greetingPtr := uint32(ptrSize[0] >> 32)
	greetingSize := uint32(ptrSize[0])

	defer func() {
		_, err = deallocate.Call(ctx, uint64(greetingPtr), uint64(greetingSize))
		if err != nil {
			log.Panicln(err)
		}
	}()

	if bytes, ok := mod.Memory().Read(greetingPtr, greetingSize); !ok {
		log.Panicf("Memory.Read(%d, %d) out of range of memory size %d",
			greetingPtr, greetingSize, mod.Memory().Size())
	} else {
		fmt.Println("go >>", string(bytes))
	}
}

func logString(ctx context.Context, m api.Module, offset, byteCount uint32) {
	_ = "STUB: not implemented"
	return
}
