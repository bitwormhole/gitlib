package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Fetch 表示一条git命令
type Fetch struct {
	services.Command

	Service FetchService

	// Path string
}

// Run ...
func (inst *Fetch) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Fetch) GetCommand() *services.Command {
	return &inst.Command
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
	cmd.Name = services.GitFetch
	cmd.Service = findServiceForCommand(&cmd.Command).(FetchService)
	return cmd
}
