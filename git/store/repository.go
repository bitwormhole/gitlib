package store

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

// RepositoryProfile 表示一个存在的git仓库的视图
type RepositoryProfile interface {
	Layout() RepositoryLayout

	Config() ConfigChain

	HEAD() HEAD

	Index() Index

	Refs() Refs

	Objects() Objects

	Digest() git.Digest           // default="SHA-1"
	Compression() git.Compression // default="DEFLATE"
	PathMapping() git.PathMapping // default="xx/xxxx"

	OpenSession() (Session, error)
}

// Repository  ...
type Repository interface {
	RepositoryProfile
}

// RepositoryLoader ...
type RepositoryLoader interface {
	Load(l RepositoryLayout) (Repository, error)
	LoadWithPath(path afs.Path) (Repository, error)
}
