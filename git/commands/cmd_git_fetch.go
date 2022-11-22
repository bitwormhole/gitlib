package commands

import (
	"strings"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/arguments"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// GitFetch ...
// 参考 https://git-scm.com/docs/git-fetch
type GitFetch struct {
}

func (inst *GitFetch) _Impl() (cli.HandlerRegistry, cli.Help) {
	return inst, inst
}

func (inst *GitFetch) name() string {
	return "git-fetch"
}

// GetHandlers ...
func (inst *GitFetch) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Help:    inst,
		Handler: inst.handle,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *GitFetch) GetHelp() *cli.HelpInfo {
	const nl = "\n"

	name := inst.name()

	usage := strings.Builder{}
	usage.WriteString("'git fetch' [<options>] [<repository> [<refspec>...]]" + nl)
	usage.WriteString("'git fetch' [<options>] <group>" + nl)
	usage.WriteString("'git fetch' --multiple [<options>] [(<repository> | <group>)...]" + nl)
	usage.WriteString("'git fetch' --all [<options>]" + nl)

	return &cli.HelpInfo{
		Name:    name,
		Title:   "Download objects and refs from another repository",
		Usage:   usage.String(),
		Content: "todo...",
	}
}

func (inst *GitFetch) handle(t1 *cli.Task) error {

	const (
		optAll                     = "--all"
		optA                       = "-a"
		optAppend                  = "--append"
		optDepth                   = "--depth"           // =<depth>
		optDeepen                  = "--deepen"          // =<depth>
		optShallowSince            = "--shallow-since"   // =<date>
		optShallowExclude          = "--shallow-exclude" // =<revision>
		optUnshallow               = "--unshallow"
		optUpdateShallow           = "--update-shallow"
		optNegotiationTip          = "--negotiation-tip" // =<commit|glob>
		optNegotiateOnly           = "--negotiate-only"
		optDryRun                  = "--dry-run"
		optWriteFetchHead          = "--write-fetch-head"
		optNoWriteFetchHead        = "--no-write-fetch-head"
		optF                       = "-f"
		optForce                   = "--force"
		optK                       = "-k"
		optKeep                    = "--keep"
		optMultiple                = "--multiple"
		optNoAutoMaintenance       = "--no-auto-maintenance"
		optAutoMaintenance         = "--auto-maintenance"
		optNoAutoGC                = "--no-auto-gc"
		optAutoGC                  = "--auto-gc"
		optNoWriteCommitGraph      = "--no-write-commit-graph"
		optWriteCommitGraph        = "--write-commit-graph"
		optPrefetch                = "--prefetch"
		optP                       = "-p"
		optPrune                   = "--prune"
		optPruneTags               = "--prune-tags"
		optN                       = "-n"
		optNoTags                  = "--no-tags"
		optRefetch                 = "--refetch"
		optRefmap                  = "--refmap" // =<refspec>
		optT                       = "-t"
		optTags                    = "--tags"
		optRecurseSubmodule        = "--recurse-submodules" // =[=yes|on-demand|no]
		optJ                       = "-j"
		optJobs                    = "--jobs" // =<n>
		optNoRecurseSubmodules     = "--no-recurse-submodules"
		optSetUpstream             = "--set-upstream"
		optSubmodulePrefix         = "--submodule-prefix"          // =<path>
		optRecurseSubmoduleDefault = "--recurse-submodule-default" // =[yes|on-demand]
		optU                       = "-u"
		optUpdateHeadOK            = "--update-head-ok"
		optUploadPack              = "--upload-pack" // <upload-pack>
		optQ                       = "-q"
		optQuiet                   = "--quiet"
		optV                       = "-v"
		optVerbose                 = "--verbose"
		optProgress                = "--progress"
		optO                       = "-o"              // <option>
		optServerOption            = "--server-option" // =<option>
		optShowForcedUpdates       = "--show-forced-updates"
		optNoShowForcedUpdates     = "--no-show-forced-updates"
		opt4                       = "-4"
		optIpv4                    = "--ipv4"
		opt6                       = "-6"
		optIpv6                    = "--ipv6"
		optStdin                   = "--stdin"
	)

	atb := arguments.NewTemplateBuilder()
	atb.AcceptOption(opt4, 0)
	atb.AcceptOption(opt6, 0)
	atb.AcceptOption(optA, 0)
	atb.AcceptOption(optAll, 0)
	atb.AcceptOption(optAppend, 0)
	atb.AcceptOption(optAutoGC, 0)
	atb.AcceptOption(optAutoMaintenance, 0)
	atb.AcceptOption(optDeepen, 0)
	atb.AcceptOption(optDepth, 0)
	atb.AcceptOption(optDryRun, 0)
	atb.AcceptOption(optF, 0)
	atb.AcceptOption(optForce, 0)
	atb.AcceptOption(optIpv4, 0)
	atb.AcceptOption(optIpv6, 0)
	atb.AcceptOption(optJ, 0)
	atb.AcceptOption(optJobs, 0)
	atb.AcceptOption(optK, 0)
	atb.AcceptOption(optKeep, 0)
	atb.AcceptOption(optMultiple, 0)
	atb.AcceptOption(optN, 0)
	atb.AcceptOption(optNegotiateOnly, 0)
	atb.AcceptOption(optNegotiationTip, 0)
	atb.AcceptOption(optNoAutoGC, 0)
	atb.AcceptOption(optNoAutoMaintenance, 0)
	atb.AcceptOption(optNoRecurseSubmodules, 0)
	atb.AcceptOption(optNoShowForcedUpdates, 0)
	atb.AcceptOption(optNoTags, 0)
	atb.AcceptOption(optNoWriteCommitGraph, 0)
	atb.AcceptOption(optNoWriteFetchHead, 0)
	atb.AcceptOption(optO, 0)
	atb.AcceptOption(optP, 0)
	atb.AcceptOption(optPrefetch, 0)
	atb.AcceptOption(optProgress, 0)
	atb.AcceptOption(optPrune, 0)
	atb.AcceptOption(optPruneTags, 0)
	atb.AcceptOption(optQ, 0)
	atb.AcceptOption(optQuiet, 0)
	atb.AcceptOption(optRecurseSubmodule, 0)
	atb.AcceptOption(optRecurseSubmoduleDefault, 0)
	atb.AcceptOption(optRefetch, 0)
	atb.AcceptOption(optRefmap, 0)
	atb.AcceptOption(optServerOption, 0)
	atb.AcceptOption(optSetUpstream, 0)
	atb.AcceptOption(optShallowExclude, 0)
	atb.AcceptOption(optShallowSince, 0)
	atb.AcceptOption(optShowForcedUpdates, 0)
	atb.AcceptOption(optStdin, 0)
	atb.AcceptOption(optSubmodulePrefix, 0)
	atb.AcceptOption(optT, 0)
	atb.AcceptOption(optTags, 0)
	atb.AcceptOption(optU, 0)
	atb.AcceptOption(optUnshallow, 0)
	atb.AcceptOption(optUpdateHeadOK, 0)
	atb.AcceptOption(optUpdateShallow, 0)
	atb.AcceptOption(optUploadPack, 0)
	atb.AcceptOption(optV, 0)
	atb.AcceptOption(optVerbose, 0)
	atb.AcceptOption(optWriteCommitGraph, 0)
	atb.AcceptOption(optWriteFetchHead, 0)

	args := atb.Create().Parse(t1.Arguments)
	ctx := t1.Context
	t2 := instructions.NewFetch(ctx)
	initMeta(&t2.Meta, t1)

	// todo ...

	t2.All = args.GetOption(optAll).Exists()
	t2.Append = args.GetOption(optAppend).Exists()
	t2.Multiple = args.GetOption(optMultiple).Exists()

	// items
	items := args.Items()
	for _, arg := range items {
		str := arg.String()
		t2.Items = append(t2.Items, str)
	}

	return t2.Run()
}
