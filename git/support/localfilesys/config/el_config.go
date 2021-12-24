package config

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// LocalGitConfigFactory ...
type LocalGitConfigFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *LocalGitConfigFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *LocalGitConfigFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitConfig struct{}

////////////////////////////////////////////////////////////////////////////////
