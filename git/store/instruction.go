package store

import (
	"context"

	"bitwormhole.com/starter/afs"
)

// Meta 包含git-command-object 的基本字段
type Meta struct {
	Context context.Context
	Name    string   // the command name
	WD      afs.Path // Working Directory
}

// Task 接口表示git-command-object 的外观
type Task interface {
	GetMeta() *Meta
	Run() error
}

// Service is the Runner for Command
type Service interface {
	// RunTask(t Task) error
}

// ServiceRegistration ...
type ServiceRegistration struct {
	Name    string
	Service Service
}

// ServiceRegistry ...
// [inject:".git-instruction-registry"]
type ServiceRegistry interface {
	ListRegistrations() []*ServiceRegistration
}

// ServiceManager 用来管理已注册的服务
type ServiceManager interface {
	Find(name string) (Service, error)
}
