package unit

import (
	"testing"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

func TestGitStatus(t *testing.T) {

	const name = "test_git_status"

	ctx := gitlib.Init(nil, nil)
	lib := store.GetLib(ctx)
	files := lib.FS()

	tmp := files.NewPath(t.TempDir())
	wd1 := tmp
	wd2 := wd1.GetChild(name)

	////////////////////////////////
	// git init

	task1 := git.NewInit(ctx)
	task1.WD = wd1
	task1.Directory = name

	err := task1.Run()
	if err != nil {
		t.Error(err)
	}

	////////////////////////////////
	// git status

	task2 := git.NewStatus(ctx)
	task2.WD = wd2

	err = task2.Run()
	if err != nil {
		t.Error(err)
	}
}
