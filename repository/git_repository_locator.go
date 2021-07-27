package repository

import "github.com/bitwormhole/starter/io/fs"

type GitRepositoryLocator interface {
	Locate(path fs.Path) (fs.Path, error)
}
