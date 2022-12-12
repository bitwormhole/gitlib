package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Push 表示一条git命令
type Push struct {
	store.Meta

	Service PushService

	All                   bool
	Mirror                bool
	Tags                  bool
	FollowTags            bool
	Atomic                bool
	DryRun                bool
	ReceivePack           string
	Repo                  string
	Force                 bool
	Delete                bool
	Prune                 bool
	Verbose               bool
	SetUpstream           bool
	PushOption            string
	Signed                bool
	SignedValue           string
	NoSigned              bool
	ForceWithLease        string
	ForceIfIncludes       bool
	NoVerify              bool
	RepositoryRefspecList []string
}

// Run ...
func (inst *Push) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Push) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// PushService 。。。
type PushService interface {
	store.Service
	Run(task *Push) error
}

////////////////////////////////////////////////////////////////////////////////

// NewPush ...
func NewPush(c context.Context) *Push {
	cmd := &Push{}
	cmd.Context = c
	cmd.Name = GitPush
	cmd.Service = findService(&cmd.Meta).(PushService)
	return cmd
}
