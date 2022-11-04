package store

import (
	"bitwormhole.com/starter/afs"
)

// Core 表示仓库对象的核心
type Core struct {
	WD afs.Path // 工作目录

	Context *Context

	Config ConfigChain // 配置链：repo -> user -> system

	Head HEAD

	Index Index

	Layout RepositoryLayout

	Objects Objects

	Refs Refs

	Repository RepositoryProfile

	Workspace Workspace

	SessionFactory SessionFactory

	Digest      Digest      // default="SHA-1"
	Compression Compression // default="DEFLATE"
	PathMapping PathMapping // default="xx/xxxx"
}

// CoreConfigurer 是用来配置 repository.Core 的组件
// [inject:".git-core-configurer"]
type CoreConfigurer interface {
	Configure(c *Core) error
}
