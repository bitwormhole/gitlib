package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitCommitService ...
type GitCommitService struct {
}

func (inst *GitCommitService) _Impl() (instructions.ServiceRegistry, git.CommitService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitCommitService) ListRegistrations() []*instructions.ServiceRegistration {
	name := inst.Name()
	reg := &instructions.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*instructions.ServiceRegistration{reg}
}

// Name ...
func (inst *GitCommitService) Name() string {
	return instructions.GitCommit
}

// Run ...
func (inst *GitCommitService) Run(task *git.Commit) error {
	return errors.New("no impl")
}
