package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Add 表示一条git命令
type Add struct {
	services.Command

	Service AddService

	// Path string
}

// Run ...
func (inst *Add) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Add) GetCommand() *services.Command {
	return &inst.Command
}

////////////////////////////////////////////////////////////////////////////////

// AddService 。。。
type AddService interface {
	Run(task *Add) error
}

////////////////////////////////////////////////////////////////////////////////

// NewAdd ...
func NewAdd(c context.Context) *Add {
	cmd := &Add{}
	cmd.Context = c
	cmd.Name = services.GitAdd
	cmd.Service = findServiceForCommand(&cmd.Command).(AddService)
	return cmd
}
