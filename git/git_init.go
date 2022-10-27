package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Init 表示一条git命令
type Init struct {
	services.Command

	Service InitService

	Directory string
	Bare      bool
}

// Run ...
func (inst *Init) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Init) GetCommand() *services.Command {
	return &inst.Command
}

////////////////////////////////////////////////////////////////////////////////

// InitService 。。。
type InitService interface {
	Run(task *Init) error
}

////////////////////////////////////////////////////////////////////////////////

// NewInit ...
func NewInit(c context.Context) *Init {
	cmd := &Init{}
	cmd.Context = c
	cmd.Name = services.GitInit
	cmd.Service = findServiceForCommand(&cmd.Command).(InitService)
	return cmd
}
