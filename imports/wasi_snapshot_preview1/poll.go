package wasi_snapshot_preview1

import (
	"context"
	"time"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var pollOneoff = newHostFunc(
	wasip1.PollOneoffName, pollOneoffFn,
	[]wasm.ValueType{i32, i32, i32, i32},
	"in", "out", "nsubscriptions", "result.nevents",
)

type event struct {
	eventType byte
	userData  []byte
	errno     wasip1.Errno
}

func pollOneoffFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

func processClockEvent(inBuf []byte) (time.Duration, sys.Errno) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(sys.Errno)
}

func isNonblock(f sys.File) bool { _ = "STUB: not implemented"; return false }

func writeEvent(outBuf []byte, evt *event) { _ = "STUB: not implemented"; return }
