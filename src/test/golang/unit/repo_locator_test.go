package unit

import (
	"testing"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/git/store"
)

func TestRepoLocator(t *testing.T) {

	ctx := gitlib.Init(nil, nil)
	lib := store.GetLib(ctx)

	files := lib.FS()
	// ctx := lib.Context()

	tmp := files.NewPath(t.TempDir())
	wd := tmp.GetChild("a/b/c")
	wd.Mkdirs(nil)

	////////////////////////////////

	locator := lib.RepositoryLocator()
	layout, err := locator.Locate(wd)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
	t.Log(layout)
}
