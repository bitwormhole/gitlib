package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Commit 表示一条git命令
type Commit struct {
	services.Command

	Service CommitService

	// Path string
}

// Run ...
func (inst *Commit) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Commit) GetCommand() *services.Command {
	return &inst.Command
}

////////////////////////////////////////////////////////////////////////////////

// CommitService 。。。
type CommitService interface {
	Run(task *Commit) error
}

////////////////////////////////////////////////////////////////////////////////

// NewCommit ...
func NewCommit(c context.Context) *Commit {
	cmd := &Commit{}
	cmd.Context = c
	cmd.Name = services.GitCommit
	cmd.Service = findServiceForCommand(&cmd.Command).(CommitService)
	return cmd
}
