package store

import (
	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// Context 表示仓库对象的周边环境
type Context struct {
	Lib Lib

	CLI cli.CLI

	Services []instructions.ServiceRegistry

	CoreConfigurers []CoreConfigurer

	FS afs.FS

	Locator RepositoryLocator

	Finder RepositoryFinder

	ServiceManager instructions.ServiceManager

	RepositoryLoader RepositoryLoader

	ConfigChainFactory ConfigChainFactory
}

// ContextConfiguration 是用来初始化模块的配置
type ContextConfiguration struct {
	Factory            ContextFactory
	ContextConfigurers []ContextConfigurer
	CoreConfigurers    []CoreConfigurer

	UseCLI    bool
	CLIConfig *cli.Configuration
	CLI       cli.CLI
}

// ContextFactory  是用来创建 repository.Context 的工厂
type ContextFactory interface {
	Create(cfg *ContextConfiguration) (*Context, error)

	// get info about the factory
	String() string
}

// ContextConfigurer 是用来配置 repository.Context 的组件
type ContextConfigurer interface {
	Configure(c *Context) error
}

////////////////////////////////////////////////////////////////////////////////
