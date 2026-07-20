//go:build !(linux || darwin || freebsd || netbsd || openbsd || dragonfly || solaris || windows)

package sys

import "io/fs"

const sysParseable = false

func statFromFileInfo(info fs.FileInfo) Stat_t { _ = "STUB: not implemented"; return *new(Stat_t) }
