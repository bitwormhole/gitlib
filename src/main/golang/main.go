package main

import (
	"os"

	"github.com/bitwormhole/gitlib/modules/gitlib"
	"github.com/starter-go/starter"
)

func main() {

	a := os.Args
	m := gitlib.Module()
	i := starter.Init(a)
	i.MainModule(m)
	i.WithPanic(true).Run()

}
