package repository

import "github.com/bitwormhole/gitlib/util"

// GitSubmodule 表示一个git子模块
type GitSubmodule interface {
	util.LocalDirectory

	OpenViewport() (Viewport, error)
}

// GitModules 表示git的模块集合
type GitModules interface {
	util.LocalDirectory

	ListNames() []string
	GetSubmodule(name string) (GitSubmodule, error)
	CreateSubmodule(name string) (GitSubmodule, error)
}
