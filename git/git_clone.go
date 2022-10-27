package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Clone 表示一条git命令
type Clone struct {
	services.Command

	Service CloneService

	// Path string
}

// Run ...
func (inst *Clone) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Clone) GetCommand() *services.Command {
	return &inst.Command
}

////////////////////////////////////////////////////////////////////////////////

// CloneService 。。。
type CloneService interface {
	Run(task *Clone) error
}

////////////////////////////////////////////////////////////////////////////////

// NewClone ...
func NewClone(c context.Context) *Clone {
	cmd := &Clone{}
	cmd.Context = c
	cmd.Name = services.GitClone
	cmd.Service = findServiceForCommand(&cmd.Command).(CloneService)
	return cmd
}
