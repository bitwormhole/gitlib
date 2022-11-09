package services

import (
	"errors"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git/instructions"

	"github.com/bitwormhole/gitlib/git/store"
)

// GitStatusService ...
type GitStatusService struct {
}

func (inst *GitStatusService) _Impl() (store.ServiceRegistry, instructions.StatusService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitStatusService) ListRegistrations() []*store.ServiceRegistration {
	name := inst.Name()
	reg := &store.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*store.ServiceRegistration{reg}
}

// Name ...
func (inst *GitStatusService) Name() string {
	return instructions.GitStatus
}

// Run ...
func (inst *GitStatusService) Run(task *instructions.Status) error {

	lib, err := store.GetLib(task.Context)
	if err != nil {
		return err
	}

	layout, err := lib.RepositoryLocator().Locate(task.WD)
	if err != nil {
		return err
	}

	repo, err := lib.RepositoryLoader().Load(layout)
	if err != nil {
		return err
	}

	session, err := repo.OpenSession()
	if err != nil {
		return err
	}
	defer func() {
		session.Close()
	}()

	head := repo.HEAD()
	if !head.Path().Exists() {
		return inst.log(task, "no HEAD")
	}
	refname, err := head.GetValue(session)
	if err != nil {
		return err
	}

	refs := repo.Refs()
	ref := refs.GetRef(refname)
	if !ref.Path().Exists() {
		return inst.log(task, "no ref: "+refname.String())
	}
	oid, err := ref.GetValue(session)
	if err != nil {
		return err
	}

	commit, err := session.LoadCommit(oid)
	if err != nil {
		return err
	}

	tree, err := session.LoadTree(commit.Tree)
	if err != nil {
		return err
	}

	count := len(tree.Items)
	vlog.Warn("todo: log tree info, items.count = ", count)

	return errors.New("no impl")
}

func (inst *GitStatusService) log(task *instructions.Status, msg string) error {
	ctx := task.Context
	b := cli.GetBinding(ctx)
	console := b.GetConsole()
	out := console.Out()
	out.WriteString(msg + "\n")
	return nil
}
