package store

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

////////////////////////////////////////////////////////////////////////////////

// Ref  is the key-value for .git/refs/*
type Ref interface {
	Path() afs.Path

	Name() git.ReferenceName

	Exists() bool

	GetValue(s Session) (git.ObjectID, error)
}

// Refs  is the key-value for .git/refs/*
type Refs interface {
	Path() afs.Path

	GetRef(name git.ReferenceName) Ref

	List() []git.ReferenceName
}
