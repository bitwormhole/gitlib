package store

import (
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/data/dxo"
)

// Session ...
type Session interface {
	io.Closer

	// 取仓库接口
	GetRepository() RepositoryProfile

	// 根据名称，取指定的组件
	GetComponent(name string) (any, error)

	// 取工作目录
	GetWD() afs.Path

	// 取仓库布局
	GetLayout() RepositoryLayout

	// config
	// SaveConfig(cfg Config) error
	// LoadConfig(cfg Config) error

	// objects
	LoadCommit(id dxo.ObjectID) (*dxo.Commit, error)
	LoadTag(id dxo.ObjectID) (*dxo.Tag, error)
	LoadTree(id dxo.ObjectID) (*dxo.Tree, error)

	// HEAD
	LoadHEAD(head HEAD) (dxo.ReferenceName, error)
}

// SessionFactory ...
type SessionFactory interface {
	OpenSession(profile RepositoryProfile) (Session, error)
}
