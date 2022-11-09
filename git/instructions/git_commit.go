package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Commit 表示一条git命令
type Commit struct {
	store.Meta

	Service CommitService

	// Path string
}

// Run ...
func (inst *Commit) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Commit) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// CommitService 。。。
type CommitService interface {
	store.Service
	Run(task *Commit) error
}

////////////////////////////////////////////////////////////////////////////////

// NewCommit ...
func NewCommit(c context.Context) *Commit {
	cmd := &Commit{}
	cmd.Context = c
	cmd.Name = GitCommit
	cmd.Service = findService(&cmd.Meta).(CommitService)
	return cmd
}
