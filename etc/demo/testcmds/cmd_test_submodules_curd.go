package testcmds

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// TestSubmodulesCURD ...
type TestSubmodulesCURD struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *TestSubmodulesCURD) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestSubmodulesCURD) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-submodules",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestSubmodulesCURD) run(task *cli.Task) error {

	return nil
}
