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

//go:embed testdata/counter.wasm
var counterWasm []byte

func main() {

	ctx := context.Background()

	cacheDir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Panicln(err)
	}
	defer os.RemoveAll(cacheDir)

	cache, err := wazero.NewCompilationCacheWithDir(cacheDir)
	if err != nil {
		log.Panicln(err)
	}
	defer cache.Close(ctx)

	runtimeConfig := wazero.NewRuntimeConfig().WithCompilationCache(cache)

	runtimeFoo := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)
	runtimeBar := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)

	m1 := instantiateWithEnv(ctx, runtimeFoo)
	m2 := instantiateWithEnv(ctx, runtimeBar)

	for i := 0; i < 2; i++ {
		fmt.Printf("m1 counter=%d\n", counterGet(ctx, m1))
		fmt.Printf("m2 counter=%d\n", counterGet(ctx, m2))
	}
}

func counterGet(ctx context.Context, mod api.Module) uint64 { _ = "STUB: not implemented"; return 0 }

type counter struct {
	counter uint32
}

func (e *counter) getAndIncrement() (ret uint32) { _ = "STUB: not implemented"; return 0 }

func instantiateWithEnv(ctx context.Context, r wazero.Runtime) api.Module {
	_ = "STUB: not implemented"
	return *new(api.Module)
}
