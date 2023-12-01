package lib

import (
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support/services"
	"github.com/bitwormhole/starter/markup"
)

// ConfigInstructions ...
type ConfigInstructions struct {
	markup.Component `class:"git-instruction-registry"`
}

func (inst *ConfigInstructions) _Impl() store.ServiceRegistry {
	return inst
}

// ListRegistrations ...
func (inst *ConfigInstructions) ListRegistrations() []*store.ServiceRegistration {

	list := []*store.ServiceRegistration{}

	list = inst.add(list, &services.GitAddService{})
	list = inst.add(list, &services.GitCommitService{})
	list = inst.add(list, &services.GitFetchService{})
	list = inst.add(list, &services.GitInitService{})
	list = inst.add(list, &services.GitPushService{})
	list = inst.add(list, &services.GitStatusService{})

	return list
}

func (inst *ConfigInstructions) add(dst []*store.ServiceRegistration, src store.ServiceRegistry) []*store.ServiceRegistration {
	some := src.ListRegistrations()
	return append(dst, some...)
}
