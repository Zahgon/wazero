package sysfs

import (
	"github.com/tetratelabs/wazero"
	experimentalsys "github.com/tetratelabs/wazero/experimental/sys"
)

type FSConfig interface {
	WithSysFSMount(fs experimentalsys.FS, guestPath string) wazero.FSConfig
}
