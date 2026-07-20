package sys

import (
	"io/fs"

	"github.com/tetratelabs/wazero/sys"
)

type FS interface {
	OpenFile(path string, flag Oflag, perm fs.FileMode) (File, Errno)

	Lstat(path string) (sys.Stat_t, Errno)

	Stat(path string) (sys.Stat_t, Errno)

	Mkdir(path string, perm fs.FileMode) Errno

	Chmod(path string, perm fs.FileMode) Errno

	Rename(from, to string) Errno

	Rmdir(path string) Errno

	Unlink(path string) Errno

	Link(oldPath, newPath string) Errno

	Symlink(oldPath, linkName string) Errno

	Readlink(path string) (string, Errno)

	Utimens(path string, atim, mtim int64) Errno
}
