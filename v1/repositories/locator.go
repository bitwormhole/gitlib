package repositories

import (
	"fmt"

	"github.com/bitwormhole/gitlib/v1/formats"
	"github.com/starter-go/v1/afs"
)

type Locator struct {
}

func (inst *Locator) Locate(wd afs.Directory) (Layout, error) {
	il := new(innerLocator)
	return il.locate(wd)
}

////////////////////////////////////////////////////////////////////////////////

type innerLocator struct{}

func (inst *innerLocator) locate(wd afs.Directory) (Layout, error) {

	path := wd
	ttl := path.CountParents()

	for ; (ttl > 0) && (path != nil); ttl-- {
		child := path.GetChild(".git")
		if child.IsDir() {
			dir := child.(afs.Directory)
			lb := new(LayoutBuilder)
			lb.DotGitDir = dir
			lb.WD = wd
			return inst.locateWithDotGitDir(dir, lb)

		} else if child.IsFile() {
			file := child.(afs.File)
			lb := new(LayoutBuilder)
			lb.DotGitFile = file
			lb.WD = wd
			return inst.locateWithDotGitFile(file, lb)

		} else {
			path = path.GetParent()
		}
	}

	layout, err := inst.tryLocateAsBareRepo(wd)
	if err == nil {
		return layout, nil
	}

	return nil, fmt.Errorf("cannot locate git repository at path [%s]", wd.GetPath())
}

func (inst *innerLocator) tryLocateAsBareRepo(path afs.Directory) (Layout, error) {

	// 这个方法主要是 通过 查找以下几个文件|文件夹,来达到定位的目的:
	// dir: objects
	// dir: refs
	// file: config
	// file: HEAD

	if path == nil {
		return nil, fmt.Errorf("innerLocator.tryLocateAsBareRepo() : param [path] is nil")
	}
	ptr := path
	ttl := ptr.CountParents()

	for ; (ttl >= 0) && (ptr != nil); ttl-- {

		dirRefs := ptr.GetChild("refs")
		dirObjects := ptr.GetChild("objects")
		fileConfig := ptr.GetChild("config")
		fileHead := ptr.GetChild("HEAD")

		b1 := dirRefs.IsDir()
		b2 := dirObjects.IsDir()
		b3 := fileConfig.IsFile()
		b4 := fileHead.IsFile()

		if b1 && b2 && b3 && b4 {
			repo := fileConfig.GetParent()
			lb := new(LayoutBuilder)
			lb.WD = path
			return inst.makeLayoutWithRepoDir(repo, lb)
		}

		ptr = ptr.GetParent()
	}

	return nil, fmt.Errorf("cannot find git-repo at path [%s]", path.String())
}

func (inst *innerLocator) locateWithDotGitFile(dotgit afs.File, lb *LayoutBuilder) (Layout, error) {

	err := inst.loadDotGitFile(dotgit, lb)
	if err != nil {
		return nil, err
	}

	repo := lb.Repository
	if repo == nil {
		return nil, fmt.Errorf("git repo dir is nil")
	}

	lb.CurrentWorktree = dotgit.GetParent()
	return inst.makeLayoutWithRepoDir(repo, lb)
}

func (inst *innerLocator) locateWithDotGitDir(dotgit afs.Directory, lb *LayoutBuilder) (Layout, error) {

	lb.CurrentWorktree = dotgit.GetParent()
	return inst.makeLayoutWithRepoDir(dotgit, lb)
}

func (inst *innerLocator) makeLayoutWithRepoDir(repo afs.Directory, lb *LayoutBuilder) (Layout, error) {

	fileConfig := repo.GetChild("config")
	fileDescription := repo.GetChild("description")
	fileHead := repo.GetChild("HEAD")
	fileIndex := repo.GetChild("index")
	filePackedRefs := repo.GetChild("packed-refs")

	dirObjects := repo.GetChild("objects")
	dirRefs := repo.GetChild("refs")
	dirWorktrees := repo.GetChild("worktrees")
	dirModules := repo.GetChild("modules")
	dirLogs := repo.GetChild("logs")
	dirHooks := repo.GetChild("hooks")
	dirInfo := repo.GetChild("info")
	dirBranches := repo.GetChild("branches")

	//files

	lb.Config = fileConfig.(afs.File)
	lb.Description = fileDescription.(afs.File)
	lb.HEAD = fileHead.(afs.File)
	lb.Index = fileIndex.(afs.File)
	lb.PackedRefs = filePackedRefs.(afs.File)

	//dirs

	lb.Objects = dirObjects.(afs.Directory)
	lb.Refs = dirRefs.(afs.Directory)
	lb.Worktrees = dirWorktrees.(afs.Directory)
	lb.Modules = dirModules.(afs.Directory)
	lb.Branches = dirBranches.(afs.Directory)
	lb.Hooks = dirHooks.(afs.Directory)
	lb.Info = dirInfo.(afs.Directory)
	lb.Logs = dirLogs.(afs.Directory)

	// complete

	lb.Repository = repo
	layout := lb.Build()
	err := inst.checkLayout(layout)
	if err != nil {
		return nil, err
	}
	return layout, nil
}

func (inst *innerLocator) checkLayout(layout Layout) error {

	const (
		required = true
	)
	var err error
	checker := new(innerLocatorNodeChecker)

	// required
	checker.reset(required)

	checker.addFile(layout.Config())

	checker.addDir(layout.Objects())
	checker.addDir(layout.Refs())

	err = checker.check()
	if err != nil {
		return err
	}

	// optional
	checker.reset(!required)

	checker.addFile(layout.Description())
	checker.addFile(layout.Index())
	checker.addFile(layout.HEAD())
	checker.addFile(layout.PackedRefs())

	checker.addDir(layout.Hooks())
	checker.addDir(layout.Branches())
	checker.addDir(layout.Info())
	checker.addDir(layout.Logs())

	return checker.check()
}

func (inst *innerLocator) loadDotGitFile(dotgit afs.File, lb *LayoutBuilder) error {

	dgfile := new(formats.DotGitFile)
	dgfile.File = dotgit
	err := dgfile.Read()
	if err != nil {
		return err
	}

	path := dgfile.Target
	node1 := dotgit.GetParent()
	node2 := node1.GetHref(path)
	endpoint := node2.(afs.Directory)

	bare, err := inst.tryLocateAsBareRepo(endpoint)
	if err != nil {
		return err
	}

	lb.Repository = bare.Repository()
	lb.DotGitEndpoint = endpoint
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type innerLocatorNodeChecker struct {
	required bool
	files    []afs.File
	dirs     []afs.Directory
}

func (inst *innerLocatorNodeChecker) reset(required bool) {
	inst.dirs = nil
	inst.files = nil
	inst.required = required
}

func (inst *innerLocatorNodeChecker) addDir(dir afs.Directory) {
	if dir == nil {
		return
	}
	inst.dirs = append(inst.dirs, dir)
}

func (inst *innerLocatorNodeChecker) addFile(file afs.File) {
	if file == nil {
		return
	}
	inst.files = append(inst.files, file)
}

func (inst *innerLocatorNodeChecker) check() error {

	for _, item := range inst.dirs {
		err := inst.onCheckDir(item)
		if err != nil {
			return err
		}
	}

	for _, item := range inst.files {
		err := inst.onCheckFile(item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *innerLocatorNodeChecker) onCheckFile(file afs.File) error {

	meta := file.GetMeta(nil)

	if meta.IsFile() {
		return nil
	} else {
		if (!inst.required) && (!meta.Exists()) {
			return nil
		}
	}

	path := file.GetPath()
	return fmt.Errorf("bad git repo: no file [%s]", path)
}

func (inst *innerLocatorNodeChecker) onCheckDir(dir afs.Directory) error {

	meta := dir.GetMeta(nil)

	if meta.IsDir() {
		return nil
	} else {
		if (!inst.required) && (!meta.Exists()) {
			return nil
		}
	}

	path := dir.GetPath()
	return fmt.Errorf("bad git repo: no dir [%s]", path)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
