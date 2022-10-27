package others

import (
	"errors"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

// RepositoryLocatorImpl ...
type RepositoryLocatorImpl struct {
}

func (inst *RepositoryLocatorImpl) _Impl() store.RepositoryLocator {
	return inst
}

// Locate ...
func (inst *RepositoryLocatorImpl) Locate(pwd afs.Path) (store.RepositoryLayout, error) {
	return nil, errors.New("no impl")
}
