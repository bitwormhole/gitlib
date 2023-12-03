package com4repository

import "github.com/bitwormhole/gitlib/git/repositories"

// RepoFacadeRegistry ...
type RepoFacadeRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *RepoFacadeRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *RepoFacadeRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:             true,
		OnInitForRepository: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *RepoFacadeRegistry) create(ctx *repositories.RepositoryContext) (any, error) {
	facade := new(repositoryFacade)
	facade.context = ctx
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type repositoryFacade struct {
	context *repositories.RepositoryContext
}
