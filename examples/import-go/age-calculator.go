package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tetratelabs/wazero"
)

//go:embed testdata/age_calculator.wasm
var ageCalculatorWasm []byte

func main() {

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	_, err := r.NewHostModuleBuilder("env").
		NewFunctionBuilder().
		WithFunc(func(v uint32) {
			fmt.Println("log_i32 >>", v)
		}).
		Export("log_i32").
		NewFunctionBuilder().
		WithFunc(func() uint32 {
			if envYear, err := strconv.ParseUint(os.Getenv("CURRENT_YEAR"), 10, 64); err == nil {
				return uint32(envYear)
			}
			return uint32(time.Now().Year())
		}).
		Export("current_year").
		Instantiate(ctx)
	if err != nil {
		log.Panicln(err)
	}

	ageCalculator, err := r.Instantiate(ctx, ageCalculatorWasm)
	if err != nil {
		log.Panicln(err)
	}

	birthYear, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Panicf("invalid arg %v: %v", os.Args[1], err)
	}

	results, err := ageCalculator.ExportedFunction("get_age").Call(ctx, birthYear)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("println >>", results[0])

	_, err = ageCalculator.ExportedFunction("log_age").Call(ctx, birthYear)
	if err != nil {
		log.Panicln(err)
	}
}
