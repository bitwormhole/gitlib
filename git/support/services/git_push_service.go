package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitPushService ...
type GitPushService struct {
}

func (inst *GitPushService) _Impl() (instructions.ServiceRegistry, git.PushService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitPushService) ListRegistrations() []*instructions.ServiceRegistration {
	name := inst.Name()
	reg := &instructions.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*instructions.ServiceRegistration{reg}
}

// Name ...
func (inst *GitPushService) Name() string {
	return instructions.GitPush
}

// Run ...
func (inst *GitPushService) Run(task *git.Push) error {
	return errors.New("no impl")
}
