package sys

type Pflag uint32

const (
	POLLIN Pflag = 1 << iota

	POLLOUT
)

type Pollable interface {
	Poll(flag Pflag, timeoutMillis int32) (ready bool, errno Errno)
}
