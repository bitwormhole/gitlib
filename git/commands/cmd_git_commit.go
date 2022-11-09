package commands

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitCommit ...
type GitCommit struct {
}

func (inst *GitCommit) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitCommit) name() string {
	return "git-commit"
}

// GetHandlers ...
func (inst *GitCommit) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitCommit) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "todo...",
		Usage:   "todo...",
		Content: "todo...",
	}
}

func (inst *GitCommit) handle(t *cli.Task) error {

	ctx := t.Context
	task := instructions.NewCommit(ctx)

	// todo ...

	return task.Run()
}
