package gitlib

import (
	"github.com/bitwormhole/gitlib/git/repositories"
)

// Agent 提供一个获取 repositories.Lib 的便捷接口
type Agent interface {
	GetLib() repositories.Lib
}
