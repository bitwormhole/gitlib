package libgit

import (
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/starter/application"
)

// Module ...
func Module() application.Module {
	return gitlib.Module()
}

// ModuleTest ...
func ModuleTest() application.Module {
	return gitlib.ModuleTest()
}
