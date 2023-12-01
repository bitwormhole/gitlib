package repositories

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/afs"
)

// Entity ...
type Entity interface {
	OpenReader() (io.ReadCloser, error)
}

// Object ... 建议直接使用 git.Object
type Object git.Object

////////////////////////////////////////////////////////////////////////////////

// InfoFolder 代表 {.git}/objects/info
type InfoFolder interface {
	Path() afs.Path
}

// Objects 接口代表 {.git}/objects
type Objects interface {
	Path() afs.Path

	GetSparseObject(oid git.ObjectID) SparseObject

	GetPack(pid git.PackID) Pack

	ListPacks() []git.PackID

	Info() InfoFolder
}
