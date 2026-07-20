package sysfs

import (
	experimentalsys "github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/sysfs"
)

type AdaptFS = sysfs.AdaptFS

func DirFS(dir string) experimentalsys.FS {
	_ = "STUB: not implemented"
	return *new(experimentalsys.FS)
}

type ReadFS = sysfs.ReadFS
