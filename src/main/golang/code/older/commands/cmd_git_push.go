package commands

import (
	"strings"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/arguments"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitPush ...
// 参考 https://git-scm.com/docs/git-push
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

	const nl = "\n"
	name := inst.name()
	usage := strings.Builder{}

	usage.WriteString("git push [--all | --mirror | --tags] [--follow-tags] [--atomic] [-n | --dry-run] [--receive-pack=<git-receive-pack>]")
	usage.WriteString("[--repo=<repository>] [-f | --force] [-d | --delete] [--prune] [-v | --verbose]")
	usage.WriteString("[-u | --set-upstream] [-o <string> | --push-option=<string>]")
	usage.WriteString("[--[no-]signed|--signed=(true|false|if-asked)]")
	usage.WriteString("[--force-with-lease[=<refname>[:<expect>]] [--force-if-includes]]")
	usage.WriteString("[--no-verify] [<repository> [<refspec>…​]]")

	return &cli.HelpInfo{
		Name:    name,
		Title:   "Update remote refs along with associated objects",
		Usage:   usage.String(),
		Content: "todo...",
	}
}

func (inst *GitPush) handle(t1 *cli.Task) error {

	const (
		optAll             = "--all"
		optMirror          = "--mirror"
		optTags            = "--tags"
		optFollowTags      = "--follow-tags"
		optAtomic          = "--atomic"
		optN               = "-n"
		optDryRun          = "--dry-run"
		optReceivePack     = "--receive-pack" // =<git-receive-pack>
		optRepo            = "--repo"         // =<repository>
		optF               = "-f"
		optForce           = "--force"
		optD               = "-d"
		optDelete          = "--delete"
		optPrune           = "--prune"
		optV               = "-v"
		optVerbose         = "--verbose"
		optU               = "-u"
		optSetUpstream     = "--set-upstream"
		optO               = "-o"            // <string>
		optPushOption      = "--push-option" // =<string>
		optSigned          = "--signed"      // =(true|false|if-asked)
		optNoSigned        = "--no-signed"
		optForceWithLease  = "--force-with-lease" // [=<refname>[:<expect>]
		optForceIfIncludes = "--force-if-includes"
		optNoVerify        = "--no-verify"
		//  [<repository> [<refspec>...]]
	)

	atb := arguments.NewTemplateBuilder()
	atb.AcceptOption(optAll, 0)
	atb.AcceptOption(optMirror, 0)
	atb.AcceptOption(optTags, 0)
	atb.AcceptOption(optFollowTags, 0)
	atb.AcceptOption(optAtomic, 0)
	atb.AcceptOption(optN, 0)
	atb.AcceptOption(optDryRun, 0)
	atb.AcceptOption(optReceivePack, 1)
	atb.AcceptOption(optRepo, 1)
	atb.AcceptOption(optF, 0)
	atb.AcceptOption(optForce, 0)
	atb.AcceptOption(optD, 0)
	atb.AcceptOption(optDelete, 0)
	atb.AcceptOption(optPrune, 0)
	atb.AcceptOption(optV, 0)
	atb.AcceptOption(optVerbose, 0)
	atb.AcceptOption(optU, 0)
	atb.AcceptOption(optSetUpstream, 0)
	atb.AcceptOption(optO, 1)
	atb.AcceptOption(optPushOption, 1)
	atb.AcceptOption(optSigned, 1)
	atb.AcceptOption(optNoSigned, 0)
	atb.AcceptOption(optForceWithLease, 1)
	atb.AcceptOption(optForceIfIncludes, 0)
	atb.AcceptOption(optNoVerify, 0)

	args := atb.Create().Parse(t1.Arguments)
	ctx := t1.Context
	t2 := instructions.NewPush(ctx)
	initMeta(&t2.Meta, t1)

	// options as simple bool
	t2.All = args.GetOption(optAll).Exists()
	t2.Mirror = args.GetOption(optMirror).Exists()
	t2.Tags = args.GetOption(optTags).Exists()
	t2.FollowTags = args.GetOption(optFollowTags).Exists()
	t2.Atomic = args.GetOption(optAtomic).Exists()
	t2.Prune = args.GetOption(optPrune).Exists()
	t2.Signed = args.GetOption(optSigned).Exists()
	t2.NoSigned = args.GetOption(optNoSigned).Exists()
	t2.ForceIfIncludes = args.GetOption(optForceIfIncludes).Exists()
	t2.NoVerify = args.GetOption(optNoVerify).Exists()

	// options with short name
	t2.DryRun = inst.existsOption(args, optN, optDryRun)
	t2.Force = inst.existsOption(args, optF, optForce)
	t2.Delete = inst.existsOption(args, optD, optDelete)
	t2.Verbose = inst.existsOption(args, optV, optVerbose)
	t2.SetUpstream = inst.existsOption(args, optU, optSetUpstream)

	// options as string
	t2.ReceivePack = args.GetOption(optReceivePack).Value(0).String()
	t2.Repo = args.GetOption(optRepo).Value(0).String()
	t2.PushOption = args.GetOption(optPushOption).Value(0).String()
	t2.ForceWithLease = args.GetOption(optForceWithLease).Value(0).String()

	t2.PushOption = inst.getPushOption(args, optO, optPushOption)

	// items
	items := args.Items()
	for _, arg := range items {
		str := arg.String()
		t2.RepositoryRefspecList = append(t2.RepositoryRefspecList, str)
	}

	return t2.Run()
}

func (inst *GitPush) existsOption(args arguments.Arguments, op ...string) bool {
	for _, name := range op {
		if args.GetOption(name).Exists() {
			return true
		}
	}
	return false
}

func (inst *GitPush) getPushOption(args arguments.Arguments, names ...string) string {
	for _, name := range names {
		op := args.GetOption(name)
		if op.Exists() {
			return op.Value(0).String()
		}
	}
	return ""
}
