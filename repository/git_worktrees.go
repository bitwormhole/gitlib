package repository

import "github.com/bitwormhole/gitlib/util"

// GitWorktree 是git工作树
type GitWorktree interface {
	util.LocalDirectory

	GetWorkspace() (GitWorkspace, error)
	OpenViewport() (Viewport, error)
}

// GitWorktrees 是git工作树的集合
type GitWorktrees interface {
	util.LocalDirectory

	ListNames() []string
	GetWorktree(name string) (GitWorktree, error)
	CreateWorktree(name string) (GitWorktree, error)
	Default() GitWorktree
}
