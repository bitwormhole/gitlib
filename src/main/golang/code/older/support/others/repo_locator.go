package others

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
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
	wd := b.WD
	l := lookup{}
	result := l.findFileOrDir(".git", wd)
	b.DotGit = result
	if result == nil {
		return errors.New("no repository found in path " + wd.GetPath())
	}
	return nil
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
			point := dotgit.GetParent().GetChild(href)
			return inst.loadSubmodule(b, point)
		}
		// worktree
		point := dotgit.GetFS().NewPath(href)
		return inst.loadWorktree(b, point)
	}

	return errors.New("bad reposiotry, path=" + dotgit.GetPath())
}

func (inst *RepositoryLocatorImpl) loadSubmodule(b *store.LayoutBuilder, point afs.Path) error {
	if !point.IsDirectory() {
		return errors.New("bad submodule, path=" + point.GetPath())
	}
	b.SubmodulePoint = point
	return inst.handleSubmodulePre(b)
}

func (inst *RepositoryLocatorImpl) loadWorktree(b *store.LayoutBuilder, point afs.Path) error {
	if !point.IsDirectory() {
		return errors.New("bad worktree, path=" + point.GetPath())
	}
	b.WorktreePoint = point
	return inst.handleWorktreePre(b)
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
		return inst.handleSubmodulePost(b)
	} else if worktree != nil {
		return inst.handleWorktreePost(b)
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

func (inst *RepositoryLocatorImpl) handleWorktreePre(b *store.LayoutBuilder) error {

	l := lookup{}
	point := b.WorktreePoint
	config := l.findFile("config", point)

	if config == nil {
		return fmt.Errorf("cannot find git repository with worktree point: " + point.GetPath())
	}

	b.Repository = config.GetParent()
	return nil
}

func (inst *RepositoryLocatorImpl) handleWorktreePost(b *store.LayoutBuilder) error {

	l := lookup{}
	point := b.WorktreePoint
	head := l.findFile("HEAD", point)
	index := l.findFile("index", point)

	if head == nil || index == nil {
		return fmt.Errorf("bad git worktree at " + point.GetPath())
	}

	// re-write:
	b.HEAD = head
	b.Index = index
	return nil
}

func (inst *RepositoryLocatorImpl) handleSubmodulePre(b *store.LayoutBuilder) error {

	l := lookup{}
	point := b.SubmodulePoint
	config := l.findFile("config", point)

	if config == nil {
		return fmt.Errorf("cannot find git repository with submodule point: " + point.GetPath())
	}

	b.Repository = config.GetParent()
	return nil
}

func (inst *RepositoryLocatorImpl) handleSubmodulePost(b *store.LayoutBuilder) error {

	// re-write:
	// b.HEAD = nil
	// b.Index = nil

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type lookup struct {
}

func (inst *lookup) findFile(name string, from afs.Path) afs.Path {
	return inst.findWithWant(name, from, true, false)
}

func (inst *lookup) findDir(name string, from afs.Path) afs.Path {
	return inst.findWithWant(name, from, false, true)
}

func (inst *lookup) findFileOrDir(name string, from afs.Path) afs.Path {
	return inst.findWithWant(name, from, true, true)
}

func (inst *lookup) findWithWant(name string, from afs.Path, wantFile bool, wantDir bool) afs.Path {
	if name == "" {
		return nil
	}
	for p := from; p != nil; p = p.GetParent() {
		tar := p.GetChild(name)
		if wantFile && wantDir {
			if tar.Exists() {
				return tar
			}
		} else if wantFile {
			if tar.IsFile() {
				return tar
			}
		} else if wantDir {
			if tar.IsDirectory() {
				return tar
			}
		} else {
			return nil
		}
	}
	return nil
}
