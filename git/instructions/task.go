package instructions

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
