package views

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// RepoViewFactory ...
type RepoViewFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *RepoViewFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *RepoViewFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {

	core := v.Core
	sh := v.Shell
	wk := v.Working
	el := &repoView{}

	if wk != nil {
		el._current = wk.Current
		el._dir = wk.DotGit
	} else if sh != nil {
		el._dir = sh.Directory
	} else if core != nil {
		el._dir = core.Directory
	}

	r.View = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type repoView struct {
	_dir     fs.Path
	_current fs.Path
	_core    repository.Core
	_shell   repository.Shell
	_wkdir   repository.WorkingDirectory
}

func (inst *repoView) _Impl() (repository.View, localfilesys.Element) {
	return inst, inst
}

func (inst *repoView) InitElement(r *localfilesys.Repo) error {
	inst._core = r.Core
	inst._shell = r.Shell
	inst._wkdir = r.Working
	return nil
}

func (inst *repoView) GetCore() repository.Core {
	return inst._core
}

func (inst *repoView) GetShell() repository.Shell {
	return inst._shell
}

func (inst *repoView) GetWorking() repository.WorkingDirectory {
	return inst._wkdir
}

func (inst *repoView) GetCurrent() string {
	return inst._current.Path()
}

func (inst *repoView) GetCurrentPath() fs.Path {
	return inst._current
}

func (inst *repoView) IsBare() bool {
	return inst._wkdir == nil
}

func (inst *repoView) OpenSession() (repository.Session, error) {
	return nil, errors.New("no impl")
}

////////////////////////////////////////////////////////////////////////////////
