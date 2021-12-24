package repository

import "github.com/bitwormhole/starter/io/fs"

// WorkingDirectory 仓库的工作目录
type WorkingDirectory interface {
	GetShell() Shell
	GetDirectory() fs.Path
}
