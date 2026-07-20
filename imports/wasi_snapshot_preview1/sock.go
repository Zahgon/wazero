package wasi_snapshot_preview1

import (
	"context"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
)

var sockAccept = newHostFunc(
	wasip1.SockAcceptName,
	sockAcceptFn,
	[]wasm.ValueType{i32, i32, i32},
	"fd", "flags", "result.fd",
)

func sockAcceptFn(_ context.Context, mod api.Module, params []uint64) (errno sys.Errno) {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

var sockRecv = newHostFunc(
	wasip1.SockRecvName,
	sockRecvFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32, i32},
	"fd", "ri_data", "ri_data_len", "ri_flags", "result.ro_datalen", "result.ro_flags",
)

func sockRecvFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

var sockSend = newHostFunc(
	wasip1.SockSendName,
	sockSendFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32},
	"fd", "si_data", "si_data_len", "si_flags", "result.so_datalen",
)

func sockSendFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}

var sockShutdown = newHostFunc(wasip1.SockShutdownName, sockShutdownFn, []wasm.ValueType{i32, i32}, "fd", "how")

func sockShutdownFn(_ context.Context, mod api.Module, params []uint64) sys.Errno {
	_ = "STUB: not implemented"
	return *new(sys.Errno)
}
