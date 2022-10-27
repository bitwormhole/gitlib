package others

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/store"
)

// RepositoryProfileFactoryImpl ...
type RepositoryProfileFactoryImpl struct {
}

func (inst *RepositoryProfileFactoryImpl) _Impl() store.RepositoryProfileFactory {
	return inst
}

// Create ...
func (inst *RepositoryProfileFactoryImpl) Create(l store.RepositoryLayout) (store.RepositoryProfile, error) {

	return nil, errors.New("no impl")
}
