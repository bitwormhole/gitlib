package unit

import (
	"testing"
)

func TestRepoFinder(t *testing.T) {

	unit := initUnit(t)

	tmp := unit.tmp
	wd := tmp.GetChild("../..")
	wd.Mkdirs(nil)

	////////////////////////////////

	finder := unit.lib.RepositoryFinder()
	all, err := finder.Find(wd)
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range all {
		t.Log("find repository ", p.GetPath())
	}
	t.Log("OK")
}
