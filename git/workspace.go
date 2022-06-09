package git

import "github.com/bitwormhole/gitlib/git/files"

// Workspace 表示工作区目录
type Workspace interface {
	Path() files.Path
}
