package repositories

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/afs"
)

// // RepositoryProfile 表示一个存在的git仓库的视图
// type RepositoryProfile interface {
// }

// Repository  ...
type Repository interface {
	Layout() Layout

	Config() ConfigChain

	HEAD() HEAD

	Index() Index

	Refs() Refs

	Objects() Objects

	Worktrees() Worktrees

	Submodules() Submodules

	Digest() git.Digest           // default="SHA-1"
	Compression() git.Compression // default="DEFLATE"
	PathMapping() git.PathMapping // default="xx/xxxx"

	OpenSession() (Session, error)
}

// Loader ...
type Loader interface {
	Load(l Layout) (Repository, error)
	LoadWithPath(path afs.Path) (Repository, error)
}
