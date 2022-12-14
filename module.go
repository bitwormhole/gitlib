package gitlib

import (
	"embed"

	"bitwormhole.com/starter/cli/libcli"
	"github.com/bitwormhole/gitlib/gen/gitlibcfg"
	"github.com/bitwormhole/gitlib/gen/gitlibdemo"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwormhole/gitlib"
	theModuleVersion  = "v0.0.12"
	theModuleRevision = 12

	theModuleResPath     = "src/main/resources"
	theModuleTestResPath = "src/test/resources"
)

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module ...
func Module() application.Module {
	mb := &application.ModuleBuilder{}
	mb = initModule(mb, false)

	mb.OnMount(gitlibcfg.ExportConfigForGitlib)

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/test/resources"
var theModuleTestResFS embed.FS

// ModuleTest ...
func ModuleTest() application.Module {

	parent := Module()
	mb := &application.ModuleBuilder{}
	mb = initModule(mb, true)

	mb.Name(parent.GetName() + "#demo")
	mb.OnMount(gitlibdemo.ExportConfigForGitlibDemo)

	mb.Dependency(parent)
	mb.Dependency(libcli.ModuleCommon())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// InitModule ...
func initModule(mb *application.ModuleBuilder, forTest bool) *application.ModuleBuilder {

	if mb == nil {
		mb = &application.ModuleBuilder{}
	}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)

	if !forTest {
		mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))
	} else {
		mb.Resources(collection.LoadEmbedResources(&theModuleTestResFS, theModuleTestResPath))
	}

	mb.Dependency(starter.Module())
	mb.Dependency(libcli.Module())

	return mb
}
