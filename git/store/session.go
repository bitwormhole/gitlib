package store

import (
	"io"

	"bitwormhole.com/starter/afs"
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

	SaveConfig(cfg Config) error

	LoadConfig(cfg Config) error
}

// SessionFactory ...
type SessionFactory interface {
	OpenSession(profile RepositoryProfile) (Session, error)
}
