package logging

import (
	"context"
	"io"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental"
	"github.com/tetratelabs/wazero/internal/logging"
)

type Writer interface {
	io.Writer
	io.StringWriter
}

type LogScopes = logging.LogScopes

const (
	LogScopeNone = logging.LogScopeNone

	LogScopeClock = logging.LogScopeClock

	LogScopeProc = logging.LogScopeProc

	LogScopeFilesystem = logging.LogScopeFilesystem

	LogScopeMemory = logging.LogScopeMemory

	LogScopePoll = logging.LogScopePoll

	LogScopeRandom = logging.LogScopeRandom

	LogScopeSock = logging.LogScopeSock

	LogScopeAll = logging.LogScopeAll
)

func NewLoggingListenerFactory(w Writer) experimental.FunctionListenerFactory {
	_ = "STUB: not implemented"
	return *new(experimental.FunctionListenerFactory)
}

func NewHostLoggingListenerFactory(w Writer, scopes logging.LogScopes) experimental.FunctionListenerFactory {
	_ = "STUB: not implemented"
	return *new(experimental.FunctionListenerFactory)
}

func toInternalWriter(w Writer) logging.Writer {
	_ = "STUB: not implemented"
	return *new(logging.Writer)
}

type loggingListenerFactory struct {
	w        logging.Writer
	hostOnly bool
	scopes   logging.LogScopes
	stack    logStack
}

type flusher interface {
	Flush() error
}

func (f *loggingListenerFactory) NewFunctionListener(fnd api.FunctionDefinition) experimental.FunctionListener {
	_ = "STUB: not implemented"
	return *new(experimental.FunctionListener)
}

type logStack struct {
	params [][]uint64
}

func (s *logStack) push(params []uint64) { _ = "STUB: not implemented"; return }

func (s *logStack) pop() []uint64 { _ = "STUB: not implemented"; return nil }

func (s *logStack) count() (n int) { _ = "STUB: not implemented"; return 0 }

type loggingListener struct {
	w                         logging.Writer
	beforePrefix, afterPrefix string
	pLoggers                  []logging.ParamLogger
	pSampler                  logging.ParamSampler
	rLoggers                  []logging.ResultLogger
	stack                     *logStack
}

func (l *loggingListener) Before(ctx context.Context, mod api.Module, def api.FunctionDefinition, params []uint64, _ experimental.StackIterator) {
	_ = "STUB: not implemented"
	return
}

func (l *loggingListener) After(ctx context.Context, mod api.Module, def api.FunctionDefinition, results []uint64) {
	_ = "STUB: not implemented"
	return
}

func (l *loggingListener) Abort(ctx context.Context, mod api.Module, def api.FunctionDefinition, _ error) {
	_ = "STUB: not implemented"
	return
}

func (l *loggingListener) logIndented(nestLevel int, prefix string, log func()) {
	_ = "STUB: not implemented"
	return
}

func (l *loggingListener) logParams(ctx context.Context, mod api.Module, params []uint64) {
	_ = "STUB: not implemented"
	return
}

func (l *loggingListener) logResults(ctx context.Context, mod api.Module, params, results []uint64) {
	_ = "STUB: not implemented"
	return
}
