package main

import (
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/src/test/golang/gen"

	"github.com/bitwormhole/starter/application"
)

func GitlibTestModule() application.Module {

	parent := gitlib.Module()
	mb := application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#unit-test").Version(parent.GetVersion()).Revision(parent.GetRevision())
	mb.OnMount(gen.ExportConfigForGitlibTest)
	mb.Dependency(parent)
	return mb.Create()
}
