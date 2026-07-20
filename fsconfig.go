package wazero

import (
	"io/fs"

	experimentalsys "github.com/tetratelabs/wazero/experimental/sys"
)

type FSConfig interface {
	WithDirMount(dir, guestPath string) FSConfig

	WithReadOnlyDirMount(dir, guestPath string) FSConfig

	WithFSMount(fs fs.FS, guestPath string) FSConfig
}

type fsConfig struct {
	fs []experimentalsys.FS

	guestPaths []string

	guestPathToFS map[string]int
}

func NewFSConfig() FSConfig { _ = "STUB: not implemented"; return *new(FSConfig) }

func (c *fsConfig) clone() *fsConfig { _ = "STUB: not implemented"; return nil }

func (c *fsConfig) WithDirMount(dir, guestPath string) FSConfig {
	_ = "STUB: not implemented"
	return *new(FSConfig)
}

func (c *fsConfig) WithReadOnlyDirMount(dir, guestPath string) FSConfig {
	_ = "STUB: not implemented"
	return *new(FSConfig)
}

func (c *fsConfig) WithFSMount(fs fs.FS, guestPath string) FSConfig {
	_ = "STUB: not implemented"
	return *new(FSConfig)
}

func (c *fsConfig) WithSysFSMount(fs experimentalsys.FS, guestPath string) FSConfig {
	_ = "STUB: not implemented"
	return *new(FSConfig)
}

func (c *fsConfig) preopens() ([]experimentalsys.FS, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
