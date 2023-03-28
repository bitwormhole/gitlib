package worktrees

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support/others"
)

type worktree struct {
	core      *store.Core
	name      string
	dotgit    afs.Path // the '.git' file
	gitdir    gitdir
	commondir commondir
}

func (inst *worktree) _Impl() store.Worktree {
	return inst
}

func (inst *worktree) Name() string {
	return inst.name
}

func (inst *worktree) Workspace() store.Workspace {
	dg := inst.dotgit
	if dg == nil {
		return nil
	}
	if !dg.IsFile() {
		return nil
	}
	builder := others.GitWorkspaceBuilder{}
	builder.Core = inst.core
	builder.DotGit = dg
	return builder.Create()
}

func (inst *worktree) Exists() bool {
	filelist := make([]afs.Path, 0)
	filelist = append(filelist, inst.dotgit)
	filelist = append(filelist, inst.gitdir.path)
	filelist = append(filelist, inst.commondir.path)
	for _, f := range filelist {
		if f == nil {
			return false
		}
		if !f.IsFile() {
			return false
		}
	}
	return true
}
