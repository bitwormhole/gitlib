package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitCommitService ...
type GitCommitService struct {
}

func (inst *GitCommitService) _Impl() (store.ServiceRegistry, instructions.CommitService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitCommitService) ListRegistrations() []*store.ServiceRegistration {
	name := inst.Name()
	reg := &store.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*store.ServiceRegistration{reg}
}

// Name ...
func (inst *GitCommitService) Name() string {
	return instructions.GitCommit
}

// Run ...
func (inst *GitCommitService) Run(task *instructions.Commit) error {
	return errors.New("no impl")
}
