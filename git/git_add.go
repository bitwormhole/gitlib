package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Add 表示一条git命令
type Add struct {
	instructions.Meta

	Service AddService

	// Path string
}

// Run ...
func (inst *Add) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Add) GetMeta() *instructions.Meta {
	return &inst.Meta
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
	cmd.Name = instructions.GitAdd
	cmd.Service = findService(&cmd.Meta).(AddService)
	return cmd
}
