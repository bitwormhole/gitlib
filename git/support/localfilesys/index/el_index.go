package index

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// GitIndexFileFactory ...
type GitIndexFileFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *GitIndexFileFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *GitIndexFileFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {
	file := v.Shell.Index
	el := &gitIndexFile{}
	el._file = file
	r.Index = el
	r.AddElement(el)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type gitIndexFile struct {
	_file fs.Path
}

func (inst *gitIndexFile) _Impl() (repository.Index, localfilesys.Element) {
	return inst, inst
}

func (inst *gitIndexFile) InitElement(r *localfilesys.Repo) error {

	return nil
}

func (inst *gitIndexFile) Demo() int {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
