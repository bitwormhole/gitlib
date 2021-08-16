package repository

import "github.com/bitwormhole/gitlib/util"

// GitPWD 表示一个viewport的当前工作目录
type GitPWD interface {
	util.LocalDirectory
}
