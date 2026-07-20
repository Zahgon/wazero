package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/tetratelabs/wazero/api"
)

//go:embed testdata/greet.wasm
var greetWasm []byte

func main() {
	if err := run(); err != nil {
		log.Panicln(err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func logString(_ context.Context, m api.Module, offset, byteCount uint32) {
	_ = "STUB: not implemented"
	return
}
