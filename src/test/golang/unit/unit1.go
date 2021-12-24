package unit

import (
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

type UnitTest1 struct {
	markup.Component `initMethod:"Init"`

	RM repository.Manager `inject:"#git-repository-manager"`

	RepoPath string `inject:"${test.repo.path}"`
}

func (inst *UnitTest1) Init() error {

	path := fs.Default().GetPath(inst.RepoPath)

	view, err := inst.RM.OpenByPath(path)
	if err != nil {
		return err
	}

	view.GetCore()

	return nil
}
