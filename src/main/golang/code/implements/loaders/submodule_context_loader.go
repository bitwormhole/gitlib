package loaders

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// SubmoduleContextLoaderImpl ...
type SubmoduleContextLoaderImpl struct {

	//starter:component
	_as func(repositories.SubmoduleContextLoader, repositories.ComponentRegistry) //starter:as("#",".")

}

func (inst *SubmoduleContextLoaderImpl) _impl() repositories.SubmoduleContextLoader {
	return inst
}

// ListRegistrations ...
func (inst *SubmoduleContextLoaderImpl) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.createComp,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SubmoduleContextLoaderImpl) createComp(ctx *repositories.SystemContext) (any, error) {
	ctx.SubmoduleContextLoader = inst // singleton
	return inst, nil
}

// Load ...
func (inst *SubmoduleContextLoaderImpl) Load(params *repositories.SubmoduleParams) (*repositories.SubmoduleContext, error) {

	err := inst.checkParams(params)
	if err != nil {
		return nil, err
	}

	ctx := &repositories.SubmoduleContext{}
	ctx.SafeMode = params.Mode
	ctx.Parent = params.Parent
	ctx.FS = params.Parent.FS
	ctx.Layout = params.Layout

	steps := make([]func(ctx *repositories.SubmoduleContext) error, 0)
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

func (inst *SubmoduleContextLoaderImpl) checkParams(params *repositories.SubmoduleParams) error {

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

func (inst *SubmoduleContextLoaderImpl) loadComponents(ctx *repositories.SubmoduleContext) error {
	list := ctx.Parent.Parent.Parent.SubmoduleComponents
	comlife := ctx.ComLifeManager.Init(list, ctx.SafeMode)
	comlife.CreateItems(func(h *repositories.ComponentHolder) error {
		maker := h.Registration.OnInitForSubmodule
		com, err := maker(ctx)
		if err == nil && com != nil {
			h.Component = com
		}
		return err
	})
	return nil
}

func (inst *SubmoduleContextLoaderImpl) openNewLife(ctx *repositories.SubmoduleContext) error {
	lm := ctx.ComLifeManager.GetLifecycles()
	closer, err := lm.Open()
	if err != nil {
		return err
	}
	ctx.Closer = closer
	return nil
}
