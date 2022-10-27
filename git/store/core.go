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
}

// CoreConfigurer 是用来配置 repository.Core 的组件
type CoreConfigurer interface {
	Configure(c *Core) error
}
