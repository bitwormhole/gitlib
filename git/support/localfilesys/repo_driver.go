package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
)

type LocalRepoDriver struct {
	markup.Component `class:"git-repository-driver"`

	Layout  files.Layout            `inject:"#git-local-repository-layout"`
	Locator files.RepositoryLocator `inject:"#git-local-repository-locator"`
	Factory RepoFactory             `inject:"#git-local-repository-factory"`
}

func (inst *LocalRepoDriver) _Impl() repository.Driver {
	return inst
}

func (inst *LocalRepoDriver) Supports(uri lang.URI) bool {
	return uri.Scheme() == "file"
}

func (inst *LocalRepoDriver) Open(uri lang.URI) (repository.View, error) {

	path, err := fs.Default().GetPathByURI(uri)
	if err != nil {
		return nil, err
	}

	location, err := inst.Locator.Locate(path)
	if err != nil {
		return nil, err
	}

	view, err := inst.Layout.MakeView(location)
	if err != nil {
		return nil, err
	}

	repo, err := inst.Factory.Create(view)
	if err != nil {
		return nil, err
	}

	return repo.View, nil
}
