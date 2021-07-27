package repository

import "github.com/bitwormhole/gitlib/util"

type GitPWD interface {
	util.LocalDirectory
	GitContextClient
}
