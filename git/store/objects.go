package store

import (
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

// Entity ...
type Entity interface {
	OpenReader() (io.ReadCloser, error)
}

// Object 表示一个git对象
type Object struct {
	ID     git.ObjectID
	Type   string
	Length int64
	// Entity Entity
}

// Pack 代表 {.git}/objects/pack
type Pack interface {
	GetID() git.PackID

	GetIndexFile() afs.Path

	GetEntityFile() afs.Path

	GetObject(oid git.ObjectID) PackObject

	Exists() bool
}

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

// PackObject 表示包内对象
type PackObject interface {
	GetID() git.ObjectID
	Container() Pack
}

// PackBuilder 读写包对象
type PackBuilder interface {
}

// PackObjectLS 读写包对象
type PackObjectLS interface {
	ReadPackObject(o PackObject) (io.ReadCloser, *Object, error)
}

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

	Info() InfoFolder
}
