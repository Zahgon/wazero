//go:build plan9 || aix

package sys

func syscallToErrno(err error) (Errno, bool) { _ = "STUB: not implemented"; return *new(Errno), false }
