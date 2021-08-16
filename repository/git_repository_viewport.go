package repository

import "io"

// Viewport 仓库的一个视图
type Viewport interface {
	io.Closer

	GetWorkspace() (GitWorkspace, error)
	GetSubmodule() (GitSubmodule, error)
	GetWorktree() (GitWorktree, error)
	GetCore() Core
	GetPWD() GitPWD

	Objects() GitObjects
	Refs() GitRefs
	Config() GitConfig
	Index() GitIndex
	HEAD() GitHEAD
}
