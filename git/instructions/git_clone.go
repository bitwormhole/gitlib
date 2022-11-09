package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Clone 表示一条git命令
type Clone struct {
	store.Meta

	Service CloneService

	// Path string
}

// Run ...
func (inst *Clone) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Clone) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// CloneService 。。。
type CloneService interface {
	store.Service
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
