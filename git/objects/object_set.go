package objects

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
)

// SetR ...
type SetR interface {
	Contains(oid git.ObjectID) bool
	ReadObjectHead(oid git.ObjectID) (*git.Object, error)
	OpenObjectReader(oid git.ObjectID) (*git.Object, io.ReadCloser, error)
}

// SetW ...
type SetW interface {
	WriteObject(head *git.Object, data io.Reader) (*git.Object, error)
}

// Set ...
type Set interface {
	SetR
	SetW
}
