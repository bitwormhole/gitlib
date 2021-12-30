package repository

import "github.com/bitwormhole/gitlib/git"

// HEAD the file of [.git/HEAD]
type HEAD interface {
	GetValue() (git.ReferenceName, error)
	SetValue(name git.ReferenceName) error
	Exists() bool
}
