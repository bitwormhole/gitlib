package loaders

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/base/util"
)

// SystemContextLoaderImpl ...
type SystemContextLoaderImpl struct {

	//starter:component
	_as func(repositories.SystemContextLoader, repositories.ComponentRegistry) //starter:as("#",".")

	FS            afs.FS                           //starter:inject("#")
	AllComponents []repositories.ComponentRegistry //starter:inject(".")
	UseSafeMode   bool                             //starter:inject("${git.threads.use-safe-mode}")

}

func (inst *SystemContextLoaderImpl) _impl() (repositories.SystemContextLoader, repositories.ComponentRegistry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *SystemContextLoaderImpl) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.createComp,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SystemContextLoaderImpl) createComp(ctx *repositories.SystemContext) (any, error) {
	ctx.SystemContextLoader = inst // singleton
	return inst, nil
}

func (inst *SystemContextLoaderImpl) getSafeMode() safe.Mode {
	if inst.UseSafeMode {
		return safe.Safe()
	}
	return safe.Fast()
}

// Load ...
func (inst *SystemContextLoaderImpl) Load(params *repositories.SystemParams) (*repositories.SystemContext, error) {

	err := inst.checkParams(params)
	if err != nil {
		return nil, err
	}

	ctx := &repositories.SystemContext{}
	ctx.SystemContextLoader = inst
	ctx.AllComponents = inst.AllComponents
	ctx.SafeMode = params.Mode
	ctx.FS = inst.FS
	ctx.Path = inst.FS.NewPath("/etc/gitconfig")

	steps := make([]func(ctx *repositories.SystemContext) error, 0)
	steps = append(steps, inst.loadAllComponents)
	steps = append(steps, inst.loadSystemComponents)
	steps = append(steps, inst.openNewLife)

	for _, step := range steps {
		err := step(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

func (inst *SystemContextLoaderImpl) checkParams(params *repositories.SystemParams) error {

	if params == nil {
		return fmt.Errorf("no params")
	}

	if params.Mode == nil {
		params.Mode = inst.getSafeMode()
	}

	return nil
}

func (inst *SystemContextLoaderImpl) loadAllComponents(ctx *repositories.SystemContext) error {

	// list
	src := ctx.AllComponents
	dst := make([]*repositories.ComponentRegistration, 0)
	for _, r1 := range src {
		r2s := r1.ListRegistrations()
		dst = append(dst, r2s...)
	}

	// sort
	sort := &util.Sorter{
		OnLen:  func() int { return len(dst) },
		OnLess: func(i1, i2 int) bool { return dst[i1].Priority < dst[i2].Priority },
		OnSwap: func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] },
	}
	sort.Sort()

	// classify
	for _, item := range dst {
		if item.Enabled {
			inst.tryAddComForSystem(ctx, item)
			inst.tryAddComForUser(ctx, item)
			inst.tryAddComForRepository(ctx, item)
			inst.tryAddComForSession(ctx, item)
			inst.tryAddComForSubmodule(ctx, item)
			inst.tryAddComForWorktree(ctx, item)
		}
	}

	return nil
}

func (inst *SystemContextLoaderImpl) loadSystemComponents(ctx *repositories.SystemContext) error {
	list := ctx.SystemComponents
	comlife := ctx.ComLifeManager.Init(list, ctx.SafeMode)
	comlife.CreateItems(func(h *repositories.ComponentHolder) error {
		maker := h.Registration.OnInitForSystem
		com, err := maker(ctx)
		if err == nil && com != nil {
			h.Component = com
		}
		return err
	})
	return nil
}

func (inst *SystemContextLoaderImpl) openNewLife(ctx *repositories.SystemContext) error {
	lifeManager := ctx.ComLifeManager.GetLifecycles()
	closer, err := lifeManager.Open()
	if err != nil {
		return err
	}
	ctx.Closer = closer
	return nil
}

func (inst *SystemContextLoaderImpl) tryAddComForSystem(ctx *repositories.SystemContext, item *repositories.ComponentRegistration) {
	if item.OnInitForSystem != nil {
		ctx.SystemComponents = append(ctx.SystemComponents, item)
	}
}

func (inst *SystemContextLoaderImpl) tryAddComForUser(ctx *repositories.SystemContext, item *repositories.ComponentRegistration) {
	if item.OnInitForUser != nil {
		ctx.UserComponents = append(ctx.UserComponents, item)
	}
}

func (inst *SystemContextLoaderImpl) tryAddComForRepository(ctx *repositories.SystemContext, item *repositories.ComponentRegistration) {
	if item.OnInitForRepository != nil {
		ctx.RepositoryComponents = append(ctx.RepositoryComponents, item)
	}
}

func (inst *SystemContextLoaderImpl) tryAddComForSession(ctx *repositories.SystemContext, item *repositories.ComponentRegistration) {
	if item.OnInitForSession != nil {
		ctx.SessionComponents = append(ctx.SessionComponents, item)
	}
}

func (inst *SystemContextLoaderImpl) tryAddComForWorktree(ctx *repositories.SystemContext, item *repositories.ComponentRegistration) {
	if item.OnInitForWorktree != nil {
		ctx.WorktreeComponents = append(ctx.WorktreeComponents, item)
	}
}

func (inst *SystemContextLoaderImpl) tryAddComForSubmodule(ctx *repositories.SystemContext, item *repositories.ComponentRegistration) {
	if item.OnInitForSubmodule != nil {
		ctx.SubmoduleComponents = append(ctx.SubmoduleComponents, item)
	}
}
