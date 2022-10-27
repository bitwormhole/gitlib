package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Push 表示一条git命令
type Push struct {
	services.Command

	Service PushService

	// Path string
}

// Run ...
func (inst *Push) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Push) GetCommand() *services.Command {
	return &inst.Command
}

////////////////////////////////////////////////////////////////////////////////

// PushService 。。。
type PushService interface {
	Run(task *Push) error
}

////////////////////////////////////////////////////////////////////////////////

// NewPush ...
func NewPush(c context.Context) *Push {
	cmd := &Push{}
	cmd.Context = c
	cmd.Name = services.GitPush
	cmd.Service = findServiceForCommand(&cmd.Command).(PushService)
	return cmd
}
