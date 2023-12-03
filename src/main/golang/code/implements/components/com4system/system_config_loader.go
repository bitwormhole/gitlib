package com4system

import (
	"github.com/bitwormhole/gitlib/git/data/config"
	"github.com/bitwormhole/gitlib/git/repositories"
)

// SystemConfigLoaderRegistry ...
type SystemConfigLoaderRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *SystemConfigLoaderRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *SystemConfigLoaderRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SystemConfigLoaderRegistry) create(ctx *repositories.SystemContext) (any, error) {
	com := &systemConfigLoaderImpl{context: ctx}
	com.load()
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type systemConfigLoaderImpl struct {
	context *repositories.SystemContext
}

func (inst *systemConfigLoaderImpl) _impl() any {
	return inst
}

func (inst *systemConfigLoaderImpl) load() error {

	ctx := inst.context
	ccf := config.NewChainFactory(ctx.FS)
	parent := ccf.Root()
	file := ctx.FS.NewPath("/etc/gitconfig")

	chain := ccf.Create(&repositories.ConfigChainParams{
		File:       file,
		Parent:     parent,
		Scope:      repositories.ConfigScopeSystem,
		IgnoreCase: true,
		Required:   false,
	})

	ctx.Config = chain.Config()
	ctx.ConfigChain = chain
	ctx.ConfigFile = file
	ctx.ConfigChainFactory = ccf

	if file.IsFile() {
		ctx.Config.Load()
	}

	return nil
}
