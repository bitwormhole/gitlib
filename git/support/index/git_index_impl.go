package index

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

type GitIndexImpl struct {
	Core *store.Core
}

func (inst *GitIndexImpl) _Impl() store.Index {
	return inst
}

func (inst *GitIndexImpl) Path() afs.Path {
	panic("no impl")
}

func (inst *GitIndexImpl) NodeType() store.NodeType {
	return store.NodeIndex
}
