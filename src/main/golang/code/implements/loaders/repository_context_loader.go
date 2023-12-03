package loaders

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// RepositoryContextLoaderImpl ...
type RepositoryContextLoaderImpl struct {

	//starter:component
	_as func(repositories.RepositoryContextLoader, repositories.ComponentRegistry) //starter:as("#",".")

}

func (inst *RepositoryContextLoaderImpl) _impl() repositories.RepositoryContextLoader {
	return inst
}

// ListRegistrations ...
func (inst *RepositoryContextLoaderImpl) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.createComp,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *RepositoryContextLoaderImpl) createComp(ctx *repositories.SystemContext) (any, error) {
	ctx.RepositoryContextLoader = inst // singleton
	return inst, nil
}

// Load ...
func (inst *RepositoryContextLoaderImpl) Load(params *repositories.RepositoryParams) (*repositories.RepositoryContext, error) {

	err := inst.checkParams(params)
	if err != nil {
		return nil, err
	}

	ctx := &repositories.RepositoryContext{}
	ctx.SafeMode = params.Mode
	ctx.Parent = params.Parent
	ctx.FS = params.Parent.FS
	ctx.Layout = params.Layout

	steps := make([]func(ctx *repositories.RepositoryContext) error, 0)
	steps = append(steps, inst.loadComponents)
	steps = append(steps, inst.openNewLife)

	for _, step := range steps {
		err := step(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil

}

func (inst *RepositoryContextLoaderImpl) checkParams(params *repositories.RepositoryParams) error {

	if params == nil {
		return fmt.Errorf("params is nil")
	}

	if params.Parent == nil {
		return fmt.Errorf("param:parent is nil")
	}

	if params.Layout == nil {
		return fmt.Errorf("param:layout is nil")
	}

	if params.Mode == nil {
		params.Mode = params.Parent.SafeMode
	}

	return nil
}

func (inst *RepositoryContextLoaderImpl) loadComponents(ctx *repositories.RepositoryContext) error {
	list := ctx.Parent.Parent.RepositoryComponents
	comlife := ctx.ComLifeManager.Init(list, ctx.SafeMode)
	comlife.CreateItems(func(h *repositories.ComponentHolder) error {
		maker := h.Registration.OnInitForRepository
		com, err := maker(ctx)
		if err == nil && com != nil {
			h.Component = com
		}
		return err
	})
	return nil
}

func (inst *RepositoryContextLoaderImpl) openNewLife(ctx *repositories.RepositoryContext) error {
	lm := ctx.ComLifeManager.GetLifecycles()
	closer, err := lm.Open()
	if err != nil {
		return err
	}
	ctx.Closer = closer
	return nil
}
