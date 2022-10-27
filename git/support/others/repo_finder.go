package others

import (
	"errors"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

// RepositoryFinderImpl ...
type RepositoryFinderImpl struct {
}

func (inst *RepositoryFinderImpl) _Impl() store.RepositoryFinder {
	return inst
}

// Find ...
func (inst *RepositoryFinderImpl) Find(pwd afs.Path) ([]store.RepositoryLayout, error) {
	return nil, errors.New("no impl")
}
