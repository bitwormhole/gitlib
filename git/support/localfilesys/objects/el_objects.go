package objects

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository/objects"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
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

	dir := v.Core.Objects

	el := &localGitObjects{}
	el._dir = dir

	r.Objects = el
	r.AddElement(el)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type localGitObjects struct {
	_dir fs.Path
}

func (inst *localGitObjects) _Impl() (objects.Directory, localfilesys.Element) {
	return inst, inst
}

func (inst *localGitObjects) InitElement(r *localfilesys.Repo) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
