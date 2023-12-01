package repositories

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/starter-go/afs"
	// "github.com/starter-go/cli"
)

// Context 表示仓库对象的周边环境
type Context struct {
	Lib Lib

	// CLI cli.CLI

	Algorithms []git.AlgorithmRegistry

	Services []ServiceRegistry

	CoreConfigurers []CoreConfigurer

	Connectors []pktline.Connector

	FS afs.FS

	Locator Locator

	Finder Finder

	Loader Loader

	ServiceManager ServiceManager

	ConfigChainFactory ConfigChainFactory

	AlgorithmManager AlgorithmManager

	ConnectorManager pktline.ConnectorManager // the main connector
}

// ContextConfiguration 是用来初始化模块的配置
type ContextConfiguration struct {
	Factory            ContextFactory
	ContextConfigurers []ContextConfigurer
	CoreConfigurers    []CoreConfigurer

	// UseCLI    bool
	// CLIConfig *cli.Configuration
	// CLI       cli.CLI
}

// ContextFactory  是用来创建 repository.Context 的工厂
type ContextFactory interface {
	Create(cfg *ContextConfiguration) (*Context, error)

	// get info about the factory
	String() string
}

// ContextConfigurer 是用来配置 repository.Context 的组件
// [inject:".git-context-configurer"]
type ContextConfigurer interface {
	Configure(c *Context) error
}

////////////////////////////////////////////////////////////////////////////////

// SystemContext ...
type SystemContext struct {
	Config      Config
	ConfigChain ConfigChain
}

// UserContext ...
type UserContext struct {
	Parent      *SystemContext
	Home        afs.Path // 用户的主文件夹
	Config      Config
	ConfigChain ConfigChain
}

// RepositoryContext ...
type RepositoryContext struct {
	Parent      *UserContext
	Config      Config
	ConfigChain ConfigChain
	Layout      Layout
	Facade      Repository
	Submodules  Submodules
	Worktrees   Worktrees
}

// ViewportContext ...
type ViewportContext struct {
	Parent *RepositoryContext
	Layout Layout
}

// SubmoduleContext ...
type SubmoduleContext struct {
	ViewportContext

	Facade Submodule
}

// WorktreeContext ...
type WorktreeContext struct {
	ViewportContext

	Facade Worktree
}
