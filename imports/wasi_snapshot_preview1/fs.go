package wasi_snapshot_preview1

import (
	"context"
	"io/fs"
	"math"

	"github.com/tetratelabs/wazero/api"
	experimentalsys "github.com/tetratelabs/wazero/experimental/sys"
	"github.com/tetratelabs/wazero/internal/sys"
	"github.com/tetratelabs/wazero/internal/wasip1"
	"github.com/tetratelabs/wazero/internal/wasm"
	sysapi "github.com/tetratelabs/wazero/sys"
)

var fdAdvise = newHostFunc(
	wasip1.FdAdviseName, fdAdviseFn,
	[]wasm.ValueType{i32, i64, i64, i32},
	"fd", "offset", "len", "advice",
)

func fdAdviseFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdAllocate = newHostFunc(
	wasip1.FdAllocateName, fdAllocateFn,
	[]wasm.ValueType{i32, i64, i64},
	"fd", "offset", "len",
)

func fdAllocateFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdClose = newHostFunc(wasip1.FdCloseName, fdCloseFn, []wasm.ValueType{i32}, "fd")

func fdCloseFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdDatasync = newHostFunc(wasip1.FdDatasyncName, fdDatasyncFn, []wasm.ValueType{i32}, "fd")

func fdDatasyncFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdFdstatGet = newHostFunc(wasip1.FdFdstatGetName, fdFdstatGetFn, []wasm.ValueType{i32, i32}, "fd", "result.stat")

func fdFdstatGetFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func isPreopenedStdio(fd int32, f *sys.FileEntry) bool { _ = "STUB: not implemented"; return false }

const fileRightsBase = wasip1.RIGHT_FD_DATASYNC |
	wasip1.RIGHT_FD_READ |
	wasip1.RIGHT_FD_SEEK |
	wasip1.RIGHT_FDSTAT_SET_FLAGS |
	wasip1.RIGHT_FD_SYNC |
	wasip1.RIGHT_FD_TELL |
	wasip1.RIGHT_FD_WRITE |
	wasip1.RIGHT_FD_ADVISE |
	wasip1.RIGHT_FD_ALLOCATE |
	wasip1.RIGHT_FD_FILESTAT_GET |
	wasip1.RIGHT_FD_FILESTAT_SET_SIZE |
	wasip1.RIGHT_FD_FILESTAT_SET_TIMES |
	wasip1.RIGHT_POLL_FD_READWRITE

const dirRightsBase = wasip1.RIGHT_FD_DATASYNC |
	wasip1.RIGHT_FDSTAT_SET_FLAGS |
	wasip1.RIGHT_FD_SYNC |
	wasip1.RIGHT_PATH_CREATE_DIRECTORY |
	wasip1.RIGHT_PATH_CREATE_FILE |
	wasip1.RIGHT_PATH_LINK_SOURCE |
	wasip1.RIGHT_PATH_LINK_TARGET |
	wasip1.RIGHT_PATH_OPEN |
	wasip1.RIGHT_FD_READDIR |
	wasip1.RIGHT_PATH_READLINK |
	wasip1.RIGHT_PATH_RENAME_SOURCE |
	wasip1.RIGHT_PATH_RENAME_TARGET |
	wasip1.RIGHT_PATH_FILESTAT_GET |
	wasip1.RIGHT_PATH_FILESTAT_SET_SIZE |
	wasip1.RIGHT_PATH_FILESTAT_SET_TIMES |
	wasip1.RIGHT_FD_FILESTAT_GET |
	wasip1.RIGHT_FD_FILESTAT_SET_TIMES |
	wasip1.RIGHT_PATH_SYMLINK |
	wasip1.RIGHT_PATH_REMOVE_DIRECTORY |
	wasip1.RIGHT_PATH_UNLINK_FILE

func writeFdstat(buf []byte, fileType uint8, fdflags uint16, fsRightsBase, fsRightsInheriting uint32) {
	_ = "STUB: not implemented"
	return
}

var fdFdstatSetFlags = newHostFunc(wasip1.FdFdstatSetFlagsName, fdFdstatSetFlagsFn, []wasm.ValueType{i32, i32}, "fd", "flags")

func fdFdstatSetFlagsFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdFdstatSetRights = stubFunction(
	wasip1.FdFdstatSetRightsName,
	[]wasm.ValueType{i32, i64, i64},
	"fd", "fs_rights_base", "fs_rights_inheriting",
)

var fdFilestatGet = newHostFunc(wasip1.FdFilestatGetName, fdFilestatGetFn, []wasm.ValueType{i32, i32}, "fd", "result.filestat")

func fdFilestatGetFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func fdFilestatGetFunc(mod api.Module, fd int32, resultBuf uint32) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func getExtendedWasiFiletype(file experimentalsys.File, fm fs.FileMode) (ftype uint8) {
	_ = "STUB: not implemented"
	return 0
}

func getWasiFiletype(fm fs.FileMode) uint8 { _ = "STUB: not implemented"; return 0 }

func writeFilestat(buf []byte, st *sysapi.Stat_t, ftype uint8) (errno experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdFilestatSetSize = newHostFunc(wasip1.FdFilestatSetSizeName, fdFilestatSetSizeFn, []wasm.ValueType{i32, i64}, "fd", "size")

func fdFilestatSetSizeFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdFilestatSetTimes = newHostFunc(
	wasip1.FdFilestatSetTimesName, fdFilestatSetTimesFn,
	[]wasm.ValueType{i32, i64, i64, i32},
	"fd", "atim", "mtim", "fst_flags",
)

func fdFilestatSetTimesFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func toTimes(walltime func() int64, atim, mtim int64, fstFlags uint16) (int64, int64, experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return 0, 0, *new(experimentalsys.Errno)
}

var fdPread = newHostFunc(
	wasip1.FdPreadName, fdPreadFn,
	[]wasm.ValueType{i32, i32, i32, i64, i32},
	"fd", "iovs", "iovs_len", "offset", "result.nread",
)

func fdPreadFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdPrestatGet = newHostFunc(wasip1.FdPrestatGetName, fdPrestatGetFn, []wasm.ValueType{i32, i32}, "fd", "result.prestat")

func fdPrestatGetFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdPrestatDirName = newHostFunc(
	wasip1.FdPrestatDirNameName, fdPrestatDirNameFn,
	[]wasm.ValueType{i32, i32, i32},
	"fd", "result.path", "result.path_len",
)

func fdPrestatDirNameFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdPwrite = newHostFunc(
	wasip1.FdPwriteName, fdPwriteFn,
	[]wasm.ValueType{i32, i32, i32, i64, i32},
	"fd", "iovs", "iovs_len", "offset", "result.nwritten",
)

func fdPwriteFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdRead = newHostFunc(
	wasip1.FdReadName, fdReadFn,
	[]wasm.ValueType{i32, i32, i32, i32},
	"fd", "iovs", "iovs_len", "result.nread",
)

type preader struct {
	f      experimentalsys.File
	offset int64
}

func (w *preader) Read(buf []byte) (n int, errno experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return 0, *new(experimentalsys.Errno)
}

func fdReadFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func fdReadOrPread(mod api.Module, params []uint64, isPread bool) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func readv(mem api.Memory, iovs uint32, iovsCount uint32, reader func(buf []byte) (nread int, errno experimentalsys.Errno)) (uint32, experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return 0, *new(experimentalsys.Errno)
}

var fdReaddir = newHostFunc(
	wasip1.FdReaddirName, fdReaddirFn,
	[]wasm.ValueType{i32, i32, i32, i64, i32},
	"fd", "buf", "buf_len", "cookie", "result.bufused",
)

func fdReaddirFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

const largestDirent = int64(math.MaxUint32 - wasip1.DirentSize)

func maxDirents(dirents []experimentalsys.Dirent, bufLen uint32) (bufToWrite uint32, direntCount int, truncatedLen uint32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func writeDirents(buf []byte, dirents []experimentalsys.Dirent, d_next uint64, direntCount int, truncatedLen uint32) {
	_ = "STUB: not implemented"
	return
}

func writeDirent(buf []byte, dNext uint64, ino sysapi.Inode, dNamlen uint32, dType fs.FileMode) {
	_ = "STUB: not implemented"
	return
}

func direntCache(fsc *sys.FSContext, fd int32) (*sys.DirentCache, experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return nil, *new(experimentalsys.Errno)
}

var fdRenumber = newHostFunc(wasip1.FdRenumberName, fdRenumberFn, []wasm.ValueType{i32, i32}, "fd", "to")

func fdRenumberFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdSeek = newHostFunc(
	wasip1.FdSeekName, fdSeekFn,
	[]wasm.ValueType{i32, i64, i32, i32},
	"fd", "offset", "whence", "result.newoffset",
)

func fdSeekFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdSync = newHostFunc(wasip1.FdSyncName, fdSyncFn, []wasm.ValueType{i32}, "fd")

func fdSyncFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdTell = newHostFunc(wasip1.FdTellName, fdTellFn, []wasm.ValueType{i32, i32}, "fd", "result.offset")

func fdTellFn(ctx context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var fdWrite = newHostFunc(
	wasip1.FdWriteName, fdWriteFn,
	[]wasm.ValueType{i32, i32, i32, i32},
	"fd", "iovs", "iovs_len", "result.nwritten",
)

func fdWriteFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

type pwriter struct {
	f      experimentalsys.File
	offset int64
}

func (w *pwriter) Write(buf []byte) (n int, errno experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return 0, *new(experimentalsys.Errno)
}

func fdWriteOrPwrite(mod api.Module, params []uint64, isPwrite bool) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func writev(mem api.Memory, iovs uint32, iovsCount uint32, writer func(buf []byte) (n int, errno experimentalsys.Errno)) (uint32, experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return 0, *new(experimentalsys.Errno)
}

var pathCreateDirectory = newHostFunc(
	wasip1.PathCreateDirectoryName, pathCreateDirectoryFn,
	[]wasm.ValueType{i32, i32, i32},
	"fd", "path", "path_len",
)

func pathCreateDirectoryFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathFilestatGet = newHostFunc(
	wasip1.PathFilestatGetName, pathFilestatGetFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32},
	"fd", "flags", "path", "path_len", "result.filestat",
)

func pathFilestatGetFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathFilestatSetTimes = newHostFunc(
	wasip1.PathFilestatSetTimesName, pathFilestatSetTimesFn,
	[]wasm.ValueType{i32, i32, i32, i32, i64, i64, i32},
	"fd", "flags", "path", "path_len", "atim", "mtim", "fst_flags",
)

func pathFilestatSetTimesFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathLink = newHostFunc(
	wasip1.PathLinkName, pathLinkFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32, i32, i32},
	"old_fd", "old_flags", "old_path", "old_path_len", "new_fd", "new_path", "new_path_len",
)

func pathLinkFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathOpen = newHostFunc(
	wasip1.PathOpenName, pathOpenFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32, i64, i64, i32, i32},
	"fd", "dirflags", "path", "path_len", "oflags", "fs_rights_base", "fs_rights_inheriting", "fdflags", "result.opened_fd",
)

func pathOpenFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

func atPath(fsc *sys.FSContext, mem api.Memory, fd int32, p, pathLen uint32) (experimentalsys.FS, string, experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return *new(experimentalsys.FS), "", *new(experimentalsys.Errno)
}

func preopenPath(fsc *sys.FSContext, fd int32) (string, experimentalsys.Errno) {
	_ = "STUB: not implemented"
	return "", *new(experimentalsys.Errno)
}

func openFlags(dirflags, oflags, fdflags uint16, rights uint32) (openFlags experimentalsys.Oflag) {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Oflag)
}

var pathReadlink = newHostFunc(
	wasip1.PathReadlinkName, pathReadlinkFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32, i32},
	"fd", "path", "path_len", "buf", "buf_len", "result.bufused",
)

func pathReadlinkFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathRemoveDirectory = newHostFunc(
	wasip1.PathRemoveDirectoryName, pathRemoveDirectoryFn,
	[]wasm.ValueType{i32, i32, i32},
	"fd", "path", "path_len",
)

func pathRemoveDirectoryFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathRename = newHostFunc(
	wasip1.PathRenameName, pathRenameFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32, i32},
	"fd", "old_path", "old_path_len", "new_fd", "new_path", "new_path_len",
)

func pathRenameFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathSymlink = newHostFunc(
	wasip1.PathSymlinkName, pathSymlinkFn,
	[]wasm.ValueType{i32, i32, i32, i32, i32},
	"old_path", "old_path_len", "fd", "new_path", "new_path_len",
)

func pathSymlinkFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}

var pathUnlinkFile = newHostFunc(
	wasip1.PathUnlinkFileName, pathUnlinkFileFn,
	[]wasm.ValueType{i32, i32, i32},
	"fd", "path", "path_len",
)

func pathUnlinkFileFn(_ context.Context, mod api.Module, params []uint64) experimentalsys.Errno {
	_ = "STUB: not implemented"
	return *new(experimentalsys.Errno)
}
