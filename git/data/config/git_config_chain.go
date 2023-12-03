package config

import (
	"os"

	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
	"github.com/starter-go/vlog"
)

// NewChainFactory ...
func NewChainFactory(fs afs.FS) repositories.ConfigChainFactory {
	return &ChainFactory{fs: fs}
}

////////////////////////////////////////////////////////////////////////////////

// ChainFactory ... 实现配置链工厂
type ChainFactory struct {
	fs afs.FS // cached
}

func (inst *ChainFactory) _Impl() repositories.ConfigChainFactory {
	return inst
}

func (inst *ChainFactory) getFS() afs.FS {
	f := inst.fs
	// if f == nil {
	// 	f = files.FS()
	// 	inst.fs = f
	// }
	return f
}

// Create ...
func (inst *ChainFactory) Create(p *repositories.ConfigChainParams) repositories.ConfigChain {
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
func (inst *ChainFactory) Root() repositories.ConfigChain {
	file1 := inst.getSystemConfigFile()
	file2 := inst.getUserConfigFile()
	builder := configChainBuilder{factory: inst}
	builder.add(nil, repositories.ConfigScopeDefault, false)
	builder.add(file1, repositories.ConfigScopeSystem, false)
	builder.add(file2, repositories.ConfigScopeUser, false)
	chain := builder.create()
	inst.initDefaultConfig(chain)
	return chain
}

func (inst *ChainFactory) initDefaultConfig(chain repositories.ConfigChain) {

	src := make(map[gitconfig.KeyTemplate]string)

	src[gitconfig.CoreCompressionAlgorithm] = "deflate"
	src[gitconfig.CoreDigestAlgorithm] = "sha1"
	src[gitconfig.CoreObjectsPathPattern] = "xx/xxxx"

	dst := chain.FindByScope(repositories.ConfigScopeDefault).Config()
	for k, v := range src {
		dst.SetProperty(k.String(), v)
	}
}

func (inst *ChainFactory) getUserConfigFile() afs.Path {
	uhdir, err := os.UserHomeDir()
	if err != nil {
		vlog.Warn(err.Error())
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
	chain   repositories.ConfigChain
	factory repositories.ConfigChainFactory
}

func (inst *configChainBuilder) add(file afs.Path, scope repositories.ConfigScope, required bool) {
	parent := inst.chain
	params := &repositories.ConfigChainParams{
		File:       file,
		Parent:     parent,
		Scope:      scope,
		Required:   required,
		IgnoreCase: true,
	}
	child := inst.factory.Create(params)
	inst.chain = child
}

func (inst *configChainBuilder) create() repositories.ConfigChain {
	return inst.chain
}

////////////////////////////////////////////////////////////////////////////////

type configChainNode struct {
	config   repositories.Config
	scope    repositories.ConfigScope
	parent   repositories.ConfigChain
	required bool
}

func (inst *configChainNode) _Impl() repositories.ConfigChain {
	return inst
}

func (inst *configChainNode) Config() repositories.Config {
	return inst.config
}

func (inst *configChainNode) Parent() repositories.ConfigChain {
	return inst.parent
}

func (inst *configChainNode) Scope() repositories.ConfigScope {
	return inst.scope
}

func (inst *configChainNode) FindByScope(scope repositories.ConfigScope) repositories.ConfigChain {
	p := inst._Impl()
	for ; p != nil; p = p.Parent() {
		if p.Scope() == scope {
			return p
		}
	}
	return nil
}

func (inst *configChainNode) Mix() repositories.ConfigChain {
	if inst.scope == repositories.ConfigScopeMix {
		return inst
	}
	cfg, _ := mix(inst)
	if cfg == nil {
		return inst
	}
	return &configChainNode{
		parent: inst,
		scope:  repositories.ConfigScopeMix,
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
