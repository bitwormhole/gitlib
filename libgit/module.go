package libgit

import (
	"bitwormhole.com/starter/cli/libcli"
	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/gen/gitlibcfg"
	"github.com/bitwormhole/gitlib/gen/gitlibdemo"
	"github.com/bitwormhole/starter/application"
)

// Module ...
func Module() application.Module {
	mb := &application.ModuleBuilder{}
	mb = gitlib.InitModule(mb, false)

	mb.OnMount(gitlibcfg.ExportConfigForGitlib)

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleDemo ...
func ModuleDemo() application.Module {

	parent := Module()
	mb := &application.ModuleBuilder{}
	mb = gitlib.InitModule(mb, true)

	mb.Name(parent.GetName() + "#demo")
	mb.OnMount(gitlibdemo.ExportConfigForGitlibDemo)

	mb.Dependency(parent)
	mb.Dependency(libcli.ModuleCommon())

	return mb.Create()
}
