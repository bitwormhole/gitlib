package others

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/services"
	"github.com/bitwormhole/gitlib/git/store"
)

// LibImpl ...
type LibImpl struct {
	Context *store.Context
}

func (inst *LibImpl) _Impl() store.Lib {
	return inst
}

// FS ...
func (inst *LibImpl) FS() afs.FS {
	return inst.Context.FS
}

// RepositoryFinder ...
func (inst *LibImpl) RepositoryFinder() store.RepositoryFinder {
	return inst.Context.Finder
}

// RepositoryLocator ...
func (inst *LibImpl) RepositoryLocator() store.RepositoryLocator {
	return inst.Context.Locator
}

// ServiceManager ...
func (inst *LibImpl) ServiceManager() services.ServiceManager {
	return inst.Context.ServiceManager
}
