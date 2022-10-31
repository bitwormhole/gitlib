package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Status 表示一条git命令
type Status struct {
	instructions.Meta

	Service StatusService

	// Path string
}

// Run ...
func (inst *Status) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Status) GetMeta() *instructions.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// StatusService 。。。
type StatusService interface {
	Run(task *Status) error
}

////////////////////////////////////////////////////////////////////////////////

// NewStatus ...
func NewStatus(c context.Context) *Status {
	cmd := &Status{}
	cmd.Context = c
	cmd.Name = instructions.GitStatus
	cmd.Service = findService(&cmd.Meta).(StatusService)
	return cmd
}
