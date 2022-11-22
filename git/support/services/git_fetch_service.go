package services

import (
	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/network/clients"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitFetchService ...
type GitFetchService struct {
}

func (inst *GitFetchService) _Impl() (store.ServiceRegistry, instructions.FetchService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitFetchService) ListRegistrations() []*store.ServiceRegistration {
	name := inst.Name()
	reg := &store.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*store.ServiceRegistration{reg}
}

// Name ...
func (inst *GitFetchService) Name() string {
	return instructions.GitFetch
}

// Run ...
func (inst *GitFetchService) Run(task *instructions.Fetch) error {
	t2 := innerGitFetchTask{instruction: task}
	return t2.run()
}

////////////////////////////////////////////////////////////////////////////////

type innerGitFetchTask struct {
	instruction *instructions.Fetch
	targets     []*gitconfig.RemoteAndBranch
}

func (inst *innerGitFetchTask) run() error {
	cc := &clients.Context{}
	return inst.forLocal(cc)
}

func (inst *innerGitFetchTask) forLocal(cc *clients.Context) error {

	task := inst.instruction
	wd := task.WD
	ctx := task.Context
	cc = cc.Clone()

	// init
	err := cc.Init(ctx, wd)
	if err != nil {
		return err
	}

	// open session
	session, err := cc.OpenSession()
	if err != nil {
		return err
	}
	defer func() { session.Close() }()
	cc.Local.Session = session

	// load remote & branch
	err = inst.loadRemoteAndBranch(cc)
	if err != nil {
		return err
	}

	remotes := inst.listRemotes(cc)
	for _, remote := range remotes {
		cc.Remote.Raw = *remote
		err := inst.forRemote(cc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *innerGitFetchTask) forRemote(cc *clients.Context) error {

	cc = cc.Clone()

	conn, err := cc.OpenConnection(pktline.ServiceGitUpload)
	if err != nil {
		return err
	}
	defer func() { conn.Close() }()
	cc.Remote.Connection = conn

	return nil
}

func (inst *innerGitFetchTask) forBranch(cc *clients.Context) error {

	cc = cc.Clone()

	return nil
}

func (inst *innerGitFetchTask) loadRemoteAndBranch(cc *clients.Context) error {

	return nil
}

func (inst *innerGitFetchTask) listBranchesByRemote(cc *clients.Context) map[string]*gitconfig.Branch {
	src := inst.targets
	dst := make(map[string]*gitconfig.Branch)
	rName1 := cc.Remote.Raw.Name
	for _, rnb := range src {
		rName2 := rnb.Remote.Name
		if rName1 == rName2 {
			branch := rnb.Branch
			name := branch.Name
			dst[name] = branch
		}
	}
	return dst
}

func (inst *innerGitFetchTask) listRemotes(cc *clients.Context) map[string]*gitconfig.Remote {
	src := inst.targets
	dst := make(map[string]*gitconfig.Remote)
	for _, rnb := range src {
		remote := rnb.Remote
		name := remote.Name
		dst[name] = remote
	}
	return dst
}
