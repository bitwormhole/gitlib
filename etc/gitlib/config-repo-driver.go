package etcgitlib

import (
	"github.com/bitwormhole/gitlib/factory/repo"
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/markup"
)

type theRepoManager struct {
	markup.Component
	instance *repo.DefaultRepositoryManager `id:"git-repository-manager"`
	Drivers  []repository.Driver            `inject:".git-repository-driver"`
}

type theRepoDriverDefault struct {
	markup.Component
	instance  *repo.DefaultRepositoryDriver `id:"file-git-repository-driver" class:"git-repository-driver"`
	MyLocator repository.Locator            `inject:"#file-git-repository-locator"`
	MyFactory repository.Factory            `inject:"#file-git-repository-factory"`
}

type theFileRepoFactory struct {
	markup.Component
	instance *repo.FileRepositoryFactory `id:"file-git-repository-factory"`
	Pipeline []repo.ElementFactory       `inject:".default-git-repository-element"`
}

type theFileRepoLocator struct {
	markup.Component
	instance *repo.FileRepositoryLocator `id:"file-git-repository-locator"`
}

////////////////////////////////////////////////////////////////////////////////
// elements

type theElementConfig struct {
	markup.Component
	instance *repo.GitConfigElementFactory `class:"default-git-repository-element"`
}

type theElementIndex struct {
	markup.Component
	instance *repo.GitIndexElementFactory `class:"default-git-repository-element"`
}

type theElementHead struct {
	markup.Component
	instance *repo.GitHeadElementFactory `class:"default-git-repository-element"`
}

type theElementRefs struct {
	markup.Component
	instance *repo.GitRefsElementFactory `class:"default-git-repository-element"`
}

type theElementObjects struct {
	markup.Component
	instance *repo.GitObjectsElementFactory `class:"default-git-repository-element"`
}
