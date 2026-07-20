package main

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/logging"
	"github.com/tetratelabs/wazero/experimental/sock"
)

func main() {
	os.Exit(doMain(os.Stdout, os.Stderr))
}

func doMain(stdOut io.Writer, stdErr logging.Writer) int { _ = "STUB: not implemented"; return 0 }

func doCompile(args []string, stdErr io.Writer) int { _ = "STUB: not implemented"; return 0 }

func doRun(args []string, stdOut io.Writer, stdErr logging.Writer) int {
	_ = "STUB: not implemented"
	return 0
}

func validateMounts(mounts sliceFlag, stdErr logging.Writer) (rc int, rootPath string, config wazero.FSConfig) {
	_ = "STUB: not implemented"
	return 0, "", *new(wazero.FSConfig)
}

func validateListens(listens sliceFlag, stdErr logging.Writer) (rc int, config sock.Config) {
	_ = "STUB: not implemented"
	return 0, *new(sock.Config)
}

const (
	modeDefault importMode = iota
	modeWasi
	modeWasiUnstable
)

type importMode uint

func detectImports(imports []api.FunctionDefinition) importMode {
	_ = "STUB: not implemented"
	return *new(importMode)
}

func maybeHostLogging(ctx context.Context, scopes logging.LogScopes, stdErr logging.Writer) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func cacheDirFlag(flags *flag.FlagSet) *string { _ = "STUB: not implemented"; return nil }

func workersFlag(flags *flag.FlagSet) *int { _ = "STUB: not implemented"; return nil }

func maybeUseCacheDir(cacheDir *string, stdErr io.Writer) (int, wazero.CompilationCache) {
	_ = "STUB: not implemented"
	return 0, *new(wazero.CompilationCache)
}

func printUsage(stdErr io.Writer) { _ = "STUB: not implemented"; return }

func printCompileUsage(stdErr io.Writer, flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func printRunUsage(stdErr io.Writer, flags *flag.FlagSet) { _ = "STUB: not implemented"; return }

func startCPUProfile(stdErr io.Writer, path string) (stopCPUProfile func()) {
	_ = "STUB: not implemented"
	return nil
}

func writeHeapProfile(stdErr io.Writer, path string) { _ = "STUB: not implemented"; return }

type sliceFlag []string

func (f *sliceFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *sliceFlag) Set(s string) error { _ = "STUB: not implemented"; return nil }

type logScopesFlag logging.LogScopes

func (f *logScopesFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *logScopesFlag) Set(input string) error { _ = "STUB: not implemented"; return nil }
