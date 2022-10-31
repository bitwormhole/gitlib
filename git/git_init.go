package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Init 表示一条git命令
type Init struct {
	instructions.Meta

	Service InitService

	Quiet             bool
	Bare              bool
	TemplateDirectory string
	SeparateGitDir    string
	ObjectFormat      string
	BranchName        string
	Shared            bool
	Permissions       string
	Directory         string
}

// Run ...
func (inst *Init) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Init) GetMeta() *instructions.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// InitService 。。。
type InitService interface {
	Run(task *Init) error
}

////////////////////////////////////////////////////////////////////////////////

// NewInit ...
func NewInit(c context.Context) *Init {
	cmd := &Init{}
	cmd.Context = c
	cmd.Name = instructions.GitInit
	cmd.Service = findService(&cmd.Meta).(InitService)
	return cmd
}
