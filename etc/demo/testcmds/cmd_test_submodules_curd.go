package testcmds

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
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

	lib, err := inst.LA.GetLib()
	if err != nil {
		return err
	}

	wd := lib.FS().NewPath(inst.WD)
	layout, err := lib.RepositoryLocator().Locate(wd)
	if err != nil {
		return err
	}

	repo, err := lib.RepositoryLoader().Load(layout)
	if err != nil {
		return err
	}

	list := repo.Submodules().List()
	for _, wt := range list {
		vlog.Info("find submodule: ", wt.Name())
	}

	return nil
}
