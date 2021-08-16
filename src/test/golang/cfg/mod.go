package cfg

import (
	"github.com/bitwormhole/gitlib/gitstarter"
	srctest "github.com/bitwormhole/gitlib/src/test"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
)

func ExportModule() application.Module {

	builder := &application.ModuleBuilder{}
	builder.Name("gitlib/src/test/golang").Version("v1").Revision(0)

	builder.Resources(srctest.ExportResources())
	builder.OnMount(func(cb application.ConfigBuilder) error { return ExportConfig(cb) })
	builder.Dependency(starter.Module())
	builder.Dependency(gitstarter.Module())

	return builder.Create()
}
