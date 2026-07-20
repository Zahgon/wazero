package sys

import "io/fs"

type Inode = uint64

type EpochNanos = int64

type Stat_t struct {
	Dev uint64

	Ino Inode

	Mode fs.FileMode

	Nlink uint64

	Size int64

	Atim EpochNanos

	Mtim EpochNanos

	Ctim EpochNanos
}

func NewStat_t(info fs.FileInfo) Stat_t { _ = "STUB: not implemented"; return *new(Stat_t) }

func defaultStatFromFileInfo(info fs.FileInfo) Stat_t {
	_ = "STUB: not implemented"
	return *new(Stat_t)
}
