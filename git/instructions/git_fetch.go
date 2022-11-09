package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Fetch 表示一条git命令
type Fetch struct {
	store.Meta

	Service FetchService

	// Path string
}

// Run ...
func (inst *Fetch) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Fetch) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// FetchService 。。。
type FetchService interface {
	store.Service
	Run(task *Fetch) error
}

////////////////////////////////////////////////////////////////////////////////

// NewFetch ...
func NewFetch(c context.Context) *Fetch {
	cmd := &Fetch{}
	cmd.Context = c
	cmd.Name = GitFetch
	cmd.Service = findService(&cmd.Meta).(FetchService)
	return cmd
}
