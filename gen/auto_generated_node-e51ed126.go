// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	files0x00707a "github.com/bitwormhole/gitlib/git/files"
	repository0x5aaf5a "github.com/bitwormhole/gitlib/git/repository"
	support0x074feb "github.com/bitwormhole/gitlib/git/support"
	localfilesys0x6be3ff "github.com/bitwormhole/gitlib/git/support/localfilesys"
	config0x71b4a2 "github.com/bitwormhole/gitlib/git/support/localfilesys/config"
	head0xb6393b "github.com/bitwormhole/gitlib/git/support/localfilesys/head"
	index0x3eb559 "github.com/bitwormhole/gitlib/git/support/localfilesys/index"
	modules0xe8c0dc "github.com/bitwormhole/gitlib/git/support/localfilesys/modules"
	objects0x11508a "github.com/bitwormhole/gitlib/git/support/localfilesys/objects"
	refs0x4e5472 "github.com/bitwormhole/gitlib/git/support/localfilesys/refs"
	views0x7d154b "github.com/bitwormhole/gitlib/git/support/localfilesys/views"
	worktrees0xa9c0a9 "github.com/bitwormhole/gitlib/git/support/localfilesys/worktrees"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComLocalGitConfigFactory struct {
	instance *config0x71b4a2.LocalGitConfigFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComGitHeadFileFactory struct {
	instance *head0xb6393b.GitHeadFileFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComGitIndexFileFactory struct {
	instance *index0x3eb559.GitIndexFileFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComGitModulesDirFactory struct {
	instance *modules0xe8c0dc.GitModulesDirFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComLocalGitObjectsFactory struct {
	instance *objects0x11508a.LocalGitObjectsFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComLocalGitRefsFactory struct {
	instance *refs0x4e5472.LocalGitRefsFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComLocalRepoDriver struct {
	instance *localfilesys0x6be3ff.LocalRepoDriver
	 markup0x23084a.Component `class:"git-repository-driver"`
	Layout files0x00707a.Layout `inject:"#git-local-repository-layout"`
	Locator files0x00707a.RepositoryLocator `inject:"#git-local-repository-locator"`
	Factory localfilesys0x6be3ff.RepoFactory `inject:"#git-local-repository-factory"`
}


type pComLocalRepoFactory struct {
	instance *localfilesys0x6be3ff.LocalRepoFactory
	 markup0x23084a.Component `id:"git-local-repository-factory"`
	Elements []localfilesys0x6be3ff.ElementFactory `inject:".git-local-element-factory"`
}


type pComLocalRepoLayout struct {
	instance *localfilesys0x6be3ff.LocalRepoLayout
	 markup0x23084a.Component `id:"git-local-repository-layout"`
}


type pComLocalRepoLocator struct {
	instance *localfilesys0x6be3ff.LocalRepoLocator
	 markup0x23084a.Component `id:"git-local-repository-locator"`
}


type pComCoreDirFactory struct {
	instance *views0x7d154b.CoreDirFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComShellDirFactory struct {
	instance *views0x7d154b.ShellDirFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComRepoViewFactory struct {
	instance *views0x7d154b.RepoViewFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComWorkingDirFactory struct {
	instance *views0x7d154b.WorkingDirFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComGitWorktreesDirFactory struct {
	instance *worktrees0xa9c0a9.GitWorktreesDirFactory
	 markup0x23084a.Component `class:"git-local-element-factory"`
}


type pComGitRepoManager struct {
	instance *support0x074feb.GitRepoManager
	 markup0x23084a.Component `id:"git-repository-manager"`
	Drivers []repository0x5aaf5a.Driver `inject:".git-repository-driver"`
}

