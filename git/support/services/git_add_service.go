package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/services"
)

// GitAddService ...
type GitAddService struct {
}

func (inst *GitAddService) _Impl() (services.ServiceRegistry, git.AddService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitAddService) ListRegistrations() []*services.ServiceRegistration {
	name := inst.Name()
	reg := &services.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*services.ServiceRegistration{reg}
}

// Name ...
func (inst *GitAddService) Name() string {
	return services.GitAdd
}

// Run ...
func (inst *GitAddService) Run(task *git.Add) error {
	return errors.New("no impl")
}
