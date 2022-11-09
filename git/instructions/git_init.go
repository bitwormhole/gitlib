package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Init 表示一条git命令
type Init struct {
	store.Meta

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
func (inst *Init) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// InitService 。。。
type InitService interface {
	store.Service
	Run(task *Init) error
}

////////////////////////////////////////////////////////////////////////////////

// NewInit ...
func NewInit(c context.Context) *Init {
	cmd := &Init{}
	cmd.Context = c
	cmd.Name = GitInit
	cmd.Service = findService(&cmd.Meta).(InitService)
	return cmd
}
