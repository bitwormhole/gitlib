package worktrees

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/utils"
	"github.com/starter-go/afs"
)

// gitdir ref '${git.repo}/worktrees/{worktree.name}/gitdir' file
type gitdir struct {
	path afs.Path
}

func (inst *gitdir) ResolveTarget() (afs.Path, error) {

	file := inst.path
	if file == nil {
		return nil, fmt.Errorf("no path of 'gitdir' file")
	}

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text)
	return utils.ComputeAbsolutePath(text, file)
}
