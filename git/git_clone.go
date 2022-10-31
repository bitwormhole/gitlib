package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Clone 表示一条git命令
type Clone struct {
	instructions.Meta

	Service CloneService

	// Path string
}

// Run ...
func (inst *Clone) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Clone) GetMeta() *instructions.Meta {
	return &inst.Meta
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
	cmd.Name = instructions.GitClone
	cmd.Service = findService(&cmd.Meta).(CloneService)
	return cmd
}
