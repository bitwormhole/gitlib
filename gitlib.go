package gitlib

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// New 初始化git模块
func New(cfg *store.ContextConfiguration) store.Lib {

	if cfg == nil {
		cfg = GetDefaultConfiguration()
	}

	storeContext, err := cfg.Factory.Create(cfg)
	if err != nil {
		panic(err)
	}

	lib := storeContext.Lib
	if lib == nil {
		panic("lib is nil")
	}
	return lib
}

// Bind ...
func Bind(cc context.Context) context.Context {
	return store.Bind(cc)
}
