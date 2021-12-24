package objects

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// LocalGitObjectsFactory ...
type LocalGitObjectsFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *LocalGitObjectsFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *LocalGitObjectsFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitObjects struct{}

////////////////////////////////////////////////////////////////////////////////
