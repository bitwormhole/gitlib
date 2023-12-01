package objects

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
)

type sparseObjectImpl struct {
	id   git.ObjectID
	file afs.Path
}

func (inst *sparseObjectImpl) _Impl() store.SparseObject {
	return inst
}

func (inst *sparseObjectImpl) Path() afs.Path {
	return inst.file
}

func (inst *sparseObjectImpl) GetID() git.ObjectID {
	return inst.id
}

func (inst *sparseObjectImpl) Exists() bool {
	return inst.file.Exists()
}
