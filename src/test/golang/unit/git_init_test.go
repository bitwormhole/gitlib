package unit

import (
	"testing"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

func TestGitInit(t *testing.T) {

	ctx := gitlib.Init(nil, nil)
	lib := store.GetLib(ctx)
	files := lib.FS()

	tmp := files.NewPath(t.TempDir())
	wd := tmp.GetChild("test_git_init")
	wd.Mkdirs(nil)

	////////////////////////////////

	task := git.NewInit(ctx)
	task.Bare = false
	task.Directory = "demo"
	task.WD = wd

	err := task.Run()
	if err != nil {
		t.Error(err)
	}
}
