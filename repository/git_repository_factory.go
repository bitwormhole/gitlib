package repository

import "github.com/bitwormhole/starter/io/fs"

type GitRepositoryFactory interface {
	Open(path fs.Path) (GitRepositoryViewport, error)
}
