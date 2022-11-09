package gitlib

import (
	"embed"

	"bitwormhole.com/starter/cli/libcli"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName        = "github.com/bitwormhole/gitlib"
	theModuleVersion     = "v0.0.5"
	theModuleRevision    = 5
	theModuleResPath     = "src/main/resources"
	theModuleTestResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

//go:embed "src/test/resources"
var theModuleTestResFS embed.FS

// InitModule ...
func InitModule(mb *application.ModuleBuilder, forTest bool) *application.ModuleBuilder {

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
