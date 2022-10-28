package refs

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/data/dxo"
	"github.com/bitwormhole/gitlib/git/store"
)

type GitRefsImpl struct {
	Core *store.Core
}

func (inst *GitRefsImpl) _Impl() store.Refs {
	return inst
}

func (inst *GitRefsImpl) Path() afs.Path {
	return nil
}

func (inst *GitRefsImpl) GetRef(name dxo.ReferenceName) store.Ref {
	return nil
}
