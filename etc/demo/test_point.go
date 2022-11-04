package demo

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// TestPoint ...
type TestPoint struct {
	markup.Component `class:"life"`

	Agent   store.LibAgent `inject:"#git-lib-agent"`
	Command string         `inject:"${test.gitlib.command}"`
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

	cli := lib.GetCLI(true)
	ctx := context.Background()
	ctx = lib.Bind(ctx)
	ctx = cli.Bind(ctx)

	cmd := inst.Command

	return cli.GetClient().RunCCA(ctx, cmd, nil)
}
