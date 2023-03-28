package submodules

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

// Facade 是实现 store.Submodules 的 Facade
type Facade struct {
	Core  *store.Core
	file1 afs.Path
	file2 afs.Path
}

func (inst *Facade) _Impl() store.Submodules {
	return inst
}

// Get ...
func (inst *Facade) Get(name string) store.Submodule {
	o := &submodule{
		name:   name,
		core:   inst.Core,
		dotgit: nil,
		gitdir: nil,
	}
	return o
}

// List ...
func (inst *Facade) List() []store.Submodule {

	dst := make([]store.Submodule, 0)

	return dst
}
