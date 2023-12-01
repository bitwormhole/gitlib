package repositories

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects"
	"github.com/starter-go/afs"
)

// Session ...
type Session interface {
	io.Closer

	// 取仓库接口
	GetRepository() Repository

	// 根据名称，取指定的组件
	GetComponent(name string) (any, error)

	// 取工作目录
	GetWD() afs.Path

	// 取仓库布局
	GetLayout() Layout

	// 新建临时文件
	NewTemporaryFile(dir afs.Path) afs.Path

	NewTemporaryBuffer(dir afs.Path) TemporaryBuffer

	NewReaderPool(size int) afs.ReaderPool

	GetReaderPool() afs.ReaderPool

	GetObjectContext() *objects.Context

	/////////////////////////////////////////////////

	// config
	// SaveConfig(cfg Config) error
	// LoadConfig(cfg Config) error

	// objects
	ReadObject(id git.ObjectID) (*git.Object, io.ReadCloser, error)

	WriteObject(o *git.Object, data io.Reader) (*git.Object, error)

	GetSparseObjects() SparseObjects

	GetPacks() Packs

	LoadText(id git.ObjectID) (string, *git.Object, error)
	LoadBinary(id git.ObjectID) ([]byte, *git.Object, error)

	// commit, tag, tree
	LoadCommit(id git.ObjectID) (*git.Commit, error)
	LoadTag(id git.ObjectID) (*git.Tag, error)
	LoadTree(id git.ObjectID) (*git.Tree, error)

	// refs
	LoadRef(r Ref) (*git.Ref, error)
	SaveRef(r *git.Ref) error

	// HEAD
	LoadHEAD(head HEAD) (*git.HEAD, error)
	SaveHEAD(h *git.HEAD) error
}

// SessionFactory ...
type SessionFactory interface {
	OpenSession(repo Repository) (Session, error)
}
