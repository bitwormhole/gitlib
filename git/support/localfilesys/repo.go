package localfilesys

import (
	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/starter/lang"
)

type Core struct {
}

type Shell struct {
}

type Working struct {
}

type Facade struct {
	View    repository.View
	Shell   repository.Shell
	Core    repository.Core
	Working repository.WorkingDirectory
}

type Repo struct {
	Factory  RepoFactory
	Elements []lang.Object

	Core    Core
	Shell   Shell
	Working Working
	Facade  Facade
}
