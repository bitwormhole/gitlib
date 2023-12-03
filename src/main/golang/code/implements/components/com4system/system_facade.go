package com4system

import "github.com/bitwormhole/gitlib/git/repositories"

// SystemFacadeRegistry ...
type SystemFacadeRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *SystemFacadeRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *SystemFacadeRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SystemFacadeRegistry) create(ctx *repositories.SystemContext) (any, error) {
	facade := new(systemSpaceImpl)
	facade.context = ctx
	ctx.Facade = facade
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type systemSpaceImpl struct {
	context *repositories.SystemContext
}

func (inst *systemSpaceImpl) _impl() repositories.SystemSpace { return inst }

func (inst *systemSpaceImpl) Lib() repositories.Lib {
	return inst.context.Lib
}
