package com4user

import (
	"github.com/bitwormhole/gitlib/git/repositories"
)

// UserConfigLoaderRegistry ...
type UserConfigLoaderRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *UserConfigLoaderRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *UserConfigLoaderRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:       true,
		OnInitForUser: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *UserConfigLoaderRegistry) create(ctx *repositories.UserContext) (any, error) {
	com := &userConfigLoader{context: ctx}
	com.load()
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type userConfigLoader struct {
	context *repositories.UserContext
}

func (inst *userConfigLoader) _impl() any {
	return inst
}

func (inst *userConfigLoader) load() error {

	ctx := inst.context
	ccFactory := ctx.Parent.ConfigChainFactory
	parent := ctx.Parent.ConfigChain
	home := ctx.Home
	file := home.GetChild(".gitconfig")

	chain := ccFactory.Create(&repositories.ConfigChainParams{
		File:       file,
		Parent:     parent,
		Scope:      repositories.ConfigScopeUser,
		Required:   false,
		IgnoreCase: true,
	})

	ctx.Config = chain.Config()
	ctx.ConfigChain = chain
	ctx.ConfigFile = file

	if file.IsFile() {
		ctx.Config.Load()
	}

	return nil
}
