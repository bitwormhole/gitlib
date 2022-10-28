package store

import "bitwormhole.com/starter/afs"

// ConfigScope 表示配置属性的作用域
type ConfigScope int

// 定义配置属性的作用域
const (
	ConfigScopeDefault    ConfigScope = 0
	ConfigScopeSystem     ConfigScope = 1
	ConfigScopeUser       ConfigScope = 2
	ConfigScopeRepository ConfigScope = 3
	ConfigScopeMix        ConfigScope = 4
)

// RepositoryConfig  表示基本的配置文件
type RepositoryConfig interface {
	Import(src map[string]string)
	Export() map[string]string
	GetProperty(name string) string
	SetProperty(name, value string)
	Clear()
}

// Config 统一的抽象配置接口
type Config interface {
	NodeLocation
	RepositoryConfig
	Save() error
	Load() error
}

// ConfigChain ... 表示配置对象构成的责任链
type ConfigChain interface {
	Config() Config

	Mix() ConfigChain

	Parent() ConfigChain

	Scope() ConfigScope

	FindByScope(scope ConfigScope) ConfigChain

	Load() error
}

// ConfigChainFactory ...
type ConfigChainFactory interface {
	Create(file afs.Path, parent ConfigChain, scope ConfigScope, required bool) ConfigChain

	Root() ConfigChain
}
