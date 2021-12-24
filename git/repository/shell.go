package repository

import "github.com/bitwormhole/starter/io/fs"

// Shell 提供git仓库的某个分支视图接口
type Shell interface {
	GetCore() Core

	GetDirectory() fs.Path

	GetHEAD() HEAD

	GetIndex() Index
}
