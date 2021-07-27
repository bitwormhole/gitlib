package repository

import "github.com/bitwormhole/starter/io/fs"

type GitRepositoryFinder interface {
	Find(path fs.Path) []fs.Path
}
