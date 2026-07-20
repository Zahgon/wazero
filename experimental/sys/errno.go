package sys

type Errno uint16

const (
	EACCES Errno = iota + 1
	EAGAIN
	EBADF
	EEXIST
	EFAULT
	EINTR
	EINVAL
	EIO
	EISDIR
	ELOOP
	ENAMETOOLONG
	ENOENT
	ENOSYS
	ENOTDIR
	ERANGE
	ENOTEMPTY
	ENOTSOCK
	ENOTSUP
	EPERM
	EROFS
)

func (e Errno) Error() string { _ = "STUB: not implemented"; return "" }
