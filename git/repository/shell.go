package repository

import "github.com/bitwormhole/starter/io/fs"

type Shell interface {
	GetCore() Core
	GetDirectory() fs.Path
}
