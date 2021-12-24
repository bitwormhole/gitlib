package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/starter/markup"
)

// LocalRepoLayout 组件提供本地仓库的布局
type LocalRepoLayout struct {
	markup.Component `id:"git-local-repository-layout"`
}

func (inst *LocalRepoLayout) _Impl() files.Layout {
	return inst
}

// MakeView 创建布局视图
func (inst *LocalRepoLayout) MakeView(location *files.RepositoryLocation) (*files.RepositoryView, error) {

	core := &files.CoreDirectory{}
	view := &files.RepositoryView{}

	wkdir := location.WorkingDirectory
	shdir := location.ShellDirectory

	if wkdir != nil && shdir != nil {
		shell := &files.ShellDirectory{}
		shell.Core = core
		shell.Directory = shdir

		shell.CommonDir = shdir.GetChild("commondir")
		shell.GitDir = shdir.GetChild("gitdir")
		shell.Head = shdir.GetChild("HEAD")
		shell.Index = shdir.GetChild("index")
		shell.Logs = shdir.GetChild("logs")
		shell.OrigHead = shdir.GetChild("ORIG_HEAD")

		working := &files.WorkingDirectory{}
		working.Current = location.Current
		working.Directory = wkdir
		working.DotGit = location.DotGit
		working.Shell = shell

		view.Shell = shell
		view.Working = working
	}

	coredir := location.CoreDirectory
	core.Config = location.ConfigFile
	core.Directory = coredir
	core.Objects = coredir.GetChild("objects")
	core.Refs = coredir.GetChild("refs")

	view.Core = core
	return view, nil
}
