package objects

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/data/dxo"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitObjectsImpl ...
type GitObjectsImpl struct {
	Core *store.Core
}

func (inst *GitObjectsImpl) _Impl() store.Objects {
	return inst
}

func (inst *GitObjectsImpl) Path() afs.Path {
	return nil
}

func (inst *GitObjectsImpl) GetObject(oid dxo.ObjectID) store.SparseObject {
	return nil
}

func (inst *GitObjectsImpl) GetPack(pid dxo.PackID) store.Pack {
	return nil
}

func (inst *GitObjectsImpl) Info() store.InfoFolder {
	return nil
}
