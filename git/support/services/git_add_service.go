package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitAddService ...
type GitAddService struct {
}

func (inst *GitAddService) _Impl() (store.ServiceRegistry, instructions.AddService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitAddService) ListRegistrations() []*store.ServiceRegistration {
	name := inst.Name()
	reg := &store.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*store.ServiceRegistration{reg}
}

// Name ...
func (inst *GitAddService) Name() string {
	return instructions.GitAdd
}

// Run ...
func (inst *GitAddService) Run(task *instructions.Add) error {
	return errors.New("no impl")
}
