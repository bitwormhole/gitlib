package demo

import (
	"context"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// TestPoint ...
type TestPoint struct {
	markup.Component `class:"life"`

	Agent   store.LibAgent `inject:"#git-lib-agent"`
	Command string         `inject:"${test.gitlib.command}"`
	WD      string         `inject:"${test.repo.path}"`
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

	cmd := inst.Command
	wd := inst.WD

	return cli2.GetClient().Run(&cli.Task{
		Context: ctx,
		Command: cmd,
		WD:      wd,
	})
}
