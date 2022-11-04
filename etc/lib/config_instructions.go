package lib

import (
	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/support/services"
	"github.com/bitwormhole/starter/markup"
)

// ConfigInstructions ...
type ConfigInstructions struct {
	markup.Component `class:"git-instruction-registry"`
}

func (inst *ConfigInstructions) _Impl() instructions.ServiceRegistry {
	return inst
}

// ListRegistrations ...
func (inst *ConfigInstructions) ListRegistrations() []*instructions.ServiceRegistration {

	list := []*instructions.ServiceRegistration{}

	list = inst.add(list, &services.GitAddService{})
	list = inst.add(list, &services.GitCommitService{})
	list = inst.add(list, &services.GitInitService{})
	list = inst.add(list, &services.GitPushService{})
	list = inst.add(list, &services.GitStatusService{})

	return list
}

func (inst *ConfigInstructions) add(dst []*instructions.ServiceRegistration, src instructions.ServiceRegistry) []*instructions.ServiceRegistration {
	some := src.ListRegistrations()
	return append(dst, some...)
}
