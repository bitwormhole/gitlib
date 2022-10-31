package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Fetch 表示一条git命令
type Fetch struct {
	instructions.Meta

	Service FetchService

	// Path string
}

// Run ...
func (inst *Fetch) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Fetch) GetMeta() *instructions.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// FetchService 。。。
type FetchService interface {
	Run(task *Fetch) error
}

////////////////////////////////////////////////////////////////////////////////

// NewFetch ...
func NewFetch(c context.Context) *Fetch {
	cmd := &Fetch{}
	cmd.Context = c
	cmd.Name = instructions.GitFetch
	cmd.Service = findService(&cmd.Meta).(FetchService)
	return cmd
}
