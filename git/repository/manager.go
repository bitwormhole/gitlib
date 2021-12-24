package repository

import (
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
)

// Manager 仓库管理器
type Manager interface {
	Open(uri lang.URI) (View, error)
	OpenByPath(path fs.Path) (View, error)
}
