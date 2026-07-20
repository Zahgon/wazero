package sys

import (
	"io/fs"
)

const sysParseable = true

func statFromFileInfo(info fs.FileInfo) Stat_t { _ = "STUB: not implemented"; return *new(Stat_t) }
