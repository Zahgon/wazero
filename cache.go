package wazero

import (
	"context"
	"sync"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/internal/filecache"
	"github.com/tetratelabs/wazero/internal/wasm"
)

type CompilationCache interface{ api.Closer }

func NewCompilationCache() CompilationCache {
	_ = "STUB: not implemented"
	return *new(CompilationCache)
}

func NewCompilationCacheWithDir(dirname string) (CompilationCache, error) {
	_ = "STUB: not implemented"
	return *new(CompilationCache), nil
}

type cache struct {
	engs      [engineKindCount]wasm.Engine
	fileCache filecache.Cache
	initOnces [engineKindCount]sync.Once
}

func (c *cache) initEngine(ek engineKind, ne newEngine, ctx context.Context, features api.CoreFeatures) wasm.Engine {
	_ = "STUB: not implemented"
	return *new(wasm.Engine)
}

func (c *cache) Close(_ context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (c *cache) ensuresFileCache(dir string, wazeroVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func mkdir(dirname string) error { _ = "STUB: not implemented"; return nil }
