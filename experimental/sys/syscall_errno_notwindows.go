//go:build !windows

package sys

func errorToErrno(err error) Errno { _ = "STUB: not implemented"; return *new(Errno) }
