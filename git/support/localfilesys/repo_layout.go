package localfilesys

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/starter/markup"
)

type LocalRepoLayout struct {
	markup.Component `id:"git-local-repository-layout"`
}

func (inst *LocalRepoLayout) _Impl() files.Layout {
	return inst
}

func (inst *LocalRepoLayout) MakeView(location *files.RepositoryLocation) (*files.RepositoryView, error) {
	return nil, errors.New("no impl")
}
