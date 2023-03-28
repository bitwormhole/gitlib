package submodules

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

type workspace struct {
	dir afs.Path
}

func (inst *workspace) _Impl() store.Workspace {
	return inst
}

func (inst *workspace) Path() afs.Path {
	return inst.dir
}
