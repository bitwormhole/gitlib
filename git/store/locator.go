package store

import "bitwormhole.com/starter/afs"

// RepositoryLocator 表示一个git仓库定位器，用来确定仓库的准确位置
type RepositoryLocator interface {
	Locate(pwd afs.Path) (RepositoryLayout, error)
}

// RepositoryFinder 表示一个git仓库查找器，用来查找指定路径下的所有仓库
type RepositoryFinder interface {
	Find(pwd afs.Path) ([]RepositoryLayout, error)
}
