package store

import (
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
)

// ImportPackParams ...
type ImportPackParams struct {
	ID        git.PackID
	PackFile  afs.Path
	IdxFile   afs.Path
	MoveFiles bool
}

// ImportPackResult ...
type ImportPackResult struct {
	Params *ImportPackParams
	ID     git.PackID
	Pack   Pack
}

// // PackIndexItem ... 建议直接使用 git.PackIndexItem
// type PackIndexItem git.PackIndexItem

////////////////////////////////////////////////////////////////////////////////

// Pack 代表 {.git}/objects/pack
type Pack interface {
	GetID() git.PackID

	GetIndexFile() afs.Path

	GetEntityFile() afs.Path

	GetObject(oid git.ObjectID) PackObject

	Exists() bool
}

// PackObject 表示包内对象
type PackObject interface {
	GetID() git.ObjectID
	Container() Pack
}

// PackBuilder 读写包对象
type PackBuilder interface {
	io.Closer

	WriteObject(o *Object, data io.Reader) (*Object, error)
}

// PackIndex 表示包内对象索引
type PackIndex interface {
	Count() int64

	Get(index int64) (*git.PackIndexItem, error)

	List(index, limit int64) ([]*git.PackIndexItem, error)

	Find(id git.ObjectID) (*git.PackIndexItem, error)
}

// PackFile 表示包内对象索引
type PackFile interface {
	ReadPackObject(item *git.PackIndexItem) (io.ReadCloser, *git.Object, error)
}

// PackReadCloser 表示包内对象索引
type PackReadCloser interface {
	io.Closer
	GetIndex() PackIndex
	GetPack() PackFile
}

// PackDAO 读写包对象
type PackDAO interface {

	// query

	FindPackObject(o *git.PackIndexItem) (*git.PackIndexItem, error)

	ReadPackObject(o *git.PackIndexItem) (io.ReadCloser, error)

	CheckPack(pid git.PackID, flags pack.CheckFlag) error

	ListPacks() ([]git.PackID, error)

	// OpenPackReader(p Pack) (PackReadCloser, error)

	// insert

	NewPackBuilder() PackBuilder

	// import

	ImportPack(p *ImportPackParams) (*ImportPackResult, error)
}
