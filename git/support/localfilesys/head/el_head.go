package head

import (
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// GitHeadFileFactory ...
type GitHeadFileFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *GitHeadFileFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *GitHeadFileFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {
	file := v.Shell.Head
	el := &headFile{}
	el._file = file
	r.Head = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type headFile struct {
	_file fs.Path
}

func (inst *headFile) _Impl() (repository.HEAD, localfilesys.Element) {
	return inst, inst
}

func (inst *headFile) InitElement(r *localfilesys.Repo) error {

	return nil
}

func (inst *headFile) GetValue() (git.ReferenceName, error) {
	return nil, nil
}

func (inst *headFile) SetValue(name git.ReferenceName) error {
	return nil
}

func (inst *headFile) Exists() bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////
