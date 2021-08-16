package repository

import (
	"github.com/bitwormhole/starter/lang"
)

// Driver 是仓库驱动
type Driver interface {
	Factory() Factory
	Locator() Locator
	Accept(uri lang.URI) bool
}

// Manager 是仓库管理器
type Manager interface {
	GetAllDrivers() []Driver
	FindDriver(uri lang.URI) (Driver, error)
	Open(uri lang.URI) (Viewport, error)
}
