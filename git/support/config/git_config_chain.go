package config

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

// ChainFactory ... 实现配置链工厂
type ChainFactory struct {
}

func (inst *ChainFactory) _Impl() store.ConfigChainFactory {
	return inst
}

// Create ...
func (inst *ChainFactory) Create(file afs.Path, parent store.ConfigChain, scope store.ConfigScope) store.ConfigChain {
	config := &simpleConfig{}
	config.init(file)
	return &configChainNode{
		scope:  scope,
		parent: parent,
		config: config,
	}
}

// Root ... 取根链
func (inst *ChainFactory) Root() store.ConfigChain {
	file1 := inst.getSystemConfigFile()
	file2 := inst.getUserConfigFile()
	builder := configChainBuilder{factory: inst}
	builder.add(file1, store.ConfigScopeSystem)
	builder.add(file2, store.ConfigScopeUser)
	return builder.create()
}

func (inst *ChainFactory) getUserConfigFile() afs.Path {
	return nil
}

func (inst *ChainFactory) getSystemConfigFile() afs.Path {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type configChainBuilder struct {
	chain   store.ConfigChain
	factory store.ConfigChainFactory
}

func (inst *configChainBuilder) add(file afs.Path, scope store.ConfigScope) {
	parent := inst.chain
	child := inst.factory.Create(file, parent, scope)
	inst.chain = child
}

func (inst *configChainBuilder) create() store.ConfigChain {
	return inst.chain
}

////////////////////////////////////////////////////////////////////////////////

type configChainNode struct {
	config store.Config
	scope  store.ConfigScope
	parent store.ConfigChain
}

func (inst *configChainNode) _Impl() store.ConfigChain {
	return inst
}

func (inst *configChainNode) Config() store.Config {
	return inst.config
}

func (inst *configChainNode) Parent() store.ConfigChain {
	return inst.parent
}

func (inst *configChainNode) Scope() store.ConfigScope {
	return inst.scope
}

func (inst *configChainNode) FindByScope(scope store.ConfigScope) store.ConfigChain {
	p := inst._Impl()
	for ; p != nil; p = p.Parent() {
		if p.Scope() == scope {
			return p
		}
	}
	return nil
}

func (inst *configChainNode) Mix() store.ConfigChain {
	if inst.scope == store.ConfigScopeMix {
		return inst
	}
	cfg, _ := mix(inst)
	if cfg == nil {
		return inst
	}
	return &configChainNode{
		parent: inst,
		scope:  store.ConfigScopeMix,
		config: cfg,
	}
}

func (inst *configChainNode) Load(se store.Session) error {
	p := inst._Impl()
	for ; p != nil; p = p.Parent() {
		cfg := p.Config()
		err := cfg.Load(se)
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////