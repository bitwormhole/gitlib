package gitlibdemo

import "github.com/bitwormhole/starter/application"

// ExportConfigForGitlibDemo ...
func ExportConfigForGitlibDemo(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
