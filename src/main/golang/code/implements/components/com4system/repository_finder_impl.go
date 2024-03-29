package com4system

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// RepositoryFinderRegistry ...
type RepositoryFinderRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *RepositoryFinderRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *RepositoryFinderRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *RepositoryFinderRegistry) create(ctx *repositories.SystemContext) (any, error) {
	com := &repositoryFinderImpl{}
	ctx.RepositoryFinder = com
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type repositoryFinderImpl struct {
	// context *repositories.SystemContext
}

func (inst *repositoryFinderImpl) _impl() repositories.Finder { return inst }

// Find ...
func (inst *repositoryFinderImpl) Find(pwd afs.Path) ([]afs.Path, error) {
	limit := 32
	r := &repositoryFinderResults{}
	err := r.scan(pwd, limit)
	r.handleError(err)
	return r.end()
}

////////////////////////////////////////////////////////////////////////////////

type repositoryFinderResults struct {
	items []afs.Path
	err   error
}

func (inst *repositoryFinderResults) handleError(err error) {
	if err == nil {
		return
	}
	inst.err = err
}

func (inst *repositoryFinderResults) scan(p afs.Path, limit int) error {

	if limit < 0 {
		return fmt.Errorf("the path is too deep, path = %v", p.GetPath())
	}

	dgit := p.GetChild(".git")
	if dgit.IsDirectory() || dgit.IsFile() {
		inst.items = append(inst.items, dgit)
		return nil
	}

	if p.IsDirectory() {
		dirName := p.GetName()
		if strings.HasSuffix(dirName, ".git") {
			// maybe a bare repository
			if inst.isBareRepo(p) {
				inst.items = append(inst.items, p)
				return nil
			}
		}
		names := p.ListNames()
		sort.Strings(names)
		for _, name := range names {
			child := p.GetChild(name)
			err := inst.scan(child, limit-1)
			inst.handleError(err)
		}
	}

	return nil
}

func (inst *repositoryFinderResults) isBareRepo(p afs.Path) bool {
	return false // todo: no impl
}

func (inst *repositoryFinderResults) end() ([]afs.Path, error) {
	dst := inst.items
	if dst == nil {
		dst = []afs.Path{}
	}
	return dst, inst.err
}
