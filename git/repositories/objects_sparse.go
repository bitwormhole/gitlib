package repositories

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/afs"
)

// SparseObject  表示稀疏对象
type SparseObject interface {
	Path() afs.Path

	GetID() git.ObjectID

	Exists() bool
}

// SparseObjects 读写稀疏对象
type SparseObjects interface {
	ReadSparseObject(o SparseObject) (*git.Object, io.ReadCloser, error)
	ReadSparseObjectRaw(o SparseObject) (io.ReadCloser, error)
	WriteSparseObject(o *git.Object, data io.Reader) (*git.Object, error)
	WriteSparseObjectRaw(o *git.Object, data io.Reader) (*git.Object, error)
}
