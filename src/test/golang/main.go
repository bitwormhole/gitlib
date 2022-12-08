package main

import (
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/libgit"
	"github.com/bitwormhole/starter"
)

func main() {
	runWithModule()

	// runWithoutModule()
}

func runWithModule() {
	i := starter.InitApp()
	// i.SetArguments(  )
	i.UseMain(libgit.ModuleTest())
	i.Run()
}

func runWithoutModule() {
	lib := gitlib.GetLib()
	path := lib.FS().NewPath(".")
	layout, err := lib.RepositoryLocator().Locate(path)
	if err != nil {
		vlog.Error(err)
	} else if layout == nil {
		vlog.Error("layout is nil")
	}
}
