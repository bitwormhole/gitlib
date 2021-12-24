package main

import (
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(gitlib.Module())
	i.Use(GitlibTestModule())
	i.Run()
}
