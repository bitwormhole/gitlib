package repository

import "github.com/bitwormhole/gitlib/util"

// ref to the [.git] directory, or not if bare.
type GitRepository interface {
	util.LocalDirectory
}
