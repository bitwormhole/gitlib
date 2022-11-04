package config

import (
	"os"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

// ChainFactory ... 实现配置链工厂
type ChainFactory struct {
	fs afs.FS // cached
}

func (inst *ChainFactory) _Impl() store.ConfigChainFactory {
	return inst
}

func (inst *ChainFactory) getFS() afs.FS {
	f := inst.fs
	if f == nil {
		f = files.FS()
		inst.fs = f
	}
	return f
}

// Create ...
func (inst *ChainFactory) Create(p *store.ConfigChainParams) store.ConfigChain {
	config := &simpleConfig{}
	config.init(p)
	return &configChainNode{
		scope:    p.Scope,
		parent:   p.Parent,
		required: p.Required,

		config: config,
	}
}

// Root ... 取根链
func (inst *ChainFactory) Root() store.ConfigChain {
	file1 := inst.getSystemConfigFile()
	file2 := inst.getUserConfigFile()
	builder := configChainBuilder{factory: inst}
	builder.add(nil, store.ConfigScopeDefault, false)
	builder.add(file1, store.ConfigScopeSystem, false)
	builder.add(file2, store.ConfigScopeUser, false)
	chain := builder.create()
	inst.initDefaultConfig(chain)
	return chain
}

func (inst *ChainFactory) initDefaultConfig(chain store.ConfigChain) {

	src := make(map[gitconfig.KeyTemplate]string)

	src[gitconfig.CoreCompressionAlgorithm] = "deflate"
	src[gitconfig.CoreDigestAlgorithm] = "sha1"
	src[gitconfig.CoreObjectsPathPattern] = "xx/xxxx"

	dst := chain.FindByScope(store.ConfigScopeDefault).Config()
	for k, v := range src {
		dst.SetProperty(k.String(), v)
	}
}

func (inst *ChainFactory) getUserConfigFile() afs.Path {
	uhdir, err := os.UserHomeDir()
	if err != nil {
		vlog.Warn(err)
		return nil
	}
	base := inst.getFS().NewPath(uhdir)
	return base.GetChild(".gitconfig")
}

func (inst *ChainFactory) getSystemConfigFile() afs.Path {
	const p = "/etc/gitconfig"
	return inst.getFS().NewPath(p)
}

////////////////////////////////////////////////////////////////////////////////

type configChainBuilder struct {
	chain   store.ConfigChain
	factory store.ConfigChainFactory
}

func (inst *configChainBuilder) add(file afs.Path, scope store.ConfigScope, required bool) {
	parent := inst.chain
	params := &store.ConfigChainParams{
		File:       file,
		Parent:     parent,
		Scope:      scope,
		Required:   required,
		IgnoreCase: true,
	}
	child := inst.factory.Create(params)
	inst.chain = child
}

func (inst *configChainBuilder) create() store.ConfigChain {
	return inst.chain
}

////////////////////////////////////////////////////////////////////////////////

type configChainNode struct {
	config   store.Config
	scope    store.ConfigScope
	parent   store.ConfigChain
	required bool
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

func (inst *configChainNode) Load() error {

	// parent
	parent := inst.parent
	if parent != nil {
		err := inst.parent.Load()
		if err != nil {
			return err
		}
	}

	// this
	cfg := inst.Config()
	err := cfg.Load()
	if err != nil {
		if inst.required {
			return err
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
