package worktrees

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/gitlib/git/repository/worktrees"
	"github.com/bitwormhole/gitlib/git/support/localfilesys"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// GitWorktreesDirFactory ...
type GitWorktreesDirFactory struct {
	markup.Component `class:"git-local-element-factory"`
}

func (inst *GitWorktreesDirFactory) _Impl() localfilesys.ElementFactory {
	return inst
}

// CreateElement ...
func (inst *GitWorktreesDirFactory) CreateElement(r *localfilesys.Repo, v *files.RepositoryView) error {

	dir := v.Core.Worktrees

	el := &gitWorktreesDir{}
	el._dir = dir

	r.Worktrees = el
	r.AddElement(el)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type gitWorktreesDir struct {
	_dir fs.Path
}

func (inst *gitWorktreesDir) _Impl() (worktrees.Directory, localfilesys.Element) {
	return inst, inst
}

func (inst *gitWorktreesDir) InitElement(r *localfilesys.Repo) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
