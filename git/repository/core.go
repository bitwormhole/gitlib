package repository

import (
	"github.com/bitwormhole/gitlib/git/repository/modules"
	"github.com/bitwormhole/gitlib/git/repository/objects"
	"github.com/bitwormhole/gitlib/git/repository/refs"
	"github.com/bitwormhole/gitlib/git/repository/worktrees"
	"github.com/bitwormhole/starter/io/fs"
)

// Core 提供git仓库的核心接口
type Core interface {
	GetDirectory() fs.Path

	GetRefs() refs.Directory

	GetObjects() objects.Directory

	GetWorktrees() worktrees.Directory

	GetModules() modules.Directory

	GetConfig() Config
}
