package main

import (
	"os"

	"github.com/bitwormhole/gitlib/modules/gitlib"
	"github.com/starter-go/units"
)

func main() {
	m := gitlib.ModuleForTest()
	units.NewRunner().Dependencies(m).Run(os.Args)
}
