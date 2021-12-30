package modules

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository/modules"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// GitModulesDirFactory ...
type GitModulesDirFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *GitModulesDirFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *GitModulesDirFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {

	dir := v.Core.Modules

	el := &gitModulesDir{}
	el._dir = dir

	r.Modules = el
	r.AddElement(el)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type gitModulesDir struct {
	_dir fs.Path
}

func (inst *gitModulesDir) _Impl() (modules.Directory, localfilesys.Element) {
	return inst, inst
}

func (inst *gitModulesDir) InitElement(r *localfilesys.Repo) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
