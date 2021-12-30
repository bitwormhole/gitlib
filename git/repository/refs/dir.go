package refs

import "github.com/bitwormhole/gitlib/git"

// Directory 表示【.git/refs】这个目录
type Directory interface {
	GetRef(name git.ReferenceName) git.Ref
}
