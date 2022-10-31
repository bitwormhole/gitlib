package unit

import (
	"context"
	"testing"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/git/store"
)

type gitUnit struct {
	lib     store.Lib
	cli     cli.CLI
	context context.Context
	tmp     afs.Path
	fs      afs.FS
}

func initUnit(t *testing.T) *gitUnit {

	ctx := context.Background()
	lib := gitlib.New(nil)
	cli := lib.GetCLI(true)
	fs1 := lib.FS()

	ctx = lib.Bind(ctx)
	ctx = cli.Bind(ctx)

	tmp := fs1.NewPath(t.TempDir())

	unit := &gitUnit{}
	unit.context = ctx
	unit.fs = fs1
	unit.lib = lib
	unit.cli = cli
	unit.tmp = tmp
	return unit
}
