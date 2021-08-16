package element

import (
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/tests"
	"github.com/bitwormhole/starter/vlog"
)

type RepoLocatorTester struct {
	markup.Component `id:"repo-locator-tester"`

	AppContext application.Context
	RM         repository.Manager
}

func (inst *RepoLocatorTester) Start() error {

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

func (inst *RepoLocatorTester) testWithPwd(pwd fs.Path) error {

	vlog.Debug("test repoLocator with repo@", pwd.Path())

	uri := pwd.URI()
	driver, err := inst.RM.FindDriver(uri)
	if err != nil {
		return err
	}

	location, err := driver.Locator().Locate(uri)
	if err != nil {
		return err
	}

	vlog.Debug("location=", location)

	return nil
}
