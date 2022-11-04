package gitlibcfg

import "github.com/bitwormhole/starter/application"

// ExportConfigForGitlib ...
func ExportConfigForGitlib(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
