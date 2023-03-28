package others

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitWorkspaceFacade ...
type GitWorkspaceFacade struct {
	Core *store.Core

	inner store.Workspace
}

func (inst *GitWorkspaceFacade) _Impl() store.Workspace {
	return inst
}

func (inst *GitWorkspaceFacade) getInner() store.Workspace {

	i := inst.inner
	if i != nil {
		return i
	}

	// load
	c := inst.Core
	builder := GitWorkspaceBuilder{}
	builder.Core = c
	builder.DotGit = c.Layout.DotGit()

	i = builder.Create()
	inst.inner = i
	return i
}

// Path ...
func (inst *GitWorkspaceFacade) Path() afs.Path {
	return inst.getInner().Path()
}

// WorkingDirectory ...
func (inst *GitWorkspaceFacade) WorkingDirectory() afs.Path {
	return inst.getInner().WorkingDirectory()
}

// DotGit ...
func (inst *GitWorkspaceFacade) DotGit() afs.Path {
	return inst.getInner().DotGit()
}

////////////////////////////////////////////////////////////////////////////////

// GitWorkspaceImpl ...
type GitWorkspaceImpl struct {
	core   *store.Core
	dotgit afs.Path
	wkdir  afs.Path
}

func (inst *GitWorkspaceImpl) _Impl() store.Workspace {
	return inst
}

// Path ...
func (inst *GitWorkspaceImpl) Path() afs.Path {
	return inst.wkdir
}

// WorkingDirectory ...
func (inst *GitWorkspaceImpl) WorkingDirectory() afs.Path {
	return inst.wkdir
}

// DotGit ...
func (inst *GitWorkspaceImpl) DotGit() afs.Path {
	return inst.dotgit
}

////////////////////////////////////////////////////////////////////////////////

// GitWorkspaceBuilder ...
type GitWorkspaceBuilder struct {
	Core   *store.Core
	DotGit afs.Path
}

// Create ...
func (inst *GitWorkspaceBuilder) Create() store.Workspace {
	dg := inst.DotGit
	o := &GitWorkspaceImpl{}
	o.core = inst.Core
	o.dotgit = dg
	o.wkdir = dg.GetParent()
	return o
}
