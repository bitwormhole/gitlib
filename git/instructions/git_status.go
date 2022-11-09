package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Status 表示一条git命令
type Status struct {
	store.Meta

	Service StatusService

	// Path string
}

// Run ...
func (inst *Status) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Status) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// StatusService 。。。
type StatusService interface {
	store.Service
	Run(task *Status) error
}

////////////////////////////////////////////////////////////////////////////////

// NewStatus ...
func NewStatus(c context.Context) *Status {
	cmd := &Status{}
	cmd.Context = c
	cmd.Name = GitStatus
	cmd.Service = findService(&cmd.Meta).(StatusService)
	return cmd
}
