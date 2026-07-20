package sys

import "github.com/tetratelabs/wazero/sys"

type File interface {
	Dev() (uint64, Errno)

	Ino() (sys.Inode, Errno)

	IsDir() (bool, Errno)

	IsAppend() bool

	SetAppend(enable bool) Errno

	Stat() (sys.Stat_t, Errno)

	Read(buf []byte) (n int, errno Errno)

	Pread(buf []byte, off int64) (n int, errno Errno)

	Seek(offset int64, whence int) (newOffset int64, errno Errno)

	Readdir(n int) (dirents []Dirent, errno Errno)

	Write(buf []byte) (n int, errno Errno)

	Pwrite(buf []byte, off int64) (n int, errno Errno)

	Truncate(size int64) Errno

	Sync() Errno

	Datasync() Errno

	Utimens(atim, mtim int64) Errno

	Close() Errno
}

type PollableFile interface {
	File
	Pollable

	IsNonblock() bool

	SetNonblock(enable bool) Errno
}
