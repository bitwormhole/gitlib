package refs

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// LocalGitRefsFactory ...
type LocalGitRefsFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *LocalGitRefsFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *LocalGitRefsFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {

	el := &localGitRefs{}
	el.dir = v.Core.Refs

	r.AddElement(el)
	r.Refs = el

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitRefs struct {
	dir fs.Path
}

////////////////////////////////////////////////////////////////////////////////
