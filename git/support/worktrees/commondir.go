package worktrees

import "bitwormhole.com/starter/afs"

// commondir ref '${git.repo}/worktrees/{worktree.name}/commondir' file
type commondir struct {
	path afs.Path
}
