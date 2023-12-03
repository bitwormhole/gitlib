package com4worktree

import "github.com/bitwormhole/gitlib/git/repositories"

// WorktreeFacadeRegistry ...
type WorktreeFacadeRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *WorktreeFacadeRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *WorktreeFacadeRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:           true,
		OnInitForWorktree: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *WorktreeFacadeRegistry) create(ctx *repositories.WorktreeContext) (any, error) {
	facade := new(worktreeFacade)
	facade.context = ctx
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type worktreeFacade struct {
	context *repositories.WorktreeContext
}
