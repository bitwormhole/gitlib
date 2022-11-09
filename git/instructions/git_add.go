package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Add 表示一条git命令
type Add struct {
	store.Meta

	Service AddService

	// Path string
}

// Run ...
func (inst *Add) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Add) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// AddService 。。。
type AddService interface {
	store.Service
	Run(task *Add) error
}

////////////////////////////////////////////////////////////////////////////////

// NewAdd ...
func NewAdd(c context.Context) *Add {
	cmd := &Add{}
	cmd.Context = c
	cmd.Name = GitAdd
	cmd.Service = findService(&cmd.Meta).(AddService)
	return cmd
}
