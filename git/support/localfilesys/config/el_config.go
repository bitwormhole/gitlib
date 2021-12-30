package config

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
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

	file := v.Core.Config

	el := &localGitConfig{}
	el._file = file

	r.Config = el
	r.AddElement(el)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitConfig struct {
	_file fs.Path
}

func (inst *localGitConfig) _Impl() (repository.Config, localfilesys.Element) {
	return inst, inst
}

func (inst *localGitConfig) InitElement(r *localfilesys.Repo) error {
	return nil
}

func (inst *localGitConfig) Demo() int {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
