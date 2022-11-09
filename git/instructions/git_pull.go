package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Pull 表示一条git命令
type Pull struct {
	store.Meta

	Service PullService

	// Path string
}

// Run ...
func (inst *Pull) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Pull) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// PullService 。。。
type PullService interface {
	store.Service
	Run(task *Pull) error
}

////////////////////////////////////////////////////////////////////////////////

// NewPull ...
func NewPull(c context.Context) *Pull {
	cmd := &Pull{}
	cmd.Context = c
	cmd.Name = GitPull
	cmd.Service = findService(&cmd.Meta).(PullService)
	return cmd
}
