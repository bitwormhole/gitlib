package repository

import "github.com/bitwormhole/gitlib/util"

// Core 是仓库的核心目录 ref to the [.git] directory, or not if bare.
type Core interface {
	util.LocalDirectory

	Config() GitConfig
	HEAD() GitHEAD
	Hooks() GitHooks
	Index() GitIndex
	Logs() GitLogs
	Modules() GitModules
	Objects() GitObjects
	Refs() GitRefs
	Worktrees() GitWorktrees
}
