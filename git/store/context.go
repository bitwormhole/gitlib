package store

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/services"
)

// Context 表示仓库对象的周边环境
type Context struct {
	// Context context.Context

	Services []services.ServiceRegistry

	FS afs.FS

	Lib Lib

	Locator RepositoryLocator

	Finder RepositoryFinder

	ServiceManager services.ServiceManager

	ProfileFactory RepositoryProfileFactory

	ConfigChainFactory ConfigChainFactory
}

// ContextConfiguration 是用来初始化模块的配置
type ContextConfiguration struct {
	Factory            ContextFactory
	ContextConfigurers []ContextConfigurer
	CoreConfigurers    []CoreConfigurer
}

// ContextFactory  是用来创建 repository.Context 的工厂
type ContextFactory interface {
	Create(cfg *ContextConfiguration) *Context

	// get info about the factory
	String() string
}

// ContextConfigurer 是用来配置 repository.Context 的组件
type ContextConfigurer interface {
	Configure(c *Context) error
}

////////////////////////////////////////////////////////////////////////////////
