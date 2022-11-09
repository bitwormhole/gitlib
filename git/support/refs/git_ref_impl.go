package refs

import (
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

type referenceImpl struct {
	path afs.Path
	name git.ReferenceName
}

func (inst *referenceImpl) _Impl() store.Ref {
	return inst
}

func (inst *referenceImpl) Path() afs.Path {
	return inst.path
}

func (inst *referenceImpl) Name() git.ReferenceName {
	return inst.name
}

func (inst *referenceImpl) Exists() bool {
	return inst.path.Exists()
}

func (inst *referenceImpl) GetValue(s store.Session) (git.ObjectID, error) {
	r, err := s.LoadRef(inst)
	if err != nil {
		return nil, err
	}
	id := r.ID
	if id == nil {
		return nil, fmt.Errorf("ref.id==nil")
	}
	return id, nil
}
