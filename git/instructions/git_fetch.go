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

	Depth                   bool // = "--depth"           // =<depth>
	Deepen                  bool // = "--deepen"          // =<depth>
	ShallowSince            bool // = "--shallow-since"   // =<date>
	ShallowExclude          bool // = "--shallow-exclude" // =<revision>
	NegotiationTip          bool // = "--negotiation-tip" // =<commit|glob>
	Refmap                  bool // = "--refmap" // =<refspec>
	RecurseSubmodule        bool // = "--recurse-submodules" // =[=yes|on-demand|no]
	Jobs                    bool // = "--jobs" // =<n>
	SubmodulePrefix         bool // = "--submodule-prefix"          // =<path>
	RecurseSubmoduleDefault bool // = "--recurse-submodule-default" // =[yes|on-demand]
	UploadPack              bool // = "--upload-pack" // <upload-pack>
	ServerOption            bool // = "--server- ion" // =< ion>

	All                 bool // = "--all"
	Append              bool // = "--append"
	Unshallow           bool // = "--unshallow"
	UpdateShallow       bool // = "--update-shallow"
	NegotiateOnly       bool // = "--negotiate-only"
	DryRun              bool // = "--dry-run"
	WriteFetchHead      bool // = "--write-fetch-head"
	NoWriteFetchHead    bool // = "--no-write-fetch-head"
	Force               bool // = "--force"
	Keep                bool // = "--keep"
	Multiple            bool // = "--multiple"
	NoAutoMaintenance   bool // = "--no-auto-maintenance"
	AutoMaintenance     bool // = "--auto-maintenance"
	NoAutoGC            bool // = "--no-auto-gc"
	AutoGC              bool // = "--auto-gc"
	NoWriteCommitGraph  bool // = "--no-write-commit-graph"
	WriteCommitGraph    bool // = "--write-commit-graph"
	Prefetch            bool // = "--prefetch"
	Prune               bool // = "--prune"
	PruneTags           bool // = "--prune-tags"
	NoTags              bool // = "--no-tags"
	Refetch             bool // = "--refetch"
	Tags                bool // = "--tags"
	NoRecurseSubmodules bool // = "--no-recurse-submodules"
	SetUpstream         bool // = "--set-upstream"
	UpdateHeadOK        bool // = "--update-head-ok"
	Quiet               bool // = "--quiet"
	Verbose             bool // = "--verbose"
	Progress            bool // = "--progress"
	ShowForcedUpdates   bool // = "--show-forced-updates"
	NoShowForcedUpdates bool // = "--no-show-forced-updates"
	Ipv4                bool // = "--ipv4"
	Ipv6                bool // = "--ipv6"
	Stdin               bool // = "--stdin"

	Items []string
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
