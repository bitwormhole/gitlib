package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/gitlib/git/repository/objects"
	"github.com/bitwormhole/gitlib/git/repository/refs"
	"github.com/bitwormhole/starter/lang"
)

type Repo struct {
	Factory  RepoFactory
	Elements []lang.Object

	// detail
	Objects objects.Directory
	Refs    refs.Directory
	Config  repository.Config

	Head  repository.HEAD
	Index repository.Index

	// facade
	View    repository.View
	Shell   repository.Shell
	Core    repository.Core
	Working repository.WorkingDirectory
}

func (inst *Repo) AddElement(o lang.Object) {
	inst.Elements = append(inst.Elements, o)
}
