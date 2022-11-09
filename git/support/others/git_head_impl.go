package others

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitHeadImpl ...
type GitHeadImpl struct {
	Core *store.Core
	path afs.Path
}

func (inst *GitHeadImpl) _Impl() store.HEAD {
	return inst
}

// Path ...
func (inst *GitHeadImpl) Path() afs.Path {
	p := inst.path
	if p == nil {
		layout := inst.Core.Layout
		p = layout.HEAD()
		inst.path = p
	}
	return p
}

// NodeType ...
func (inst *GitHeadImpl) NodeType() store.NodeType {
	return store.NodeHEAD
}

// GetValue ...
func (inst *GitHeadImpl) GetValue(s store.Session) (git.ReferenceName, error) {
	h, err := s.LoadHEAD(inst)
	if err != nil {
		return "", err
	}
	return h.Name, nil
}
