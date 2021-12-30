package refs

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository/refs"
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

	r.Refs = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitRefs struct {
	dir fs.Path
}

func (inst *localGitRefs) _Impl() refs.Directory {
	return inst
}

func (inst *localGitRefs) GetRef(name git.ReferenceName) git.Ref {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
