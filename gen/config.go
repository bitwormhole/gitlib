package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigForGitlib(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
