package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// Clone 表示一条git命令
type Clone struct {
	repositories.Meta

	Service CloneService

	// Path string
}

// Run ...
func (inst *Clone) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Clone) GetMeta() *repositories.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// CloneService 。。。
type CloneService interface {
	repositories.Service
	Run(task *Clone) error
}

////////////////////////////////////////////////////////////////////////////////

// NewClone ...
func NewClone(c context.Context) *Clone {
	cmd := &Clone{}
	cmd.Context = c
	cmd.Name = GitClone
	cmd.Service = findService(&cmd.Meta).(CloneService)
	return cmd
}
