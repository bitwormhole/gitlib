package unit

import (
	"testing"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/git/store"
)

func TestRepoFinder(t *testing.T) {

	ctx := gitlib.Init(nil, nil)
	lib := store.GetLib(ctx)
	files := lib.FS()

	tmp := files.NewPath(t.TempDir())
	wd := tmp.GetChild("a/b/c")
	wd.Mkdirs(nil)

	////////////////////////////////

	finder := lib.RepositoryFinder()
	layout, err := finder.Find(wd)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
	t.Log(layout)
}
