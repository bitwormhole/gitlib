package commands

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
)

// GitExample ...
type GitExample struct {
}

func (inst *GitExample) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitExample) name() string {
	return "git-example"
}

// GetHandlers ...
func (inst *GitExample) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitExample) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "todo...",
		Usage:   "todo...",
		Content: "todo...",
	}
}

func (inst *GitExample) handle(t *cli.Task) error {

	ctx := t.Context
	task := git.NewExample(ctx)

	// todo ...

	return task.Run()
}
