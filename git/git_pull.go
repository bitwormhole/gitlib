package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Pull 表示一条git命令
type Pull struct {
	services.Command

	Service PullService

	// Path string
}

// Run ...
func (inst *Pull) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Pull) GetCommand() *services.Command {
	return &inst.Command
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
	cmd.Name = services.GitPull
	cmd.Service = findServiceForCommand(&cmd.Command).(PullService)
	return cmd
}
