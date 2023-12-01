package services

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/network/clients"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// GitFetchService ...
type GitFetchService struct {
	markup.Component `class:"git-instruction-registry"`

	MainClient clients.MainClient `inject:"#git-main-client"`
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
	t2 := innerGitFetchTask{parent: inst, instruction: task}
	return t2.run()
}

////////////////////////////////////////////////////////////////////////////////

type innerGitFetchTask struct {
	parent      *GitFetchService
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
	cc.Session = session

	// load remote & branch
	actions := GetActions()
	err = actions.LoadRemoteConfig(cc).Load()
	if err != nil {
		return err
	}

	err = inst.prepareRnB(cc)
	if err != nil {
		return err
	}

	return inst.forRemotes(cc)
}

func (inst *innerGitFetchTask) forRemotes(cc *clients.Context) error {
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

func (inst *innerGitFetchTask) forRemote(cc *clients.Context) error {

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

func (inst *innerGitFetchTask) forBranch(cc *clients.Context) error {

	cc = cc.Clone()

	return nil
}

func (inst *innerGitFetchTask) prepareRnB(cc *clients.Context) error {

	// GitFetch ...
	// 参考 https://git-scm.com/docs/git-fetch

	fetch := inst.instruction
	if fetch.All {
		return inst.tryPrepareAllRnB(cc)
	} else if fetch.Multiple {
		return inst.tryPrepareMultipleRnB(cc)
	}

	err := inst.tryPrepareGroupRnB(cc)
	if err != nil {
		err = inst.tryPrepareRepositoryRnB(cc)
	}
	return err
}

func (inst *innerGitFetchTask) getHeadBranch(cc *clients.Context) (string, error) {
	session := cc.Session
	head, err := cc.Repository.HEAD().GetValue(session)
	if err != nil {
		return "", err
	}
	return head.String(), nil
}

func (inst *innerGitFetchTask) tryPrepareGroupRnB(cc *clients.Context) error {

	// git fetch [<options>] <group>

	return fmt.Errorf("no impl")
}

func (inst *innerGitFetchTask) tryPrepareRepositoryRnB(cc *clients.Context) error {

	// git fetch [<options>] [<repository> [<refspec>…​]]

	items := inst.instruction.Items
	dst := make([]*gitconfig.RemoteAndBranch, 0)

	for i, item := range items {
		if i == 0 {
			repo, err := inst.parseParamRepository(cc, item)
			if err != nil {
				return err
			}
			dst = append(dst, repo)
		} else {
			refspec, err := inst.parseParamRefspec(cc, item)
			if err != nil {
				return err
			}
			item0 := dst[0]
			name1 := item0.Remote.Name
			name2 := refspec.Branch.Remote
			if name1 != name2 {
				return fmt.Errorf("the remote [%v] do not contains a branch named [%v]", name1, name2)
			}
			dst = append(dst, refspec)
		}
	}

	size := len(dst)
	if size > 1 {
		dst = dst[1:]
	}

	cc.RnB = dst
	return nil
}

func (inst *innerGitFetchTask) tryPrepareMultipleRnB(cc *clients.Context) error {

	// git fetch --multiple [<options>] [(<repository> | <group>)…​]

	return fmt.Errorf("no impl")
}

func (inst *innerGitFetchTask) tryPrepareAllRnB(cc *clients.Context) error {

	// git fetch --all [<options>]

	return fmt.Errorf("no impl")
}

func (inst *innerGitFetchTask) parseParamRepository(cc *clients.Context, str string) (*gitconfig.RemoteAndBranch, error) {
	src := cc.Remotes
	rnb := &gitconfig.RemoteAndBranch{}
	rnb.Branch = &gitconfig.Branch{
		Name: "*",
	}
	r1 := src[str]
	if r1 != nil && r1.Name == str {
		rnb.Remote = r1
		return rnb, nil
	}
	for _, r := range src {
		if r.URL == str {
			rnb.Remote = r
			break
		}
	}
	if rnb.Remote == nil {
		return nil, fmt.Errorf("cannot find remote by name:%v", str)
	}
	return rnb, nil
}

func (inst *innerGitFetchTask) parseParamRefspec(cc *clients.Context, str string) (*gitconfig.RemoteAndBranch, error) {
	src := cc.Branches
	b1 := src[str]
	rnb := &gitconfig.RemoteAndBranch{}
	if b1 != nil && b1.Name == str {
		rnb.Branch = b1
		return rnb, nil
	}
	str1 := str
	if !strings.Contains(str1, "/") {
		str1 = "/" + str1
	}
	for _, b := range src {
		if strings.HasSuffix(b.Merge, str1) {
			rnb.Branch = b
			return rnb, nil
		}
	}
	return nil, fmt.Errorf("cannot find branch by name: %v", str)
}

func (inst *innerGitFetchTask) parseParamGroup(cc *clients.Context, str string) (*gitconfig.RemoteAndBranch, error) {
	return nil, fmt.Errorf("no impl")
}
