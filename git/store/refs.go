package store

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/data/dxo"
)

////////////////////////////////////////////////////////////////////////////////

// Ref  is the key-value for .git/refs/*
type Ref interface {
	Path() afs.Path

	Name() dxo.ReferenceName

	Exists() bool

	GetValue(s Session) (dxo.ObjectID, error)
}

// Refs  is the key-value for .git/refs/*
type Refs interface {
	Path() afs.Path

	GetRef(name dxo.ReferenceName) Ref
}
