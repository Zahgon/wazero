package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

//go:embed testdata/test.txt
var catFS embed.FS

//go:embed testdata/cargo-wasi/cat.wasm
var catWasmCargoWasi []byte

//go:embed testdata/tinygo/cat.wasm
var catWasmTinyGo []byte

//go:embed testdata/zig/cat.wasm
var catWasmZig []byte

//go:embed testdata/zig-cc/cat.wasm
var catWasmZigCc []byte

func main() {

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	rooted, err := fs.Sub(catFS, "testdata")
	if err != nil {
		log.Panicln(err)
	}

	config := wazero.NewModuleConfig().
		WithStdout(os.Stdout).WithStderr(os.Stderr).WithFS(rooted)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	var catWasm []byte
	toolchain := os.Getenv("TOOLCHAIN")
	switch toolchain {
	case "":
		fallthrough
	case "cargo-wasi":
		catWasm = catWasmCargoWasi
	case "tinygo":
		catWasm = catWasmTinyGo
	case "zig":
		catWasm = catWasmZig
	case "zig-cc":
		catWasm = catWasmZigCc
	default:
		log.Panicln("unknown toolchain", toolchain)
	}

	if _, err = r.InstantiateWithConfig(ctx, catWasm, config.WithArgs("wasi", os.Args[1])); err != nil {

		if exitErr, ok := err.(*sys.ExitError); ok && exitErr.ExitCode() != 0 {
			fmt.Fprintf(os.Stderr, "exit_code: %d\n", exitErr.ExitCode())
		} else if !ok {
			log.Panicln(err)
		}
	}
}
