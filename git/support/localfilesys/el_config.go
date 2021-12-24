package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

type LocalGitConfigFactory struct {
	markup.Component `class:""`
}

func (inst *LocalGitConfigFactory) _Impl() ElementFactory {
	return inst
}

func (inst *LocalGitConfigFactory) InitElement(r *Repo, v *files.RepositoryView) error {

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitConfig struct{}

////////////////////////////////////////////////////////////////////////////////
