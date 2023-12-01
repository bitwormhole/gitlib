package submodules

import (
	"github.com/bitwormhole/gitlib/git/store"
)

// Facade 是实现 store.Submodules 的 Facade
type Facade struct {
	Core *store.Core

	// modules afs.Path // the '.git/modules' dir
	// dotgit  afs.Path
	// config  afs.Path

	cached []store.Submodule
}

func (inst *Facade) _Impl() store.Submodules {
	return inst
}

// Get ...
func (inst *Facade) Get(name string) store.Submodule {
	all := inst.List()
	for _, item := range all {
		if item.Name() == name {
			return item
		}
	}
	return nil
}

// List ...
func (inst *Facade) List() []store.Submodule {
	list := inst.cached
	if list == nil {
		list = inst.load()
		inst.cached = list
	}
	return list
}

func (inst *Facade) load() []store.Submodule {

	l := loader{}
	dst := make([]store.Submodule, 0)

	src, err := l.load(inst.Core)

	if err != nil {
		return dst
	}

	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}
