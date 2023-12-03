package cases

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/application"
)

// RepositoryPathProvider ...
type RepositoryPathProvider interface {
	GetRepoPath(name string) (afs.Path, error)
}

// RepositoryPathProviderImpl ...
type RepositoryPathProviderImpl struct {

	//starter:component
	_as func(RepositoryPathProvider) //starter:as("#")

	AC application.Context //starter:inject("context")
	FS afs.FS              //starter:inject("#")

}

func (inst *RepositoryPathProviderImpl) _impl() RepositoryPathProvider { return inst }

// GetRepoPath ...
func (inst *RepositoryPathProviderImpl) GetRepoPath(name1 string) (afs.Path, error) {

	name2 := "test." + name1 + ".repo"
	path, err := inst.AC.GetEnvironment().GetEnvRequired(name2)
	if err != nil {
		return nil, err
	}

	p2 := inst.FS.NewPath(path)
	return p2, nil
}
