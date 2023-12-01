package unit

import (
	"testing"

	"bitwormhole.com/starter/cli"
)

func disableTestGitCommit(t *testing.T) {

	const name = "test_git_commit"

	unit := initUnit(t)
	wd1 := unit.tmp
	wd2 := wd1.GetChild(name)
	ctx := unit.context

	////////////////////////////////
	// git init

	task := &cli.Task{
		Context: ctx,
		Command: "git init " + name,
		WD:      wd1.GetPath(),
	}
	err := unit.cli.GetClient().Run(task)
	if err != nil {
		t.Error(err)
	}
	ctx = task.Context

	////////////////////////////////
	// git commit

	err = unit.cli.GetClient().Run(&cli.Task{
		Context: ctx,
		Command: "git commit -m 'test'",
		WD:      wd2.GetPath(),
	})
	if err != nil {
		t.Error(err)
	}

}
