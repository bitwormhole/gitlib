package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// Pull 表示一条git命令
type Pull struct {
	repositories.Meta

	Service PullService

	// Path string
}

// Run ...
func (inst *Pull) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Pull) GetMeta() *repositories.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// PullService 。。。
type PullService interface {
	repositories.Service
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
