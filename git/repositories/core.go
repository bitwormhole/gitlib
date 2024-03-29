package repositories

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/afs"
)

// Core 表示仓库对象的核心
type Core struct {
	WD afs.Path // 工作目录

	Context *Context

	Config ConfigChain // 配置链：repo -> user -> system

	Head HEAD

	Index Index

	Layout Layout

	Objects Objects

	Refs Refs

	Repository Repository

	Workspace Workspace // aka the default worktree

	Worktrees Worktrees

	Submodules Submodules

	SessionFactory SessionFactory

	Digest      git.Digest      // default="SHA-1"
	Compression git.Compression // default="DEFLATE"
	PathMapping git.PathMapping // default="xx/xxxx"
}

// CoreConfigurer 是用来配置 repository.Core 的组件
// [inject:".git-core-configurer"]
type CoreConfigurer interface {
	Configure(c *Core) error
}
