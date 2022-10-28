package support

import (
	"github.com/bitwormhole/gitlib/git/store"
)

// DefaultContextFactory ...
type DefaultContextFactory struct {
}

func (inst *DefaultContextFactory) _Impl() store.ContextFactory {
	return inst
}

func (inst *DefaultContextFactory) String() string {
	return "DefaultContextFactory"
}

// Create ...
func (inst *DefaultContextFactory) Create(cfg *store.ContextConfiguration) *store.Context {
	c2 := &store.Context{}
	confs := cfg.ContextConfigurers
	for _, conf := range confs {
		err := conf.Configure(c2)
		if err != nil {
			panic(err)
		}
	}
	c2.CoreConfigurers = cfg.CoreConfigurers
	return c2
}
