package unit

import (
	"testing"

	"bitwormhole.com/starter/cli"
)

func TestRepoLocator(t *testing.T) {

	unit := initUnit(t)
	repoName := "foo"

	wd := unit.tmp.GetChild("a/b/c")
	wd.Mkdirs(nil)
	wd2 := wd.GetChild(repoName)

	client := unit.lib.GetCLI(true).GetClient()
	err := client.Run(&cli.Task{
		Context: unit.context,
		Command: "git init " + repoName,
		WD:      wd.GetPath(),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	////////////////////////////////

	locator := unit.lib.RepositoryLocator()
	layout, err := locator.Locate(wd2)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("OK")
	t.Log(layout)
}
