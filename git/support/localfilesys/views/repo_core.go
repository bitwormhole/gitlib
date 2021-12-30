package views

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/repository/modules"
	"github.com/bitwormhole/gitlib/git/repository/objects"
	"github.com/bitwormhole/gitlib/git/repository/refs"
	"github.com/bitwormhole/gitlib/git/repository/worktrees"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// CoreDirFactory ...
type CoreDirFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *CoreDirFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *CoreDirFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {
	path := v.Core.Directory
	el := &repoCoreDir{}
	el._dir = path
	r.Core = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type repoCoreDir struct {
	_dir       fs.Path
	_refs      refs.Directory
	_objects   objects.Directory
	_worktrees worktrees.Directory
	_modules   modules.Directory
	_config    repository.Config
}

func (inst *repoCoreDir) _Impl() (repository.Core, localfilesys.Element) {
	return inst, inst
}

func (inst *repoCoreDir) InitElement(r *localfilesys.Repo) error {
	inst._config = r.Config
	inst._modules = r.Modules
	inst._objects = r.Objects
	inst._refs = r.Refs
	inst._worktrees = r.Worktrees
	return nil
}

func (inst *repoCoreDir) GetDirectory() fs.Path {
	return inst._dir
}

func (inst *repoCoreDir) GetRefs() refs.Directory {
	return inst._refs
}

func (inst *repoCoreDir) GetObjects() objects.Directory {
	return inst._objects
}

func (inst *repoCoreDir) GetWorktrees() worktrees.Directory {
	return inst._worktrees
}

func (inst *repoCoreDir) GetModules() modules.Directory {
	return inst._modules
}

func (inst *repoCoreDir) GetConfig() repository.Config {
	return inst._config
}

////////////////////////////////////////////////////////////////////////////////
