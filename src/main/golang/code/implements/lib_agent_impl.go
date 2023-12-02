package implements

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/application"
)

// LibAgentImpl ...
type LibAgentImpl struct {

	//starter:component
	_as func(repositories.LibAgent) //starter:as("#")

	Loader repositories.SystemContextLoader //starter:inject("#")

	cached *repositories.SystemContext
}

func (inst *LibAgentImpl) _impl() (repositories.LibAgent, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *LibAgentImpl) Life() *application.Life {
	return &application.Life{
		OnDestroy: inst.close,
	}
}

func (inst *LibAgentImpl) close() error {
	sc := inst.cached
	if sc == nil {
		return nil
	}
	cl := sc.Closer
	if cl == nil {
		return nil
	}
	return cl.Close()
}

// GetLib ...
func (inst *LibAgentImpl) GetLib() (repositories.Lib, error) {
	sc, err := inst.getSC()
	if err != nil {
		return nil, err
	}
	lib := sc.Facade
	if lib == nil {
		return nil, fmt.Errorf("bad repositories.SystemContext, facade is nil")
	}
	return lib, nil
}

func (inst *LibAgentImpl) getSC() (*repositories.SystemContext, error) {
	sysctx := inst.cached
	if sysctx != nil {
		return sysctx, nil
	}
	sysctx, err := inst.loadSystemContext()
	if err != nil {
		return nil, err
	}
	inst.cached = sysctx
	return sysctx, nil
}

func (inst *LibAgentImpl) loadSystemContext() (*repositories.SystemContext, error) {
	return inst.Loader.Load()
}
