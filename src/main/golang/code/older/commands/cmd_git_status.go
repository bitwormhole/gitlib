package commands

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitStatus ...
// 参考 https://git-scm.com/docs/git-status
type GitStatus struct {
}

func (inst *GitStatus) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitStatus) name() string {
	return "git-status"
}

// GetHandlers ...
func (inst *GitStatus) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitStatus) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "todo...",
		Usage:   "todo...",
		Content: "todo...",
	}
}

func (inst *GitStatus) handle(t1 *cli.Task) error {

	//
	// git status [<options>…​] [--] [<pathspec>…​]
	// -s | --short
	// -b | --branch
	// 	--show-stash
	// 	--porcelain[=<version>]
	// 	--long
	// 	-v | --verbose
	// 	-u[<mode>] | --untracked-files[=<mode>]
	// 	--ignore-submodules[=<when>]
	// 	--ignored[=<mode>]
	// 	-z
	// 	--column[=<选项>] | --no-column
	// 	--ahead-behind | --no-ahead-behind
	// 	--renames | --no-renames
	// 	--find-renames[=<n>]
	// 	<pathspec>…​
	//
	//////////////

	t2 := instructions.NewStatus(t1.Context)
	initMeta(&t2.Meta, t1)

	// todo ...

	return t2.Run()
}
