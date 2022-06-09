package git

import "github.com/bitwormhole/gitlib/git/files"

// Repository 表示一个存在的git仓库
type Repository interface {
	Layout() Layout

	Config() RepositoryConfig

	HEAD() HEAD

	Index() Index

	Refs() Refs

	Objects() Objects
}

// RepositoryLocator 表示一个git仓库定位器，用来确定仓库的准确位置
type RepositoryLocator interface {
	Locate(pwd files.Path) (Layout, error)
}

// RepositoryFinder 表示一个git仓库查找器，用来查找指定路径下的所有仓库
type RepositoryFinder interface {
	Find(pwd files.Path) ([]Layout, error)
}

// RepositoryFactory 表示用来创建Repository对象的工厂
type RepositoryFactory interface {
	Open(l Layout) (Repository, error)
}
