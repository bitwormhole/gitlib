package repository

import "github.com/bitwormhole/gitlib/util"

// GitWorkspace aka WorkingDirectory
type GitWorkspace interface {
	util.LocalDirectory
}
