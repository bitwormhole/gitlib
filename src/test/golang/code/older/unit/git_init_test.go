package unit

import (
	"testing"

	"bitwormhole.com/starter/cli"
)

func TestGitInit(t *testing.T) {

	unit := initUnit(t)

	tmp := unit.tmp
	wd := tmp.GetChild("test_git_init")
	wd.Mkdirs(nil)

	ctx := unit.context

	////////////////////////////////

	task := &cli.Task{
		Context: ctx,
		Command: "git init foo",
		WD:      wd.GetPath(),
	}
	err := unit.cli.GetClient().Run(task)
	if err != nil {
		t.Error(err)
	}
}
