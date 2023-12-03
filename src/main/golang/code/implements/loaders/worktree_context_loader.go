package loaders

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// WorktreeContextLoaderImpl ...
type WorktreeContextLoaderImpl struct {

	//starter:component
	_as func(repositories.WorktreeContextLoader, repositories.ComponentRegistry) //starter:as("#",".")

}

func (inst *WorktreeContextLoaderImpl) _impl() repositories.WorktreeContextLoader {
	return inst
}

// ListRegistrations ...
func (inst *WorktreeContextLoaderImpl) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.createComp,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *WorktreeContextLoaderImpl) createComp(ctx *repositories.SystemContext) (any, error) {
	ctx.WorktreeContextLoader = inst // singleton
	return inst, nil
}

// Load ...
func (inst *WorktreeContextLoaderImpl) Load(params *repositories.WorktreeParams) (*repositories.WorktreeContext, error) {

	err := inst.checkParams(params)
	if err != nil {
		return nil, err
	}

	ctx := &repositories.WorktreeContext{}
	ctx.SafeMode = params.Mode
	ctx.Parent = params.Parent
	ctx.FS = params.Parent.FS
	ctx.Layout = params.Layout

	steps := make([]func(ctx *repositories.WorktreeContext) error, 0)
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

func (inst *WorktreeContextLoaderImpl) checkParams(params *repositories.WorktreeParams) error {

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

func (inst *WorktreeContextLoaderImpl) loadComponents(ctx *repositories.WorktreeContext) error {
	list := ctx.Parent.Parent.Parent.WorktreeComponents
	comlife := ctx.ComLifeManager.Init(list, ctx.SafeMode)
	comlife.CreateItems(func(h *repositories.ComponentHolder) error {
		maker := h.Registration.OnInitForWorktree
		com, err := maker(ctx)
		if err == nil && com != nil {
			h.Component = com
		}
		return err
	})
	return nil
}

func (inst *WorktreeContextLoaderImpl) openNewLife(ctx *repositories.WorktreeContext) error {
	lm := ctx.ComLifeManager.GetLifecycles()
	closer, err := lm.Open()
	if err != nil {
		return err
	}
	ctx.Closer = closer
	return nil
}
