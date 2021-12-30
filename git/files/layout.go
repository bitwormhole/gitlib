package files

import "github.com/bitwormhole/starter/io/fs"

// CoreDirectory 核心布局
type CoreDirectory struct {
	Directory fs.Path // self
	Config    fs.Path // the 'config' as file
	Objects   fs.Path // the 'objects' as dir
	Refs      fs.Path // the 'refs' as dir
	Modules   fs.Path // the 'modules' as dir
	Worktrees fs.Path // the 'worktrees' as dir
}

// ShellDirectory 壳子布局
type ShellDirectory struct {
	Core      *CoreDirectory // inner
	Directory fs.Path        // self
	Index     fs.Path        // the 'index' as file
	Head      fs.Path        // the 'HEAD' as file
	OrigHead  fs.Path        // the 'ORIG_HEAD' as file
	GitDir    fs.Path        // the 'gitdir' as file
	CommonDir fs.Path        // the 'commondir' as file
	Logs      fs.Path        // the 'logs' as dir
}

// WorkingDirectory 工作区布局
type WorkingDirectory struct {
	Shell     *ShellDirectory // inner
	Directory fs.Path         // self
	DotGit    fs.Path
	Current   fs.Path
}

// RepositoryView  仓库布局视图
type RepositoryView struct {
	Working *WorkingDirectory
	Shell   *ShellDirectory
	Core    *CoreDirectory
}

// Layout 仓库布局对象
type Layout interface {
	MakeView(location *RepositoryLocation) (*RepositoryView, error)
}
