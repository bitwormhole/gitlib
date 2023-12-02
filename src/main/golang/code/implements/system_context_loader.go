package implements

import (
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/base/util"
)

// SystemContextLoaderImpl ...
type SystemContextLoaderImpl struct {

	//starter:component
	_as func(repositories.SystemContextLoader) //starter:as("#")

	AllComponents []repositories.ComponentRegistry //starter:inject(".")
	UseSafeMode   bool                             //starter:inject("${git.threads.use-safe-mode}")

}

func (inst *SystemContextLoaderImpl) _impl() repositories.SystemContextLoader {
	return inst
}

func (inst *SystemContextLoaderImpl) getSafeMode() safe.Mode {
	if inst.UseSafeMode {
		return safe.Safe()
	}
	return safe.Fast()
}

// Load ...
func (inst *SystemContextLoaderImpl) Load() (*repositories.SystemContext, error) {

	mode := inst.getSafeMode()
	ctx := &repositories.SystemContext{}

	ctx.SystemContextLoader = inst
	ctx.AllComponents = inst.AllComponents
	ctx.SafeMode = mode
	ctx.ComponentLifecycleManager.Init(mode)

	err := inst.loadComponents(ctx)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

func (inst *SystemContextLoaderImpl) loadComponents(ctx *repositories.SystemContext) error {

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

	// load components for system-context
	return inst.loadSystemComponents(ctx)
}

func (inst *SystemContextLoaderImpl) loadSystemComponents(ctx *repositories.SystemContext) error {

	src := ctx.SystemComponents
	dst := ctx.ComponentLifecycleManager.GetComponents()
	holders := make([]*repositories.ComponentHolder, 0)

	for _, item := range src {
		h, err := dst.Register(item)
		if err != nil {
			return err
		}
		holders = append(holders, h)
		onInit := h.Registration.OnInitForSystem
		com, err := onInit(ctx)
		if err != nil {
			return err
		}
		h.Component = com
	}

	// load & open
	ctx.ComponentLifecycleManager.Load()
	lifeManager := ctx.ComponentLifecycleManager.GetLifecycles()
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
