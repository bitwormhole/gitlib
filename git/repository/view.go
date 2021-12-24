package repository

import "github.com/bitwormhole/starter/io/fs"

// View 仓库的视图
type View interface {
	GetCore() Core
	GetShell() Shell
	GetWorking() WorkingDirectory
	GetCurrent() string
	GetCurrentPath() fs.Path
	IsBare() bool
	OpenSession() (Session, error)
}

// ViewFactory 用来创建View
type ViewFactory interface {
	CreateView(path fs.Path) (View, error)
}
