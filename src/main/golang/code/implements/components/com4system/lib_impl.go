package com4system

import (
	"context"

	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// LibComReg ...
type LibComReg struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *LibComReg) _impl() repositories.ComponentRegistry {
	return inst
}

// ListRegistrations ...
func (inst *LibComReg) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.create,
	}
	list := make([]*repositories.ComponentRegistration, 0)
	list = append(list, r1)
	return list
}

func (inst *LibComReg) create(ctx *repositories.SystemContext) (any, error) {
	com := &libImpl{context: ctx}
	ctx.Facade = com
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type libImpl struct {
	context *repositories.SystemContext
}

func (inst *libImpl) _impl() repositories.Lib { return inst }

func (inst *libImpl) FS() afs.FS {
	return inst.context.FS
}

func (inst *libImpl) Loader() repositories.Loader {
	return inst.context.RepositoryLoader
}

func (inst *libImpl) Finder() repositories.Finder {
	return inst.context.RepositoryFinder
}

func (inst *libImpl) Locator() repositories.Locator {
	return inst.context.RepositoryLocator
}

func (inst *libImpl) InstructionServiceManager() repositories.ServiceManager {
	panic("no impl")
}

func (inst *libImpl) Connectors() pktline.ConnectorManager {
	panic("no impl")
}

// 把这个 Lib 绑定到指定的 Context
func (inst *libImpl) Bind(cc context.Context) context.Context {
	panic("no impl")
}
