package lib

import (
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// ConfigContextBase ...
type ConfigContextBase struct {
	markup.Component `class:"git-context-configurer"`
	support.BaseContextConfigurer
}

func (inst *ConfigContextBase) _Impl() store.ContextConfigurer {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// ConfigContextWithInstructions ...
type ConfigContextWithInstructions struct {
	markup.Component `class:"git-context-configurer"`

	Instructions []store.ServiceRegistry `inject:".git-instruction-registry"`
}

func (inst *ConfigContextWithInstructions) _Impl() store.ContextConfigurer {
	return inst
}

// Configure ...
func (inst *ConfigContextWithInstructions) Configure(c *store.Context) error {
	c.Services = inst.Instructions
	return nil
}

////////////////////////////////////////////////////////////////////////////////
