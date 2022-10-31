package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Pull 表示一条git命令
type Pull struct {
	instructions.Meta

	Service PullService

	// Path string
}

// Run ...
func (inst *Pull) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Pull) GetMeta() *instructions.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// PullService 。。。
type PullService interface {
	Run(task *Pull) error
}

////////////////////////////////////////////////////////////////////////////////

// NewPull ...
func NewPull(c context.Context) *Pull {
	cmd := &Pull{}
	cmd.Context = c
	cmd.Name = instructions.GitPull
	cmd.Service = findService(&cmd.Meta).(PullService)
	return cmd
}
