package experimental

import (
	"context"

	"github.com/tetratelabs/wazero/api"
)

type StackIterator interface {
	Next() bool

	Function() InternalFunction

	ProgramCounter() ProgramCounter
}

func WithFunctionListenerFactory(ctx context.Context, factory FunctionListenerFactory) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type FunctionListenerFactory interface {
	NewFunctionListener(api.FunctionDefinition) FunctionListener
}

type FunctionListener interface {
	Before(ctx context.Context, mod api.Module, def api.FunctionDefinition, params []uint64, stackIterator StackIterator)

	After(ctx context.Context, mod api.Module, def api.FunctionDefinition, results []uint64)

	Abort(ctx context.Context, mod api.Module, def api.FunctionDefinition, err error)
}

type FunctionListenerFunc func(context.Context, api.Module, api.FunctionDefinition, []uint64, StackIterator)

func (f FunctionListenerFunc) Before(ctx context.Context, mod api.Module, def api.FunctionDefinition, params []uint64, stackIterator StackIterator) {
	_ = "STUB: not implemented"
	return
}

func (f FunctionListenerFunc) After(context.Context, api.Module, api.FunctionDefinition, []uint64) {
	_ = "STUB: not implemented"
	return
}

func (f FunctionListenerFunc) Abort(context.Context, api.Module, api.FunctionDefinition, error) {
	_ = "STUB: not implemented"
	return
}

type FunctionListenerFactoryFunc func(api.FunctionDefinition) FunctionListener

func (f FunctionListenerFactoryFunc) NewFunctionListener(def api.FunctionDefinition) FunctionListener {
	_ = "STUB: not implemented"
	return *new(FunctionListener)
}

func MultiFunctionListenerFactory(factories ...FunctionListenerFactory) FunctionListenerFactory {
	_ = "STUB: not implemented"
	return *new(FunctionListenerFactory)
}

type multiFunctionListenerFactory []FunctionListenerFactory

func (multi multiFunctionListenerFactory) NewFunctionListener(def api.FunctionDefinition) FunctionListener {
	_ = "STUB: not implemented"
	return *new(FunctionListener)
}

type multiFunctionListener struct {
	lstns []FunctionListener
	stack stackIterator
}

func (multi *multiFunctionListener) Before(ctx context.Context, mod api.Module, def api.FunctionDefinition, params []uint64, si StackIterator) {
	_ = "STUB: not implemented"
	return
}

func (multi *multiFunctionListener) After(ctx context.Context, mod api.Module, def api.FunctionDefinition, results []uint64) {
	_ = "STUB: not implemented"
	return
}

func (multi *multiFunctionListener) Abort(ctx context.Context, mod api.Module, def api.FunctionDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

type stackIterator struct {
	base  StackIterator
	index int
	pcs   []uint64
	fns   []InternalFunction
}

func (si *stackIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (si *stackIterator) ProgramCounter() ProgramCounter {
	_ = "STUB: not implemented"
	return *new(ProgramCounter)
}

func (si *stackIterator) Function() InternalFunction {
	_ = "STUB: not implemented"
	return *new(InternalFunction)
}

type StackFrame struct {
	Function     api.Function
	Params       []uint64
	Results      []uint64
	PC           uint64
	SourceOffset uint64
}

type internalFunction struct {
	definition   api.FunctionDefinition
	sourceOffset uint64
}

func (f internalFunction) Definition() api.FunctionDefinition {
	_ = "STUB: not implemented"
	return *new(api.FunctionDefinition)
}

func (f internalFunction) SourceOffsetForPC(pc ProgramCounter) uint64 {
	_ = "STUB: not implemented"
	return 0
}

type stackFrameIterator struct {
	index int
	stack []StackFrame
	fndef []api.FunctionDefinition
}

func (si *stackFrameIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (si *stackFrameIterator) Function() InternalFunction {
	_ = "STUB: not implemented"
	return *new(InternalFunction)
}

func (si *stackFrameIterator) ProgramCounter() ProgramCounter {
	_ = "STUB: not implemented"
	return *new(ProgramCounter)
}

func NewStackIterator(stack ...StackFrame) StackIterator {
	_ = "STUB: not implemented"
	return *new(StackIterator)
}

func BenchmarkFunctionListener(n int, module api.Module, stack []StackFrame, listener FunctionListener) {
	_ = "STUB: not implemented"
	return
}
