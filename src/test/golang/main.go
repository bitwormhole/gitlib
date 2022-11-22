package main

import (
	"github.com/bitwormhole/gitlib/libgit"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	// i.SetArguments(  )
	i.UseMain(libgit.ModuleDemo())
	i.Run()
}
