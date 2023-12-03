package com4user

import (
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// UserFacadeRegistry ...
type UserFacadeRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *UserFacadeRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *UserFacadeRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:       true,
		OnInitForUser: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *UserFacadeRegistry) create(ctx *repositories.UserContext) (any, error) {
	facade := new(userSpaceImpl)
	facade.context = ctx
	ctx.Facade = facade
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type userSpaceImpl struct {
	context *repositories.UserContext
}

func (inst *userSpaceImpl) _impl() repositories.UserSpace { return inst }

func (inst *userSpaceImpl) Home() afs.Path {
	return inst.context.Home
}

func (inst *userSpaceImpl) SystemSpace() repositories.SystemSpace {
	return inst.context.Parent.Facade
}
