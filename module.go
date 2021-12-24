package gitlib

import (
	"embed"

	"github.com/bitwormhole/gitlib/gen"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/gitlib"
	myModuleVer  = "v0.0.1"
	myModuleRev  = 1
)

//go:embed src/main/resources
var theResFS embed.FS

// Module 导出模块【github.com/bitwormhole/gitlib】
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theResFS, "src/main/resources"))
	mb.OnMount(gen.ExportConfigForGitlib)

	mb.Dependency(starter.Module())

	return mb.Create()
}
