package commands

import (
	"bitwormhole.com/starter/cli"
)

// GitAdd ...
// 参考 https://git-scm.com/docs/git-add
type GitAdd struct {
}

func (inst *GitAdd) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitAdd) name() string {
	return "git-add"
}

// GetHandlers ...
func (inst *GitAdd) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitAdd) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "todo...",
		Usage:   "todo...",
		Content: "todo...",
	}
}

func (inst *GitAdd) handle(t *cli.Task) error {

	// ctx := t.Context
	// task := instructions.NewAdd(ctx)
	// todo ...
	// return task.Run()

	// 暂时 delegate to go-git
	t.Command = "go-git-add"
	return t.Client.Run(t)
}
