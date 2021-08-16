package element

import (
	"github.com/bitwormhole/gitlib/repository"
)

type GitRepoTester struct {
	Path string             // `inject:"${test.repo.path}"`
	RM   repository.Manager // `inject:"#git-repository-manager"`
}

func (inst *GitRepoTester) Start() error {

	// pwd := fs.Default().GetPath(inst.Path)
	// vpt, err := inst.RM.Open(pwd)
	// if err != nil {
	// 	return err
	// }
	// defer vpt.Close()

	return nil
}
