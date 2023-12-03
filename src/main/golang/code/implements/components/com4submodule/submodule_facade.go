package com4submodule

import "github.com/bitwormhole/gitlib/git/repositories"

// SubmoduleFacadeRegistry ...
type SubmoduleFacadeRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *SubmoduleFacadeRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *SubmoduleFacadeRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:            true,
		OnInitForSubmodule: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SubmoduleFacadeRegistry) create(ctx *repositories.SubmoduleContext) (any, error) {
	facade := new(submoduleFacade)
	facade.context = ctx
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type submoduleFacade struct {
	context *repositories.SubmoduleContext
}
