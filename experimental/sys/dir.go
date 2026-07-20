package sys

import (
	"io/fs"

	"github.com/tetratelabs/wazero/sys"
)

type FileType = fs.FileMode

type Dirent struct {
	Ino sys.Inode

	Name string

	Type fs.FileMode
}

func (d *Dirent) String() string { _ = "STUB: not implemented"; return "" }

func (d *Dirent) IsDir() bool { _ = "STUB: not implemented"; return false }

type DirFile struct{}

func (DirFile) IsAppend() bool { _ = "STUB: not implemented"; return false }

func (DirFile) SetAppend(bool) Errno { _ = "STUB: not implemented"; return *new(Errno) }

func (DirFile) IsDir() (bool, Errno) { _ = "STUB: not implemented"; return false, *new(Errno) }

func (DirFile) Read([]byte) (int, Errno) { _ = "STUB: not implemented"; return 0, *new(Errno) }

func (DirFile) Pread([]byte, int64) (int, Errno) { _ = "STUB: not implemented"; return 0, *new(Errno) }

func (DirFile) Write([]byte) (int, Errno) { _ = "STUB: not implemented"; return 0, *new(Errno) }

func (DirFile) Pwrite([]byte, int64) (int, Errno) { _ = "STUB: not implemented"; return 0, *new(Errno) }

func (DirFile) Truncate(int64) Errno { _ = "STUB: not implemented"; return *new(Errno) }
