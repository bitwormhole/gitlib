package demo

import (
	"context"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// TestPoint ...
type TestPoint struct {
	markup.Component `class:"life"`

	Context application.Context `inject:"context"`
	Agent   store.LibAgent      `inject:"#git-lib-agent"`
	CmdKey  string              `inject:"${test.gitlib.command}"` // use 'command.n' to get content
	WD      string              `inject:"${test.repo.path}"`
}

func (inst *TestPoint) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *TestPoint) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.start,
	}
}

func (inst *TestPoint) start() error {

	lib, err := inst.Agent.GetLib()
	if err != nil {
		return err
	}

	cli2 := lib.GetCLI(true)
	ctx := context.Background()
	ctx = lib.Bind(ctx)
	ctx = cli2.Bind(ctx)

	wd := inst.WD
	cmd, err := inst.getCommand()
	if err != nil {
		return err
	}

	vlog.Info("do test with command: ", cmd)
	vlog.Info("                  wd: ", wd)

	return cli2.GetClient().Run(&cli.Task{
		Context: ctx,
		Command: cmd,
		WD:      wd,
	})
}

func (inst *TestPoint) getCommand() (string, error) {
	key := "command." + inst.CmdKey
	ctx := inst.Context
	return ctx.GetProperties().GetPropertyRequired(key)
}
