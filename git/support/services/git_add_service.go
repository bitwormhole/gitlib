package services

import (
	"errors"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitAddService ...
type GitAddService struct {
}

func (inst *GitAddService) _Impl() (instructions.ServiceRegistry, git.AddService) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitAddService) ListRegistrations() []*instructions.ServiceRegistration {
	name := inst.Name()
	reg := &instructions.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*instructions.ServiceRegistration{reg}
}

// Name ...
func (inst *GitAddService) Name() string {
	return instructions.GitAdd
}

// Run ...
func (inst *GitAddService) Run(task *git.Add) error {
	return errors.New("no impl")
}
