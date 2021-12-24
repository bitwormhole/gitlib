package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigForGitlibTest(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
