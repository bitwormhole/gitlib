package others

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/store"
)

// RepositoryLoaderImpl ...
type RepositoryLoaderImpl struct {
	Context *store.Context
}

func (inst *RepositoryLoaderImpl) _Impl() store.RepositoryLoader {
	return inst
}

// Load ...
func (inst *RepositoryLoaderImpl) Load(l store.RepositoryLayout) (store.Repository, error) {

	core := &store.Core{}
	core.Context = inst.Context
	core.Layout = l
	core.WD = l.WD()

	err := inst.configCore(core)
	if err != nil {
		return nil, err
	}

	repo := core.Repository
	if repo == nil {
		return nil, errors.New("repository object is nil")
	}

	return repo, nil
}

func (inst *RepositoryLoaderImpl) configCore(core *store.Core) error {
	src := inst.Context.CoreConfigurers
	for _, cc := range src {
		err := cc.Configure(core)
		if err != nil {
			return err
		}
	}
	return nil
}