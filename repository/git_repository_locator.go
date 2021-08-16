package repository

import (
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
)

// Location 表示一个仓库视图的位置
type Location struct {
	PWD    fs.Path
	DotGit fs.Path

	WorkingDirectory   fs.Path
	SubmoduleDirectory fs.Path
	WorktreeDirectory  fs.Path

	CoreDirectory fs.Path
}

// Locator 是仓库的定位器
type Locator interface {
	Locate(uri lang.URI) (*Location, error)
	Accept(uri lang.URI) bool
}
