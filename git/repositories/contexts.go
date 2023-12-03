package repositories

import (
	"io"

	"github.com/starter-go/afs"
	"github.com/starter-go/base/safe"
)

// Context 表示仓库对象的周边环境
type Context struct {
	Closer         io.Closer
	ComLifeManager ComponentLifecycleManager
	FS             afs.FS
	SafeMode       safe.Mode
	Path           afs.Path
}

// type Context struct {
// 	Lib Lib

// 	// CLI cli.CLI

// 	Algorithms []git.AlgorithmRegistry

// 	Services []ServiceRegistry

// 	CoreConfigurers []CoreConfigurer

// 	Connectors []pktline.Connector

// 	FS afs.FS

// 	Locator Locator

// 	Finder Finder

// 	Loader Loader

// 	ServiceManager ServiceManager

// 	ConfigChainFactory ConfigChainFactory

// 	AlgorithmManager AlgorithmManager

// 	ConnectorManager pktline.ConnectorManager // the main connector
// }

// // ContextConfiguration 是用来初始化模块的配置
// type ContextConfiguration struct {
// 	Factory            ContextFactory
// 	ContextConfigurers []ContextConfigurer
// 	CoreConfigurers    []CoreConfigurer

// 	// UseCLI    bool
// 	// CLIConfig *cli.Configuration
// 	// CLI       cli.CLI
// }

// // ContextFactory  是用来创建 repository.Context 的工厂
// type ContextFactory interface {
// 	Create(cfg *ContextConfiguration) (*Context, error)

// 	// get info about the factory
// 	String() string
// }

// // ContextConfigurer 是用来配置 repository.Context 的组件
// // [inject:".git-context-configurer"]
// type ContextConfigurer interface {
// 	Configure(c *Context) error
// }

////////////////////////////////////////////////////////////////////////////////

// SystemContext ...
type SystemContext struct {
	Context

	Config             Config
	ConfigChain        ConfigChain
	ConfigChainFactory ConfigChainFactory
	ConfigFile         afs.Path // like '/etc/gitconfig'

	Facade SystemSpace
	Lib    Lib

	// components
	AllComponents        []ComponentRegistry
	SystemComponents     []*ComponentRegistration
	UserComponents       []*ComponentRegistration
	RepositoryComponents []*ComponentRegistration
	SessionComponents    []*ComponentRegistration
	WorktreeComponents   []*ComponentRegistration
	SubmoduleComponents  []*ComponentRegistration

	// loaders
	SystemContextLoader     SystemContextLoader
	UserContextLoader       UserContextLoader
	RepositoryContextLoader RepositoryContextLoader
	SessionContextLoader    SessionContextLoader
	WorktreeContextLoader   WorktreeContextLoader
	SubmoduleContextLoader  SubmoduleContextLoader

	// for repo
	RepositoryFinder  Finder
	RepositoryLocator Locator
	RepositoryLoader  Loader
}

// UserContext ...
type UserContext struct {
	Context

	Parent      *SystemContext
	Home        afs.Path // 用户的主文件夹
	Config      Config
	ConfigChain ConfigChain
	ConfigFile  afs.Path // like '~/.gitconfig'
	Facade      UserSpace
}

// RepositoryContext ...
type RepositoryContext struct {
	Context

	Parent      *UserContext
	Config      Config
	ConfigChain ConfigChain
	ConfigFile  afs.Path // like '.git/config'
	Layout      Layout
	Facade      Repository
	Submodules  Submodules
	Worktrees   Worktrees
}

// SubmoduleContext ...
type SubmoduleContext struct {
	Context

	Parent *RepositoryContext
	Layout Layout
	Facade Submodule
}

// WorktreeContext ...
type WorktreeContext struct {
	Context

	Parent *RepositoryContext
	Layout Layout
	Facade Worktree
}

// SessionContext ...
type SessionContext struct {
	Context

	Parent *RepositoryContext
	Facade Session
}

////////////////////////////////////////////////////////////////////////////////
