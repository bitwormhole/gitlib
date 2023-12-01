package unit

import (
	"context"
	"testing"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/libgit"
	"github.com/bitwormhole/starter"
	"github.com/starter-go/afs"
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
	lib := initLib()
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

func initLib() store.Lib {

	mod := libgit.Module()
	i := starter.InitApp()
	i.UseMain(mod)
	rt, err := i.RunEx()
	if err != nil {
		panic(err)
	}

	ctx := rt.Context()
	o, err := ctx.GetComponent("#git-lib-agent")
	if err != nil {
		panic(err)
	}

	la := o.(store.LibAgent)
	lib, err := la.GetLib()
	if err != nil {
		panic(err)
	}

	return lib
}
