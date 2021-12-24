package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/starter/markup"
)

// RepoFactory 生产仓库的工厂对象
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

// Create 创建仓库对象
func (inst *LocalRepoFactory) Create(view *files.RepositoryView) (*Repo, error) {

	repo := &Repo{}
	efs := inst.Elements

	for _, ef := range efs {
		err := ef.CreateElement(repo, view)
		if err != nil {
			return nil, err
		}
	}

	err := inst.initElements(repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (inst *LocalRepoFactory) initElements(repo *Repo) error {
	elist := repo.Elements
	for _, o := range elist {
		el, ok := o.(Element)
		if ok {
			err := el.InitElement()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
