package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/network/clients"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// GitPushService ...
type GitPushService struct {
	markup.Component `class:"git-instruction-registry"`

	MainClient clients.MainClient `inject:"#git-main-client"`
}

func (inst *GitPushService) _Impl() (store.ServiceRegistry, instructions.PushService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitPushService) ListRegistrations() []*store.ServiceRegistration {
	name := inst.Name()
	reg := &store.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*store.ServiceRegistration{reg}
}

// Name ...
func (inst *GitPushService) Name() string {
	return instructions.GitPush
}

// Run ...
func (inst *GitPushService) Run(task *instructions.Push) error {

	c := &clients.Context{}
	actions := GetActions()
	wd := task.WD
	ctx := task.Context

	// init
	err := c.Init(ctx, wd)
	if err != nil {
		return err
	}

	// open session
	session, err := c.OpenSession()
	if err != nil {
		return err
	}
	defer func() { session.Close() }()
	c.Session = session

	err = actions.LoadLocalConfig(c).Load()
	if err != nil {
		return err
	}

	err = actions.LoadRemoteConfig(c).Load()
	if err != nil {
		return err
	}

	return errors.New("no impl")
}

////////////////////////////////////////////////////////////////////////////////

type innerGitPushTask struct {
	parent *GitPushService
}

func (inst *innerGitPushTask) forRemotes(cc *clients.Context) error {
	// for remotes
	remotes := cc.Remotes
	for _, remote := range remotes {
		cc.RawRemote = *remote
		cc.Intent.URL = remote.URL
		err := inst.forRemote(cc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *innerGitPushTask) forRemote(cc *clients.Context) error {

	cc = cc.Clone()

	// open connection
	conn, err := cc.OpenConnection(&pktline.ConnParams{})
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()
	cc.Connection = conn

	return inst.parent.MainClient.Execute(cc)
}

////////////////////////////////////////////////////////////////////////////////
