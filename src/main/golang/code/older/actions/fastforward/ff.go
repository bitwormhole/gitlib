package fastforward

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

// FastForward ...
type FastForward struct {
	theRepo store.Repository
	theOld  git.ObjectID
	theNew  git.ObjectID
	theRef  git.ReferenceName
}

// SetRepository ...
func (inst *FastForward) SetRepository(repo store.Repository) *FastForward {
	inst.theRepo = repo
	return inst
}

// SetOld ...
func (inst *FastForward) SetOld(commit git.ObjectID) *FastForward {
	inst.theOld = commit
	return inst
}

// SetNew ...
func (inst *FastForward) SetNew(commit git.ObjectID) *FastForward {
	inst.theNew = commit
	return inst
}

// SetRef ...
func (inst *FastForward) SetRef(name git.ReferenceName) *FastForward {
	inst.theRef = name
	return inst
}

// Run ...
func (inst *FastForward) Run() error {
	r := &fastForwardRunner{parent: inst}
	return r.run()
}
