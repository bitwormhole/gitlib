package services

import (
	"context"

	"bitwormhole.com/starter/afs"
)

// Command 包含git-command-object 的基本字段
type Command struct {
	Context context.Context
	Name    string   // the command name
	WD      afs.Path // Working Directory
}

// Task 接口表示git-command-object 的外观
type Task interface {
	GetCommand() *Command
	Run() error
}
