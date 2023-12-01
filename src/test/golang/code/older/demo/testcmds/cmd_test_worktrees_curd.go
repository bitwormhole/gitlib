package testcmds

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// TestWorktreesCURD ...
type TestWorktreesCURD struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *TestWorktreesCURD) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestWorktreesCURD) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-worktrees",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestWorktreesCURD) run(task *cli.Task) error {

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

	list := repo.Worktrees().List()
	for _, wt := range list {
		vlog.Info("find worktree: ", wt.Name())
	}

	return nil
}
