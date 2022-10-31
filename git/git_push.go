package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Push 表示一条git命令
type Push struct {
	instructions.Meta

	Service PushService

	// Path string
}

// Run ...
func (inst *Push) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Push) GetMeta() *instructions.Meta {
	return &inst.Meta
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
	cmd.Name = instructions.GitPush
	cmd.Service = findService(&cmd.Meta).(PushService)
	return cmd
}
