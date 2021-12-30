package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/repository/modules"
	"github.com/bitwormhole/gitlib/git/repository/objects"
	"github.com/bitwormhole/gitlib/git/repository/refs"
	"github.com/bitwormhole/gitlib/git/repository/worktrees"
	"github.com/bitwormhole/starter/lang"
)

// Repo 仓库结构体
type Repo struct {
	Factory  RepoFactory
	Elements []lang.Object

	// detail

	Objects   objects.Directory
	Refs      refs.Directory
	Worktrees worktrees.Directory
	Modules   modules.Directory

	Config repository.Config
	Head   repository.HEAD
	Index  repository.Index

	// facade
	View    repository.View
	Shell   repository.Shell
	Core    repository.Core
	Working repository.WorkingDirectory
}

// AddElement ...
func (inst *Repo) AddElement(o lang.Object) {
	inst.Elements = append(inst.Elements, o)
}
