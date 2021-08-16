package gitstarter

import (
	etcgitlib "github.com/bitwormhole/gitlib/etc/gitlib"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
)

const (
	myName     = "github.com/bitwormhole/gitlib"
	myVersion  = "v0.0.1"
	myRevision = 1
)

// Module 导出模块：[github.com/bitwormhole/gitlib]
func Module() application.Module {

	builder := &application.ModuleBuilder{}
	builder.Name(myName).Version(myVersion).Revision(myRevision)
	builder.Dependency(starter.Module())

	builder.OnMount(func(cb application.ConfigBuilder) error { return etcgitlib.ExportConfig(cb) })

	return builder.Create()
}
