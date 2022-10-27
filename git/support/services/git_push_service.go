package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/services"
)

// GitPushService ...
type GitPushService struct {
}

func (inst *GitPushService) _Impl() (services.ServiceRegistry, git.PushService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitPushService) ListRegistrations() []*services.ServiceRegistration {
	name := inst.Name()
	reg := &services.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*services.ServiceRegistration{reg}
}

// Name ...
func (inst *GitPushService) Name() string {
	return services.GitPush
}

// Run ...
func (inst *GitPushService) Run(task *git.Push) error {
	return errors.New("no impl")
}
