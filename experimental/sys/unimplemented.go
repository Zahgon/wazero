package sys

import (
	"io/fs"

	"github.com/tetratelabs/wazero/sys"
)

type UnimplementedFS struct{}

func (UnimplementedFS) OpenFile(path string, flag Oflag, perm fs.FileMode) (File, Errno) {
	_ = "STUB: not implemented"
	return *new(File), *new(Errno)
}

func (UnimplementedFS) Lstat(path string) (sys.Stat_t, Errno) {
	_ = "STUB: not implemented"
	return *new(sys.Stat_t), *new(Errno)
}

func (UnimplementedFS) Stat(path string) (sys.Stat_t, Errno) {
	_ = "STUB: not implemented"
	return *new(sys.Stat_t), *new(Errno)
}

func (UnimplementedFS) Readlink(path string) (string, Errno) {
	_ = "STUB: not implemented"
	return "", *new(Errno)
}

func (UnimplementedFS) Mkdir(path string, perm fs.FileMode) Errno {
	_ = "STUB: not implemented"
	return *new(Errno)
}

func (UnimplementedFS) Chmod(path string, perm fs.FileMode) Errno {
	_ = "STUB: not implemented"
	return *new(Errno)
}

func (UnimplementedFS) Rename(from, to string) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFS) Rmdir(path string) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFS) Link(_, _ string) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFS) Symlink(_, _ string) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFS) Unlink(path string) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFS) Utimens(path string, atim, mtim int64) Errno {
	_ = "STUB: not implemented"
	return *new(Errno)
}

type UnimplementedFile struct{}

func (UnimplementedFile) Dev() (uint64, Errno) { _ = "STUB: not implemented"; return 0, *new(Errno) }

func (UnimplementedFile) Ino() (sys.Inode, Errno) {
	_ = "STUB: not implemented"
	return *new(sys.Inode), *new(Errno)
}

func (UnimplementedFile) IsDir() (bool, Errno) {
	_ = "STUB: not implemented"
	return false, *new(Errno)
}

func (UnimplementedFile) IsAppend() bool { _ = "STUB: not implemented"; return false }

func (UnimplementedFile) SetAppend(bool) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFile) Stat() (sys.Stat_t, Errno) {
	_ = "STUB: not implemented"
	return *new(sys.Stat_t), *new(Errno)
}

func (UnimplementedFile) Read([]byte) (int, Errno) {
	_ = "STUB: not implemented"
	return 0, *new(Errno)
}

func (UnimplementedFile) Pread([]byte, int64) (int, Errno) {
	_ = "STUB: not implemented"
	return 0, *new(Errno)
}

func (UnimplementedFile) Seek(int64, int) (int64, Errno) {
	_ = "STUB: not implemented"
	return 0, *new(Errno)
}

func (UnimplementedFile) Readdir(int) (dirents []Dirent, errno Errno) {
	_ = "STUB: not implemented"
	return nil, *new(Errno)
}

func (UnimplementedFile) Write([]byte) (int, Errno) {
	_ = "STUB: not implemented"
	return 0, *new(Errno)
}

func (UnimplementedFile) Pwrite([]byte, int64) (int, Errno) {
	_ = "STUB: not implemented"
	return 0, *new(Errno)
}

func (UnimplementedFile) Truncate(int64) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFile) Sync() Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFile) Datasync() Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFile) Utimens(int64, int64) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (UnimplementedFile) Close() (errno Errno) { _ = "STUB: not implemented"; return *new(Errno) }
