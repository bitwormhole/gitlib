package commands

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/store"
)

func initMeta(meta *store.Meta, t *cli.Task) {
	lib, err := store.GetLib(t.Context)
	if err != nil {
		panic(err)
	}
	meta.Context = t.Context
	meta.WD = lib.FS().NewPath(t.WD)
}
