package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Commit 表示一条git命令
type Commit struct {
	instructions.Meta

	Service CommitService

	// Path string
}

// Run ...
func (inst *Commit) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Commit) GetMeta() *instructions.Meta {
	return &inst.Meta
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
	cmd.Name = instructions.GitCommit
	cmd.Service = findService(&cmd.Meta).(CommitService)
	return cmd
}
