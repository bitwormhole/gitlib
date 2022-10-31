package others

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/store"
)

// CommandServiceManagerImpl ...
type CommandServiceManagerImpl struct {
	Context *store.Context

	table map[string]*instructions.ServiceRegistration
}

func (inst *CommandServiceManagerImpl) _Impl() instructions.ServiceManager {
	return inst
}

func (inst *CommandServiceManagerImpl) getTable() map[string]*instructions.ServiceRegistration {
	t := inst.table
	if t == nil {
		t = inst.loadTable()
		inst.table = t
	}
	return t
}

func (inst *CommandServiceManagerImpl) loadTable() map[string]*instructions.ServiceRegistration {
	src := inst.Context.Services
	dst := make(map[string]*instructions.ServiceRegistration)
	for _, r1 := range src {
		mid := r1.ListRegistrations()
		for _, r2 := range mid {
			name := r2.Name
			dst[name] = r2
		}
	}
	return dst
}

// Find ...
func (inst *CommandServiceManagerImpl) Find(name string) (instructions.Service, error) {
	t := inst.getTable()
	reg := t[name]
	if reg != nil {
		ser := reg.Service
		if ser != nil {
			return ser, nil
		}
	}
	return nil, errors.New("no git-command-service for name: " + name)
}
