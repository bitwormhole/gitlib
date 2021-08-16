package repo

import (
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/io/fs"
)

// Layout 仓库的布局
type Layout struct {
	core fs.Path

	objects fs.Path
	refs    fs.Path
	logs    fs.Path
	info    fs.Path
	hooks   fs.Path

	config      fs.Path
	description fs.Path
	HEAD        fs.Path
	index       fs.Path
	OrigHead    fs.Path // the file named 'ORIG_HEAD'

	gitdir    fs.Path // the file named 'gitdir'
	commondir fs.Path // the file named 'commondir'

	worktrees fs.Path
	modules   fs.Path

	workspace fs.Path
}

func (inst *Layout) initWithCoreDir(coreDir fs.Path) error {

	core := coreDir
	inst.core = core

	inst.description = core.GetChild("description")
	inst.config = core.GetChild("config")
	inst.index = core.GetChild("index")
	inst.HEAD = core.GetChild("HEAD")

	inst.objects = core.GetChild("objects")
	inst.refs = core.GetChild("refs")
	inst.logs = core.GetChild("logs")
	inst.hooks = core.GetChild("hooks")
	inst.info = core.GetChild("info")

	return nil
}

func (inst *Layout) initCore(location *repository.Location) error {
	return inst.initWithCoreDir(location.CoreDirectory)
}

func (inst *Layout) initView(location *repository.Location) error {

	if location.SubmoduleDirectory != nil {
		return inst.initSubmodule(location)

	} else if location.WorktreeDirectory != nil {
		return inst.initAsWorktree(location)

	} else if location.WorkingDirectory != nil {
		return inst.initNormal(location)
	}
	return inst.initBare(location)
}

func (inst *Layout) initAsWorktree(location *repository.Location) error {

	err := inst.initWithCoreDir(location.CoreDirectory)
	if err != nil {
		return err
	}

	wktree := location.WorktreeDirectory
	inst.logs = wktree.GetChild("logs")
	inst.commondir = wktree.GetChild("commondir")
	inst.gitdir = wktree.GetChild("gitdir")
	inst.HEAD = wktree.GetChild("HEAD")
	inst.index = wktree.GetChild("index")
	inst.OrigHead = wktree.GetChild("ORIG_HEAD")

	inst.workspace = location.WorkingDirectory
	inst.core = location.CoreDirectory
	return nil
}

func (inst *Layout) initSubmodule(location *repository.Location) error {
	err := inst.initWithCoreDir(location.SubmoduleDirectory)
	if err != nil {
		return err
	}
	inst.workspace = location.WorkingDirectory
	inst.core = location.CoreDirectory
	return nil
}

func (inst *Layout) initNormal(location *repository.Location) error {
	err := inst.initWithCoreDir(location.CoreDirectory)
	if err != nil {
		return err
	}
	inst.core = location.CoreDirectory
	inst.workspace = location.WorkingDirectory
	return nil
}

func (inst *Layout) initBare(location *repository.Location) error {
	err := inst.initWithCoreDir(location.CoreDirectory)
	if err != nil {
		return err
	}
	inst.core = location.CoreDirectory
	inst.workspace = nil
	return nil
}
