package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Example 表示一条git命令
type Example struct {
	services.Command

	Service ExampleService

	// Path string
}

// Run ...
func (inst *Example) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Example) GetCommand() *services.Command {
	return &inst.Command
}

////////////////////////////////////////////////////////////////////////////////

// ExampleService 。。。
type ExampleService interface {
	Run(task *Example) error
}

////////////////////////////////////////////////////////////////////////////////

// NewExample ...
func NewExample(c context.Context) *Example {
	cmd := &Example{}
	cmd.Context = c
	cmd.Name = services.GitExample
	cmd.Service = findServiceForCommand(&cmd.Command).(ExampleService)
	return cmd
}
