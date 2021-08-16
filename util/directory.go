package util

import "github.com/bitwormhole/starter/io/fs"

type LocalNode interface {
	Path() fs.Path
}

type LocalDirectory interface {
	LocalNode
	IsDir() bool
}

type LocalFile interface {
	LocalNode
	IsFile() bool
}
