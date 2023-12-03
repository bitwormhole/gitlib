package com4repository

import (
	"github.com/bitwormhole/gitlib/git/repositories"
)

// RepoConfigLoaderRegistry ...
type RepoConfigLoaderRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *RepoConfigLoaderRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *RepoConfigLoaderRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:             true,
		OnInitForRepository: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *RepoConfigLoaderRegistry) create(ctx *repositories.RepositoryContext) (any, error) {
	com := &repositoryConfigLoader{context: ctx}
	err := com.load()
	if err != nil {
		return nil, err
	}
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type repositoryConfigLoader struct {
	context *repositories.RepositoryContext
}

func (inst *repositoryConfigLoader) _impl() any {
	return inst
}

func (inst *repositoryConfigLoader) load() error {

	ctx := inst.context
	ccf := ctx.Parent.Parent.ConfigChainFactory
	parent := ctx.Parent.ConfigChain
	layout := ctx.Layout
	file := layout.Config()

	chain := ccf.Create(&repositories.ConfigChainParams{
		File:       file,
		Parent:     parent,
		Scope:      repositories.ConfigScopeRepository,
		Required:   true,
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
