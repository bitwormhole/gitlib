package store

import (
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

// SparseObject  表示稀疏对象
type SparseObject interface {
	Path() afs.Path

	GetID() git.ObjectID

	Exists() bool
}

// SparseObjectLS 读写稀疏对象
type SparseObjectLS interface {
	ReadSparseObject(o SparseObject) (io.ReadCloser, *Object, error)
	ReadSparseObjectRaw(o SparseObject) (io.ReadCloser, error)
	WriteSparseObject(o *Object, data io.Reader) (*Object, error)
	WriteSparseObjectRaw(o *Object, data io.Reader) (*Object, error)
}
