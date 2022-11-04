package main

import (
	"github.com/bitwormhole/gitlib/libgit"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.SetArguments([]string{"--application.profiles.active=git-status"})
	i.UseMain(libgit.ModuleDemo())
	i.Run()
}
