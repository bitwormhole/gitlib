package cfg

import (
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/gitlib/src/test/golang/element"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

func ExportConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}

type theGitRepoTester struct {
	markup.Component
	instance *element.GitRepoTester `initMethod:"Start"`
	Path     string                 `inject:"${test.repo.path}"`
	RM       repository.Manager     `inject:"#git-repository-manager"`
}

type theGitRepoLocatorTester struct {
	markup.Component
	instance   *element.RepoLocatorTester `initMethod:"Start"`
	AppContext application.Context        `inject:"context"`
	RM         repository.Manager         `inject:"#git-repository-manager"`
}

type theGitRepoLayoutTester struct {
	markup.Component
	instance   *element.RepoLayoutTester `initMethod:"Start"`
	AppContext application.Context       `inject:"context"`
	RM         repository.Manager        `inject:"#git-repository-manager"`
}
