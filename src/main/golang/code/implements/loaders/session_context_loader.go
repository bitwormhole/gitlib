package loaders

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// SessionContextLoaderImpl ...
type SessionContextLoaderImpl struct {

	//starter:component
	_as func(repositories.SessionContextLoader, repositories.ComponentRegistry) //starter:as("#",".")

}

func (inst *SessionContextLoaderImpl) _impl() repositories.SessionContextLoader {
	return inst
}

// ListRegistrations ...
func (inst *SessionContextLoaderImpl) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.createComp,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SessionContextLoaderImpl) createComp(ctx *repositories.SystemContext) (any, error) {
	ctx.SessionContextLoader = inst // singleton
	return inst, nil
}

// Load ...
func (inst *SessionContextLoaderImpl) Load(params *repositories.SessionParams) (*repositories.SessionContext, error) {

	err := inst.checkParams(params)
	if err != nil {
		return nil, err
	}

	ctx := &repositories.SessionContext{}
	ctx.SafeMode = params.Mode
	ctx.Parent = params.Parent
	ctx.FS = params.Parent.FS

	steps := make([]func(ctx *repositories.SessionContext) error, 0)
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

func (inst *SessionContextLoaderImpl) checkParams(params *repositories.SessionParams) error {

	if params == nil {
		return fmt.Errorf("params is nil")
	}

	if params.Parent == nil {
		return fmt.Errorf("param:parent is nil")
	}

	if params.Mode == nil {
		params.Mode = params.Parent.SafeMode
	}

	return nil
}

func (inst *SessionContextLoaderImpl) loadComponents(ctx *repositories.SessionContext) error {
	list := ctx.Parent.Parent.Parent.SessionComponents
	comlife := ctx.ComLifeManager.Init(list, ctx.SafeMode)
	comlife.CreateItems(func(h *repositories.ComponentHolder) error {
		maker := h.Registration.OnInitForSession
		com, err := maker(ctx)
		if err == nil && com != nil {
			h.Component = com
		}
		return err
	})
	return nil
}

func (inst *SessionContextLoaderImpl) openNewLife(ctx *repositories.SessionContext) error {
	lm := ctx.ComLifeManager.GetLifecycles()
	closer, err := lm.Open()
	if err != nil {
		return err
	}
	ctx.Closer = closer
	return nil
}
