package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitExampleService ...
type GitExampleService struct {
}

func (inst *GitExampleService) _Impl() (instructions.ServiceRegistry, git.ExampleService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitExampleService) ListRegistrations() []*instructions.ServiceRegistration {
	name := inst.Name()
	reg := &instructions.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*instructions.ServiceRegistration{reg}
}

// Name ...
func (inst *GitExampleService) Name() string {
	return instructions.GitExample
}

// Run ...
func (inst *GitExampleService) Run(task *git.Example) error {
	return errors.New("no impl")
}
