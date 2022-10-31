package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Example 表示一条git命令
type Example struct {
	instructions.Meta

	Service ExampleService

	// Path string
}

// Run ...
func (inst *Example) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Example) GetMeta() *instructions.Meta {
	return &inst.Meta
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
	cmd.Name = instructions.GitExample
	cmd.Service = findService(&cmd.Meta).(ExampleService)
	return cmd
}
