package support

import (
	"github.com/bitwormhole/gitlib/git/store"
)

// DefaultContextFactory ...
type DefaultContextFactory struct {
	// configurers []store.ContextConfigurer
}

func (inst *DefaultContextFactory) _Impl() store.ContextFactory {
	return inst
}

func (inst *DefaultContextFactory) String() string {
	return "DefaultContextFactory"
}

// func (inst *DefaultContextFactory) makeConfigurers() []store.ContextConfigurer {
// 	list := make([]store.ContextConfigurer, 0)

// 	list = append(list, &BaseContextConfigurer{})

// 	return list
// }

// func (inst *DefaultContextFactory) listConfigurers() []store.ContextConfigurer {
// 	list := inst.configurers
// 	if list == nil {
// 		list = inst.makeConfigurers()
// 		inst.configurers = list
// 	}
// 	return list
// }

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
	return c2
}
