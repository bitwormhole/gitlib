package repositories

import "github.com/starter-go/afs"

// Locator 表示一个git仓库定位器，用来确定仓库的准确位置
type Locator interface {
	Locate(pwd afs.Path) (Layout, error)
}

// Finder 表示一个git仓库查找器，用来查找指定路径下的所有仓库
type Finder interface {
	Find(pwd afs.Path) ([]afs.Path, error)
}
