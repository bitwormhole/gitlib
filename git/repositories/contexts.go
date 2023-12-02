package repositories

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/safe"
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
	Facade      Lib
	FS          afs.FS
	SafeMode    safe.Mode

	Closer                    io.Closer
	ComponentLifecycleManager ComponentLifecycleManager

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
	Parent                    *SystemContext
	Home                      afs.Path // 用户的主文件夹
	Config                    Config
	ConfigChain               ConfigChain
	ComponentLifecycleManager ComponentLifecycleManager
}

// RepositoryContext ...
type RepositoryContext struct {
	Parent                    *UserContext
	Config                    Config
	ConfigChain               ConfigChain
	Layout                    Layout
	Facade                    Repository
	Submodules                Submodules
	Worktrees                 Worktrees
	ComponentLifecycleManager ComponentLifecycleManager
}

// ViewportContext 是一个抽象的上下文，具体参考 SubmoduleContext & WorktreeContext
type ViewportContext struct {
	Parent                    *RepositoryContext
	Layout                    Layout
	ComponentLifecycleManager ComponentLifecycleManager
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

// SessionContext ...
type SessionContext struct {
	Parent                    *RepositoryContext
	ComponentLifecycleManager ComponentLifecycleManager
	Session                   Session
}

////////////////////////////////////////////////////////////////////////////////
