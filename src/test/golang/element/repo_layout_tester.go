package element

import (
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/tests"
)

type RepoLayoutTester struct {
	markup.Component `id:"repo-layout-tester"`

	AppContext application.Context
	RM         repository.Manager
}

func (inst *RepoLayoutTester) Start() error {

	tc := tests.ContextForApp(inst.AppContext)

	loader := &TestingRepositoriesLoader{}
	loader.Init(inst.AppContext, tc)
	repos, err := loader.Load("test/data/repo.list.properties")
	if err != nil {
		return err
	}

	for _, info := range repos {
		err = inst.testWithPwd(info.PWD)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *RepoLayoutTester) testWithPwd(pwd fs.Path) error {

	return nil
}
