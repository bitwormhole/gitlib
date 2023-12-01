package commands

import (
	"bitwormhole.com/starter/cli"
)

// GitCmd ...
type GitCmd struct {
}

func (inst *GitCmd) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitCmd) name() string {
	return "git"
}

// GetHandlers ...
func (inst *GitCmd) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitCmd) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "Git Commands",
		Usage:   "git [sub-command] ...",
		Content: "todo...",
	}
}

func (inst *GitCmd) handle(t *cli.Task) error {
	args := t.Arguments
	if len(args) < 1 {
		args = []string{"help"}
	}
	t.Command = "git-" + args[0]
	t.Arguments = args[1:]
	client := t.Client
	return client.Run(t)
}
