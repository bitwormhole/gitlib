package lib

import (
	"fmt"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// GitlibAgent ...
type GitlibAgent struct {
	markup.Component `id:"git-lib-agent" class:"life"`

	CLI                cli.CLI                   `inject:"#cli"`
	ContextConfigurers []store.ContextConfigurer `inject:".git-context-configurer"`
	CoreConfigurers    []store.CoreConfigurer    `inject:".git-core-configurer"`

	lib store.Lib // cached
}

func (inst *GitlibAgent) _Impl() (store.LibAgent, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *GitlibAgent) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.start,
	}
}

func (inst *GitlibAgent) start() error {
	_, err := inst.innerGetlib()
	return err
}

func (inst *GitlibAgent) innerGetlib() (store.Lib, error) {
	i := inst.lib
	if i != nil {
		return i, nil
	}
	i, err := inst.loadLib()
	if err != nil {
		return nil, err
	}
	inst.lib = i
	return i, nil
}

func (inst *GitlibAgent) loadLib() (store.Lib, error) {

	cfg := &store.ContextConfiguration{}
	cfg.Factory = &support.DefaultContextFactory{}

	cfg.ContextConfigurers = inst.ContextConfigurers
	cfg.CoreConfigurers = inst.CoreConfigurers

	cfg.CLI = inst.CLI
	cfg.CLIConfig = nil
	cfg.UseCLI = true

	/////////////////////////////////////////////////
	// lib := gitlib.New(cfg)

	storeContext, err := cfg.Factory.Create(cfg)
	if err != nil {
		return nil, err
	}
	lib := storeContext.Lib
	if lib == nil {
		return nil, fmt.Errorf("lib is nil")
	}
	return lib, nil
}

// GetLib ...
func (inst *GitlibAgent) GetLib() (store.Lib, error) {
	return inst.innerGetlib()
}
