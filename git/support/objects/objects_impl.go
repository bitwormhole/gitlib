package objects

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitObjectsImpl ...
type GitObjectsImpl struct {
	Core *store.Core

	// cache
	path    afs.Path
	pathMap store.PathMapping
}

func (inst *GitObjectsImpl) _Impl() store.Objects {
	return inst
}

func (inst *GitObjectsImpl) getPathMap() store.PathMapping {
	pm := inst.pathMap
	if pm != nil {
		return pm
	}
	pm = inst.Core.PathMapping
	inst.pathMap = pm
	return pm
}

// Path ...
func (inst *GitObjectsImpl) Path() afs.Path {
	p := inst.path
	if p != nil {
		return p
	}
	layout := inst.Core.Layout
	p = layout.Objects()
	inst.path = p
	return p
}

// GetSparseObject ...
func (inst *GitObjectsImpl) GetSparseObject(oid git.ObjectID) store.SparseObject {
	base := inst.Path()
	pm := inst.getPathMap()
	file := pm.Map(base, oid)
	return &sparseObjectImpl{
		file: file,
		id:   oid,
	}
}

func (inst *GitObjectsImpl) GetPack(pid git.PackID) store.Pack {
	panic("no impl")
}

func (inst *GitObjectsImpl) Info() store.InfoFolder {
	panic("no impl")
}
