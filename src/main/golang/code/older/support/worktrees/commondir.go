package worktrees

import "github.com/starter-go/afs"

// commondir ref '${git.repo}/worktrees/{worktree.name}/commondir' file
type commondir struct {
	path afs.Path
}
