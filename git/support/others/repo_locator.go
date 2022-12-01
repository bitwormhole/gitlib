package others

import (
	"errors"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

// RepositoryLocatorImpl ...
type RepositoryLocatorImpl struct {
}

func (inst *RepositoryLocatorImpl) _Impl() store.RepositoryLocator {
	return inst
}

// Locate ...
func (inst *RepositoryLocatorImpl) Locate(wd afs.Path) (store.RepositoryLayout, error) {

	type stepFn func(b *store.LayoutBuilder) error
	steps := make([]stepFn, 0)

	/////////////////////////////////////////

	steps = append(steps, inst.findDotGit)
	steps = append(steps, inst.loadDotGit)
	steps = append(steps, inst.initOtherDirs)

	/////////////////////////////////////////

	builder := &store.LayoutBuilder{}
	builder.WD = wd

	for _, step := range steps {
		err := step(builder)
		if err != nil {
			return nil, err
		}
	}

	layout := builder.Create()
	return layout, nil
}

func (inst *RepositoryLocatorImpl) findDotGit(b *store.LayoutBuilder) error {
	const name = ".git"
	wd := b.WD
	for p := wd; p != nil; p = p.GetParent() {
		dotgit := p.GetChild(name)
		if dotgit.Exists() {
			b.DotGit = dotgit
			return nil
		}
	}
	return errors.New("no repository found in path " + wd.GetPath())
}

func (inst *RepositoryLocatorImpl) loadDotGit(b *store.LayoutBuilder) error {

	dotgit := b.DotGit
	b.Workspace = dotgit.GetParent()

	if dotgit.IsDirectory() {
		b.Repository = dotgit
		return nil
	} else if !dotgit.IsFile() {
		return errors.New("the path is not a repository, path=" + dotgit.GetPath())
	}

	// load as file
	text, err := dotgit.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	// submodule: "gitdir: ../../../.git/modules/into/mods/foo-bar"
	// worktree: "gitdir: E:/home/dev/test/example_git_repos/repo4_worktree/.git/worktrees/tree1"

	const prefix = "gitdir:"
	if strings.HasPrefix(text, prefix) {
		href := strings.TrimSpace(text[len(prefix):])
		if strings.HasPrefix(href, ".") {
			// submodule
			point := dotgit.GetChild(href)
			b.SubmodulePoint = point
			return inst.loadSubmodule(b, point)
		}
		// worktree
		point := dotgit.GetFS().NewPath(href)
		b.WorktreePoint = point
		return inst.loadWorktree(b, point)
	}

	return errors.New("bad reposiotry, path=" + dotgit.GetPath())
}

func (inst *RepositoryLocatorImpl) loadSubmodule(b *store.LayoutBuilder, point afs.Path) error {

	if !point.IsDirectory() {
		return errors.New("bad submodule, path=" + point.GetPath())
	}

	return errors.New("no impl: loadSubmodule")
}

func (inst *RepositoryLocatorImpl) loadWorktree(b *store.LayoutBuilder, point afs.Path) error {

	if !point.IsDirectory() {
		return errors.New("bad worktree, path=" + point.GetPath())
	}

	return errors.New("no impl: loadWorktree")
}

func (inst *RepositoryLocatorImpl) initOtherDirs(b *store.LayoutBuilder) error {

	repo := b.Repository
	b.Objects = repo.GetChild("objects")
	b.Refs = repo.GetChild("refs")
	b.Index = repo.GetChild("index")
	b.HEAD = repo.GetChild("HEAD")
	b.Config = repo.GetChild("config")

	worktree := b.WorktreePoint
	submodule := b.SubmodulePoint

	if submodule != nil {
		// todo ...
	}

	if worktree != nil {
		// todo ...
	}

	return nil
}

func (inst *RepositoryLocatorImpl) findInBareRepo(b *store.LayoutBuilder) error {
	return errors.New("no impl: findInBareRepo")
}

func (inst *RepositoryLocatorImpl) findConfigFile(p afs.Path) (afs.Path, error) {
	const name = "config"
	p0 := p
	for ; p != nil; p = p.GetParent() {
		config := p.GetChild(name)
		if inst.isConfigFile(config) {
			return config, nil
		}
	}
	return nil, errors.New("no config found in path " + p0.GetPath())
}

func (inst *RepositoryLocatorImpl) isConfigFile(file afs.Path) bool {

	parent := file.GetParent()
	count1 := 0
	count2 := 0

	count1 += inst.checkDir(parent, "objects")
	count1 += inst.checkDir(parent, "refs")
	count1 += inst.checkFile(parent, "config")
	count1 += inst.checkFile(parent, "HEAD")

	count2 += inst.checkDir(parent, "info")
	count2 += inst.checkDir(parent, "logs")
	count2 += inst.checkDir(parent, "hooks")
	count2 += inst.checkFile(parent, "index")
	count2 += inst.checkFile(parent, "description")

	return (count1 >= 4) && (count2 >= 2)
}

// 如果成立，返回1，否则返回0
func (inst *RepositoryLocatorImpl) checkFile(parent afs.Path, child string) int {
	node := parent.GetChild(child)
	if node.IsFile() {
		return 1
	}
	return 0
}

// 如果成立，返回1，否则返回0
func (inst *RepositoryLocatorImpl) checkDir(parent afs.Path, child string) int {
	node := parent.GetChild(child)
	if node.IsDirectory() {
		return 1
	}
	return 0
}
