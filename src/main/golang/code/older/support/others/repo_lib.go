package others

import (
	"context"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
)

// LibImpl ...
type LibImpl struct {
	Context *store.Context
}

func (inst *LibImpl) _Impl() store.Lib {
	return inst
}

// FS ...
func (inst *LibImpl) FS() afs.FS {
	return inst.Context.FS
}

// RepositoryFinder ...
func (inst *LibImpl) RepositoryFinder() store.RepositoryFinder {
	return inst.Context.Finder
}

// RepositoryLocator ...
func (inst *LibImpl) RepositoryLocator() store.RepositoryLocator {
	return inst.Context.Locator
}

// RepositoryLoader ...
func (inst *LibImpl) RepositoryLoader() store.RepositoryLoader {
	return inst.Context.RepositoryLoader
}

// InstructionServiceManager ...
func (inst *LibImpl) InstructionServiceManager() store.ServiceManager {
	return inst.Context.ServiceManager
}

// Connectors ...
func (inst *LibImpl) Connectors() pktline.ConnectorManager {
	return inst.Context.ConnectorManager
}

// Bind ...
func (inst *LibImpl) Bind(cc context.Context) context.Context {
	cc = store.Bind(cc)
	b, err := store.GetBinding(cc)
	if err != nil {
		panic(err)
	}
	err = b.SetLib(inst)
	if err != nil {
		panic(err)
	}
	return cc
}

// GetCLI ...
func (inst *LibImpl) GetCLI(required bool) cli.CLI {
	o := inst.Context.CLI
	if o == nil && required {
		panic("no cli in this lib config")
	}
	return o
}
