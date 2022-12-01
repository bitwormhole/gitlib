package store

import (
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
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
	GetLayout() RepositoryLayout

	// 新建临时文件
	NewTemporaryFile(dir afs.Path) afs.Path

	NewTemporaryBuffer(dir afs.Path) TemporaryBuffer

	NewReaderPool(size int) afs.ReaderPool

	GetReaderPool() afs.ReaderPool

	/////////////////////////////////////////////////

	// config
	// SaveConfig(cfg Config) error
	// LoadConfig(cfg Config) error

	// objects
	ReadObject(id git.ObjectID) (io.ReadCloser, *Object, error)
	// WriteObject(o *Object) (io.WriteCloser, error)

	SparseObjectLS

	GetPacks() PackDAO

	LoadText(id git.ObjectID) (string, error)
	LoadBinary(id git.ObjectID) ([]byte, error)

	// commit, tag, tree
	LoadCommit(id git.ObjectID) (*git.Commit, error)
	LoadTag(id git.ObjectID) (*git.Tag, error)
	LoadTree(id git.ObjectID) (*git.Tree, error)

	// refs
	LoadRef(r Ref) (*git.Ref, error)

	// HEAD
	LoadHEAD(head HEAD) (*git.HEAD, error)
}

// SessionFactory ...
type SessionFactory interface {
	OpenSession(repo Repository) (Session, error)
}
