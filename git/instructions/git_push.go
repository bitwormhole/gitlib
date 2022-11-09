package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Push 表示一条git命令
type Push struct {
	store.Meta

	Service PushService

	// Path string
}

// Run ...
func (inst *Push) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Push) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// PushService 。。。
type PushService interface {
	store.Service
	Run(task *Push) error
}

////////////////////////////////////////////////////////////////////////////////

// NewPush ...
func NewPush(c context.Context) *Push {
	cmd := &Push{}
	cmd.Context = c
	cmd.Name = GitPush
	cmd.Service = findService(&cmd.Meta).(PushService)
	return cmd
}
