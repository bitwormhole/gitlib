package repo

import (
	"errors"

	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/lang"
)

// CenterContext 是多个 ViewportContext 的共享中心
type CenterContext struct {
	pool lang.ReleasePool
	core *ViewportContext
}

// ViewportContext 仓库的视口上下文
type ViewportContext struct {
	center   *CenterContext
	layout   *Layout
	elements []Element
	isCore   bool

	// elements
	objects   repository.GitObjects
	refs      repository.GitRefs
	config    repository.GitConfig
	index     repository.GitIndex
	HEAD      repository.GitHEAD
	workspace repository.GitWorkspace
	worktree  repository.GitWorktree
	submodule repository.GitSubmodule

	worktrees repository.GitWorktrees
	modules   repository.GitModules

	// facade
	core     repository.Core
	viewport repository.Viewport
	pwd      repository.GitPWD
}

func (inst *ViewportContext) ToViewport() (repository.Viewport, error) {
	vpt := inst.viewport
	if vpt == nil {
		return nil, errors.New("no viewport")
	}
	return vpt, nil
}
