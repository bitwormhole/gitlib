package store

import "bitwormhole.com/starter/afs"

// ReferenceName is the name for .git/refs/*
type ReferenceName string

func (v ReferenceName) String() string {
	return string(v)
}

// Normalize ...
func (v ReferenceName) Normalize() ReferenceName {
	return "todo..."
}

////////////////////////////////////////////////////////////////////////////////

// Ref  is the key-value for .git/refs/*
type Ref interface {
	Path() afs.Path

	Name() ReferenceName

	Exists() bool
}

// Refs  is the key-value for .git/refs/*
type Refs interface {
	Path() afs.Path

	GetRef(name ReferenceName) Ref
}
