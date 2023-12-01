package commands

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/arguments"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitInit ...
// 参考 https://git-scm.com/docs/git-init
type GitInit struct {
}

func (inst *GitInit) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitInit) name() string {
	return "git-init"
}

// GetHandlers ...
func (inst *GitInit) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitInit) GetHelp() *cli.HelpInfo {
	name := inst.name()
	return &cli.HelpInfo{
		Name:    name,
		Title:   "todo...",
		Usage:   "todo...",
		Content: "todo...",
	}
}

func (inst *GitInit) handle(t1 *cli.Task) error {

	// 参考 https://git-scm.com/docs/git-init
	//
	// git init [-q | --quiet] [--bare] [--template=<template-directory>]
	// [--separate-git-dir <git-dir>] [--object-format=<format>]
	// [-b <branch-name> | --initial-branch=<branch-name>]
	// [--shared[=<permissions>]] [<directory>]

	const (
		theQ              = "-q"                 //
		theQuiet          = "--quiet"            //
		theBare           = "--bare"             //
		theTemplate       = "--template"         // =<template-directory>
		theSeparateGitDir = "--separate-git-dir" // <git-dir>
		theObjectFormat   = "--object-format"    // =<format>
		theB              = "-b"                 // <branch-name>
		theInitialBranch  = "--initial-branch"   // =<branch-name>
		theShared         = "--shared"           // [=<permissions>]
	)

	tb := arguments.NewTemplateBuilder()
	tb.AcceptOption(theQ, 0)
	tb.AcceptOption(theQuiet, 0)
	tb.AcceptOption(theBare, 0)
	tb.AcceptOption(theTemplate, 1)
	tb.AcceptOption(theSeparateGitDir, 1)
	tb.AcceptOption(theObjectFormat, 1)
	tb.AcceptOption(theB, 1)
	tb.AcceptOption(theInitialBranch, 1)
	tb.AcceptOption(theShared, 1)
	args := tb.Create().Parse(t1.Arguments)

	aQ := args.GetOption(theQ)
	aQuiet := args.GetOption(theQuiet)
	aBare := args.GetOption(theBare)
	aTemplate := args.GetOption(theTemplate)
	aSeparateGitDir := args.GetOption(theSeparateGitDir)
	aObjectFormat := args.GetOption(theObjectFormat)
	aB := args.GetOption(theB)
	aInitialBranch := args.GetOption(theInitialBranch)
	aShared := args.GetOption(theShared)
	aDirectory := args.GetItem(0)

	// for instruction
	t2 := instructions.NewInit(t1.Context)
	initMeta(&t2.Meta, t1)

	t2.Quiet = aQ.Exists() || aQuiet.Exists()
	t2.Bare = aBare.Exists()
	t2.TemplateDirectory = aTemplate.Value(0).String()
	t2.SeparateGitDir = aSeparateGitDir.Value(0).String()
	t2.ObjectFormat = aObjectFormat.Value(0).String()

	if aInitialBranch.Exists() {
		t2.BranchName = aInitialBranch.Value(0).String()
	} else if aB.Exists() {
		t2.BranchName = aB.Value(0).String()
	}

	if aShared.Exists() {
		t2.Shared = true
		t2.Permissions = aShared.Value(0).String()
	}

	if aDirectory.Exists() {
		t2.Directory = aDirectory.String()
	}

	return t2.Run()
}
