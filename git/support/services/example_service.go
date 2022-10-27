package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/services"
)

// GitExampleService ...
type GitExampleService struct {
}

func (inst *GitExampleService) _Impl() (services.ServiceRegistry, git.ExampleService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitExampleService) ListRegistrations() []*services.ServiceRegistration {
	name := inst.Name()
	reg := &services.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*services.ServiceRegistration{reg}
}

// Name ...
func (inst *GitExampleService) Name() string {
	return services.GitExample
}

// Run ...
func (inst *GitExampleService) Run(task *git.Example) error {
	return errors.New("no impl")
}
