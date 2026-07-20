package sys

const (
	ExitCodeContextCanceled uint32 = 0xffffffff

	ExitCodeDeadlineExceeded uint32 = 0xefffffff
)

type ExitError struct {
	exitCode uint32
}

var exitZero = &ExitError{}

func NewExitError(exitCode uint32) *ExitError { _ = "STUB: not implemented"; return nil }

func (e *ExitError) ExitCode() uint32 { _ = "STUB: not implemented"; return 0 }

func (e *ExitError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ExitError) Is(err error) bool { _ = "STUB: not implemented"; return false }
