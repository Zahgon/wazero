package sys

type Oflag uint32

const (
	O_RDONLY Oflag = iota

	O_RDWR

	O_WRONLY

	O_APPEND Oflag = 1 << iota

	O_CREAT

	O_DIRECTORY

	O_DSYNC

	O_EXCL

	O_NOFOLLOW

	O_NONBLOCK

	O_RSYNC

	O_SYNC

	O_TRUNC
)
