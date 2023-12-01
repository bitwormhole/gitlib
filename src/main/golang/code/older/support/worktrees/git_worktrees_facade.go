package worktrees

import (
	"sort"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
)

// Facade 实现 store.Worktrees 的 Facade
type Facade struct {
	Core         *store.Core
	worktreesDir afs.Path
}

func (inst *Facade) _Impl() store.Worktrees {
	return inst
}

func (inst *Facade) getWorktreesDir() afs.Path {
	wtdir := inst.worktreesDir
	if wtdir == nil {
		parent := inst.Core.Layout.Repository()
		wtdir = parent.GetChild("worktrees")
		inst.worktreesDir = wtdir
	}
	return wtdir
}

// Get ...
func (inst *Facade) Get(name string) store.Worktree {

	theWorktrees := inst.getWorktreesDir()
	theWorktree := theWorktrees.GetChild(name)
	theCommondirFile := theWorktree.GetChild("commondir")
	theGitdirFile := theWorktree.GetChild("gitdir")

	o := &worktree{
		name:   name,
		core:   inst.Core,
		dotgit: nil,
	}
	o.commondir.path = theCommondirFile
	o.gitdir.path = theGitdirFile

	tar, err := o.gitdir.ResolveTarget()
	if err == nil {
		o.dotgit = tar
	}

	return o
}

// List ...
func (inst *Facade) List() []store.Worktree {

	dst := make([]store.Worktree, 0)
	theWorktrees := inst.getWorktreesDir()
	if !theWorktrees.IsDirectory() {
		return dst
	}

	namelist := theWorktrees.ListNames()
	sort.Strings(namelist)
	for _, name := range namelist {
		gitdirfile := theWorktrees.GetChild(name + "/gitdir")
		if gitdirfile.IsFile() {
			wt := inst.Get(name)
			dst = append(dst, wt)
		}
	}

	return dst
}
