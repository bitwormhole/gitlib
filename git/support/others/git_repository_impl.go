package others

import "github.com/bitwormhole/gitlib/git/store"

// GitRepositoryImpl ...
type GitRepositoryImpl struct {
	Core *store.Core
}

func (inst *GitRepositoryImpl) _Impl() store.Repository {
	return inst
}

// Layout ...
func (inst *GitRepositoryImpl) Layout() store.RepositoryLayout {
	return inst.Core.Layout
}

// Config ...
func (inst *GitRepositoryImpl) Config() store.ConfigChain {
	return inst.Core.Config
}

// HEAD ...
func (inst *GitRepositoryImpl) HEAD() store.HEAD {
	return inst.Core.Head
}

// Index ...
func (inst *GitRepositoryImpl) Index() store.Index {
	return inst.Core.Index
}

// Refs ...
func (inst *GitRepositoryImpl) Refs() store.Refs {
	return inst.Core.Refs
}

// Objects ...
func (inst *GitRepositoryImpl) Objects() store.Objects {
	return inst.Core.Objects
}

// OpenSession ...
func (inst *GitRepositoryImpl) OpenSession() (store.Session, error) {
	factory := inst.Core.SessionFactory
	return factory.OpenSession(inst)
}
