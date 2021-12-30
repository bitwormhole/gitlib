package views

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// WorkingDirFactory ...
type WorkingDirFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *WorkingDirFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *WorkingDirFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {
	path := v.Working.Directory
	el := &workingDir{}
	el.dir = path
	r.Working = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type workingDir struct {
	dir   fs.Path
	shell repository.Shell
}

func (inst *workingDir) _Impl() (repository.WorkingDirectory, localfilesys.Element) {
	return inst, inst
}

func (inst *workingDir) InitElement(r *localfilesys.Repo) error {
	inst.shell = r.Shell
	return nil
}

func (inst *workingDir) GetShell() repository.Shell {
	return inst.shell
}

func (inst *workingDir) GetDirectory() fs.Path {
	return inst.dir
}

////////////////////////////////////////////////////////////////////////////////
