package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// Example 表示一条git命令
type Example struct {
	repositories.Meta

	Service ExampleService

	// Path string
}

// Run ...
func (inst *Example) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Example) GetMeta() *repositories.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// ExampleService 。。。
type ExampleService interface {
	repositories.Service
	Run(task *Example) error
}

////////////////////////////////////////////////////////////////////////////////

// NewExample ...
func NewExample(c context.Context) *Example {
	cmd := &Example{}
	cmd.Context = c
	cmd.Name = GitExample
	cmd.Service = findService(&cmd.Meta).(ExampleService)
	return cmd
}
