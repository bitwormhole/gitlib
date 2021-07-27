package repository

import "github.com/bitwormhole/gitlib/util"

type GitWorktree interface {
	util.LocalDirectory
	GitContextClient
}
