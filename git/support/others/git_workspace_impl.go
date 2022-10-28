package others

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitWorkspaceImpl ...
type GitWorkspaceImpl struct {
	Core *store.Core
}

func (inst *GitWorkspaceImpl) _Impl() store.Workspace {
	return inst
}

// Path ...
func (inst *GitWorkspaceImpl) Path() afs.Path {
	panic("no impl")
}
