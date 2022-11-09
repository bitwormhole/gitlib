package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitExampleService ...
type GitExampleService struct {
}

func (inst *GitExampleService) _Impl() (store.ServiceRegistry, instructions.ExampleService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitExampleService) ListRegistrations() []*store.ServiceRegistration {
	name := inst.Name()
	reg := &store.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*store.ServiceRegistration{reg}
}

// Name ...
func (inst *GitExampleService) Name() string {
	return instructions.GitExample
}

// Run ...
func (inst *GitExampleService) Run(task *instructions.Example) error {
	return errors.New("no impl")
}
