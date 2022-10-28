package store

import (
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/data/dxo"
)

// Entity ...
type Entity interface {
	OpenReader() (io.ReadCloser, error)
}

// Object 表示一个git对象
type Object struct {
	ID     dxo.ObjectID
	Type   string
	Length int64
	Entity Entity
}

// Pack 代表 {.git}/objects/pack
type Pack interface {
	GetID() dxo.PackID

	GetIndexFile() afs.Path

	GetEntityFile() afs.Path

	Exists() bool
}

// SparseObject  表示稀疏对象
type SparseObject interface {
	Path() afs.Path

	GetID() dxo.ObjectID

	Exists() bool
}

// InfoFolder 代表 {.git}/objects/info
type InfoFolder interface {
	Path() afs.Path
}

// Objects 接口代表 {.git}/objects
type Objects interface {
	Path() afs.Path

	GetObject(oid dxo.ObjectID) SparseObject

	GetPack(pid dxo.PackID) Pack

	Info() InfoFolder
}
