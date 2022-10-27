package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/services"
)

// GitCommitService ...
type GitCommitService struct {
}

func (inst *GitCommitService) _Impl() (services.ServiceRegistry, git.CommitService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitCommitService) ListRegistrations() []*services.ServiceRegistration {
	name := inst.Name()
	reg := &services.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*services.ServiceRegistration{reg}
}

// Name ...
func (inst *GitCommitService) Name() string {
	return services.GitCommit
}

// Run ...
func (inst *GitCommitService) Run(task *git.Commit) error {
	return errors.New("no impl")
}
