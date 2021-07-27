package repository

import "github.com/bitwormhole/gitlib/util"

// ref to the [.git] directory, or not if bare.
type GitRepository interface {
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
