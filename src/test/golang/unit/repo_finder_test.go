package unit

import (
	"testing"
)

func TestRepoFinder(t *testing.T) {

	unit := initUnit(t)

	tmp := unit.tmp
	wd := tmp.GetChild("a/b/c")
	wd.Mkdirs(nil)

	////////////////////////////////

	finder := unit.lib.RepositoryFinder()
	layout, err := finder.Find(wd)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
	t.Log(layout)
}
