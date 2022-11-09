package commands

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitPush ...
type GitPush struct {
}

func (inst *GitPush) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitPush) name() string {
	return "git-push"
}

// GetHandlers ...
func (inst *GitPush) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitPush) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "todo...",
		Usage:   "todo...",
		Content: "todo...",
	}
}

func (inst *GitPush) handle(t *cli.Task) error {

	ctx := t.Context
	task := instructions.NewPush(ctx)

	// todo ...

	return task.Run()
}
