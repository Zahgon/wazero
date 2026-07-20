package sys

func UnwrapOSError(err error) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func underlyingError(err error) error { _ = "STUB: not implemented"; return nil }
