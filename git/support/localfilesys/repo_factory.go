package localfilesys

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/starter/markup"
)

type RepoFactory interface {
	Create(view *files.RepositoryView) (*Repo, error)
}

////////////////////////////////////////////////////////////////////////////////

// LocalRepoFactory 用来创建本地仓库
type LocalRepoFactory struct {
	markup.Component `id:"git-local-repository-factory"`

	Elements []ElementFactory `inject:".git-local-element-factory"`
}

func (inst *LocalRepoFactory) _Impl() RepoFactory {
	return inst
}

func (inst *LocalRepoFactory) Create(view *files.RepositoryView) (*Repo, error) {

	return nil, errors.New("no impl")
}
