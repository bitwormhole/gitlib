package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Status 表示一条git命令
type Status struct {
	services.Command

	Service StatusService

	// Path string
}

// Run ...
func (inst *Status) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Status) GetCommand() *services.Command {
	return &inst.Command
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
	cmd.Name = services.GitStatus
	cmd.Service = findServiceForCommand(&cmd.Command).(StatusService)
	return cmd
}
