package services

import (
	"errors"

	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/instructions"

	"github.com/bitwormhole/gitlib/git/store"
)

// GitStatusService ...
type GitStatusService struct {
}

func (inst *GitStatusService) _Impl() (instructions.ServiceRegistry, git.StatusService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitStatusService) ListRegistrations() []*instructions.ServiceRegistration {
	name := inst.Name()
	reg := &instructions.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*instructions.ServiceRegistration{reg}
}

// Name ...
func (inst *GitStatusService) Name() string {
	return instructions.GitStatus
}

// Run ...
func (inst *GitStatusService) Run(task *git.Status) error {

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
	refname, err := head.GetValue(session)
	if err != nil {
		return err
	}

	refs := repo.Refs()
	ref := refs.GetRef(refname)
	oid, err := ref.GetValue(session)
	if err != nil {
		return err
	}

	commit, err := session.LoadCommit(oid)
	if err != nil {
		return err
	}

	vlog.Warn("todo: log commit info", commit)

	return errors.New("no impl")
}
