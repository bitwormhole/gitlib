package util

import "github.com/bitwormhole/starter/io/fs"

type LocalNode interface {
	Path() fs.Path
	IsFile() bool
	IsDirectory() bool
	Exists() bool
}

type LocalDirectory interface {
	LocalNode
	List() []string
}

type LocalFile interface {
	LocalNode
	Length() int64
}
