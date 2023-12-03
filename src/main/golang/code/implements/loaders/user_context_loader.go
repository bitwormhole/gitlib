package loaders

import (
	"fmt"
	"os"

	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// UserContextLoaderImpl ...
type UserContextLoaderImpl struct {

	//starter:component
	_as func(repositories.UserContextLoader, repositories.ComponentRegistry) //starter:as("#",".")

}

func (inst *UserContextLoaderImpl) _impl() (repositories.UserContextLoader, repositories.ComponentRegistry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *UserContextLoaderImpl) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.createComp,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *UserContextLoaderImpl) createComp(ctx *repositories.SystemContext) (any, error) {
	ctx.UserContextLoader = inst // singleton
	return inst, nil
}

// Load ...
func (inst *UserContextLoaderImpl) Load(params *repositories.UserParams) (*repositories.UserContext, error) {

	err := inst.checkParams(params)
	if err != nil {
		return nil, err
	}

	ctx := &repositories.UserContext{}
	ctx.SafeMode = params.Mode
	ctx.Parent = params.Parent
	ctx.FS = params.Parent.FS
	ctx.Path = params.Home
	ctx.Home = params.Home

	steps := make([]func(ctx *repositories.UserContext) error, 0)
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

func (inst *UserContextLoaderImpl) checkParams(params *repositories.UserParams) error {

	if params == nil {
		return fmt.Errorf("params is nil")
	}

	if params.Parent == nil {
		return fmt.Errorf("param:parent is nil")
	}

	if params.Mode == nil {
		params.Mode = params.Parent.SafeMode
	}

	if params.Home == nil {
		params.Home = inst.getCurrentUserHomeDir(params)
	}

	return nil
}

func (inst *UserContextLoaderImpl) getCurrentUserHomeDir(params *repositories.UserParams) afs.Path {
	path, err := os.UserHomeDir()
	if err != nil {
		path = "/home/current"
	}
	return params.Parent.FS.NewPath(path)
}

func (inst *UserContextLoaderImpl) loadComponents(ctx *repositories.UserContext) error {
	list := ctx.Parent.UserComponents
	comlife := ctx.ComLifeManager.Init(list, ctx.SafeMode)
	comlife.CreateItems(func(h *repositories.ComponentHolder) error {
		maker := h.Registration.OnInitForUser
		com, err := maker(ctx)
		if err == nil && com != nil {
			h.Component = com
		}
		return err
	})
	return nil
}

func (inst *UserContextLoaderImpl) openNewLife(ctx *repositories.UserContext) error {
	lm := ctx.ComLifeManager.GetLifecycles()
	closer, err := lm.Open()
	if err != nil {
		return err
	}
	ctx.Closer = closer
	return nil
}
