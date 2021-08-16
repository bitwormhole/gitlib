package repository

import "github.com/bitwormhole/starter/io/fs"

// Finder 是仓库的搜索器
type Finder interface {
	Find(path fs.Path) []fs.Path
}
