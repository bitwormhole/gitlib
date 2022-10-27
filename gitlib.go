package gitlib

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support"
	"github.com/bitwormhole/gitlib/git/support/services"
)

// Init 初始化git模块
func Init(ctx context.Context, cfg *store.ContextConfiguration) context.Context {

	ctx = store.SetupBinding(ctx)
	binding := store.GetBinding(ctx)

	if cfg == nil {
		cfg = makeDefaultConfig()
	}

	factory := cfg.Factory
	if factory == nil {
		factory = &support.DefaultContextFactory{}
		cfg.Factory = factory
	}

	binding.Config(cfg)
	lib := binding.GetLib()
	lib.RepositoryFinder() // panic if nil

	return ctx
}

func makeDefaultConfig() *store.ContextConfiguration {

	cfg := &store.ContextConfiguration{}
	c4ctx := cfg.ContextConfigurers
	c4core := cfg.CoreConfigurers

	c4ctx = append(c4ctx, &support.BaseContextConfigurer{})
	c4ctx = append(c4ctx, &services.Configurer{})

	cfg.ContextConfigurers = c4ctx
	cfg.CoreConfigurers = c4core
	cfg.Factory = &support.DefaultContextFactory{}

	return cfg
}
