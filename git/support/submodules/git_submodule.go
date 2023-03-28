package submodules

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

type submodule struct {
	core   *store.Core
	name   string
	dotgit afs.Path // the '.git' file
	gitdir afs.Path
}

func (inst *submodule) _Impl() store.Submodule {
	return inst
}

func (inst *submodule) Name() string {
	return inst.name
}

func (inst *submodule) Workspace() store.Workspace {
	dir := inst.dotgit.GetParent()
	return &workspace{dir}
}

func (inst *submodule) Exists() bool {
	p1 := inst.dotgit
	p2 := inst.gitdir
	if p1 == nil || p2 == nil {
		return false
	}
	return p1.Exists() && p2.Exists()
}
