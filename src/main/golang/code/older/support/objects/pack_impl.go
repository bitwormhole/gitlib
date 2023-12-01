package objects

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
)

////////////////////////////////////////////////////////////////////////////////

type packImpl struct {
	id             git.PackID
	packEntityFile afs.Path
	packIndexFile  afs.Path
}

func (inst *packImpl) _Impl() store.Pack {
	return inst
}

func (inst *packImpl) GetID() git.PackID {
	return inst.id
}

func (inst *packImpl) GetDotIdx() afs.Path {
	return inst.packIndexFile
}

func (inst *packImpl) GetDotPack() afs.Path {
	return inst.packEntityFile
}

func (inst *packImpl) Exists() bool {
	file := inst.packIndexFile
	return file.Exists()
}

func (inst *packImpl) GetObject(oid git.ObjectID) store.PackObject {
	return &packObjectImpl{
		id:        oid,
		container: inst,
	}
}

////////////////////////////////////////////////////////////////////////////////

type packObjectImpl struct {
	container store.Pack
	id        git.ObjectID
}

func (inst *packObjectImpl) _Impl() store.PackObject {
	return inst
}

func (inst *packObjectImpl) GetID() git.ObjectID {
	return inst.id
}

func (inst *packObjectImpl) Container() store.Pack {
	return inst.container
}

////////////////////////////////////////////////////////////////////////////////
