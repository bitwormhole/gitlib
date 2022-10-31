package unit

import (
	"testing"
)

func TestRepoLocator(t *testing.T) {

	unit := initUnit(t)

	wd := unit.tmp.GetChild("a/b/c")
	wd.Mkdirs(nil)

	////////////////////////////////

	locator := unit.lib.RepositoryLocator()
	layout, err := locator.Locate(wd)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
	t.Log(layout)
}
