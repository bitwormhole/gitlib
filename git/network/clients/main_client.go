package clients

import (
	"fmt"

	"github.com/bitwormhole/starter/markup"
)

// MainClientImpl ...
type MainClientImpl struct {
	markup.Component `id:"git-main-client"`

	ClientRegistryList []ClientRegistry `inject:".git-client-registry"`

	// cached
	allClients []*ClientRegistration
}

func (inst *MainClientImpl) _Impl() MainClient {
	return inst
}

// Accept ...
func (inst *MainClientImpl) Accept(c *Context) bool {
	return true
}

// Execute ...
func (inst *MainClientImpl) Execute(c *Context) error {
	intent := &c.Intent
	all := inst.getAll()
	for _, r2 := range all {
		c2 := r2.Client
		if c2 == nil {
			continue
		}
		if c2.Accept(c) {
			return c2.Execute(c)
		}
	}
	url := intent.URL
	return fmt.Errorf("no git-client support this remote repository: %v", url)
}

func (inst *MainClientImpl) getAll() []*ClientRegistration {
	all := inst.allClients
	if all != nil {
		return all
	}
	all = make([]*ClientRegistration, 0)
	src := inst.ClientRegistryList
	for _, r1 := range src {
		r2 := r1.GetClientRegistration()
		if r2 != nil {
			all = append(all, r2)
		}
	}
	inst.allClients = all
	return all
}
