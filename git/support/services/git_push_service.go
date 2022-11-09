package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitPushService ...
type GitPushService struct {
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
	return errors.New("no impl")
}
