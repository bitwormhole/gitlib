package views

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// ShellDirFactory ...
type ShellDirFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *ShellDirFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *ShellDirFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {
	path := v.Shell.Directory
	el := &repoShellDir{}
	el._dir = path
	r.Shell = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type repoShellDir struct {
	_dir   fs.Path
	_core  repository.Core
	_head  repository.HEAD
	_index repository.Index
}

func (inst *repoShellDir) _Impl() (repository.Shell, localfilesys.Element) {
	return inst, inst
}

func (inst *repoShellDir) InitElement(r *localfilesys.Repo) error {
	inst._core = r.Core
	inst._head = r.Head
	inst._index = r.Index
	return nil
}

func (inst *repoShellDir) GetCore() repository.Core {
	return inst._core
}

func (inst *repoShellDir) GetDirectory() fs.Path {
	return inst._dir
}

func (inst *repoShellDir) GetHEAD() repository.HEAD {
	return inst._head
}

func (inst *repoShellDir) GetIndex() repository.Index {
	return inst._index
}

////////////////////////////////////////////////////////////////////////////////
