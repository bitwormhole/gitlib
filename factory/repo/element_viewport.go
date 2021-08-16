package repo

import (
	"errors"

	"github.com/bitwormhole/gitlib/repository"
)

// ViewportElementFactory vpt 工厂
type ViewportElementFactory struct {
}

func (inst *ViewportElementFactory) _Impl() ElementFactory {
	return inst
}

// Make 生成Viewport
func (inst *ViewportElementFactory) Make(ctx *ViewportContext) error {
	vpt := &viewportElement{}
	ctx.viewport = vpt
	vpt.context = ctx
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type viewportElement struct {
	context *ViewportContext
}

func (inst *viewportElement) _Impl() repository.Viewport {
	return inst
}

func (inst *viewportElement) Close() error {
	pool := inst.context.center.pool
	if pool == nil {
		return nil
	}
	return pool.Release()
}

func (inst *viewportElement) GetPWD() repository.GitPWD {
	return inst.context.pwd
}

// func (inst *viewportElement) Workspace() repository.GitWorkspace {
// 	return inst.context.workspace
// }

func (inst *viewportElement) GetCore() repository.Core {
	ctx := inst.context
	if ctx.isCore && ctx.core != nil {
		return ctx.core
	}
	ctx = ctx.center.core
	if ctx.isCore && ctx.core != nil {
		return ctx.core
	}
	panic("no core in current viewport")
}

func (inst *viewportElement) Config() repository.GitConfig {
	return inst.context.config
}

func (inst *viewportElement) Objects() repository.GitObjects {
	return inst.context.objects
}

func (inst *viewportElement) Index() repository.GitIndex {
	return inst.context.index
}

func (inst *viewportElement) Refs() repository.GitRefs {
	return inst.context.refs
}

func (inst *viewportElement) HEAD() repository.GitHEAD {
	return inst.context.HEAD
}

func (inst *viewportElement) GetWorktree() (repository.GitWorktree, error) {
	wktree := inst.context.worktree
	if wktree == nil {
		return nil, errors.New("No worktree in this viewport")
	}
	return wktree, nil
}

func (inst *viewportElement) GetSubmodule() (repository.GitSubmodule, error) {
	mod := inst.context.submodule
	if mod == nil {
		return nil, errors.New("No submodule in this viewport")
	}
	return mod, nil
}

func (inst *viewportElement) GetWorkspace() (repository.GitWorkspace, error) {
	space := inst.context.workspace
	if space == nil {
		return nil, errors.New("No workspace in this viewport")
	}
	return space, nil
}
