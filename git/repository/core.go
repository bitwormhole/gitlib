package repository

import "github.com/bitwormhole/starter/io/fs"

type Core interface {
	GetDirectory() fs.Path
}
