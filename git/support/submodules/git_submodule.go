package submodules

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

type submodule struct {
	core *store.Core

	// properties:
	name   string // as property name
	path   string
	url    string
	active bool

	// lazyload:
	dotgit    afs.Path // the '.git' file
	subconfig afs.Path
}

func (inst *submodule) _Impl() store.Submodule {
	return inst
}

func (inst *submodule) Name() string {
	return inst.name
}

func (inst *submodule) Path() string {
	return inst.path
}

func (inst *submodule) URL() string {
	return inst.url
}

func (inst *submodule) IsActive() bool {
	return inst.active
}

func (inst *submodule) Workspace() store.Workspace {
	dir := inst.dotgit.GetParent()
	return &workspace{dir}
}

func (inst *submodule) Exists() bool {
	p1 := inst.dotgit
	p2 := inst.subconfig
	if p1 == nil || p2 == nil {
		return false
	}
	return p1.Exists() && p2.Exists()
}