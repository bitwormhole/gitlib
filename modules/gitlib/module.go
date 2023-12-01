package gitlib

import (
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/gen/main4gitlib"
	"github.com/bitwormhole/gitlib/gen/test4gitlib"
	"github.com/starter-go/application"
)

// Module ...
func Module() application.Module {

	mb := gitlib.NewMainModule()
	mb.Components(main4gitlib.ExportComponents)

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	parent := Module()

	mb := gitlib.NewTestModule()
	mb.Components(test4gitlib.ExportComponents)

	mb.Depend(parent)
	return mb.Create()
}
