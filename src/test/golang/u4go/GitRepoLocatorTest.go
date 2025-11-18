package u4go

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bitwormhole/gitlib/v1/repositories"
	"github.com/starter-go/units"
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/libafs"
)

type GitRepoLocatorTest struct {

	//starter:component

	_as func(units.Units) //starter:as(".")
}

// Units implements units.Units.
func (inst *GitRepoLocatorTest) Units(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:    "test-git-repo-locator",
		Enabled: true,
		Test:    inst.run,
	}
	list = append(list, r1)
	return list
}

func (inst *GitRepoLocatorTest) _impl() units.Units {
	return inst
}

func (inst *GitRepoLocatorTest) run() error {

	afs.Init(libafs.DefaultDriver())

	fsys := afs.Default()
	locator := new(repositories.Locator)
	wdp := inst.getWkdirPath()
	wd := fsys.GetNode(wdp)

	layout, err := locator.Locate(wd.(afs.Directory))
	if err != nil {
		return err
	}

	msg := inst.stringifyLayout(layout)
	fmt.Println("layout:\n", msg)

	return nil
}

func (inst *GitRepoLocatorTest) getWkdirPath() afs.Path {

	args := os.Args

	for i, it := range args {
		if i == 0 {
			continue
		}
		it = strings.TrimSpace(it)
		if strings.HasPrefix(it, "-") {
			continue
		}
		return afs.Path(it)
	}

	return "~/tmp/no/this/path/"
}

func (inst *GitRepoLocatorTest) stringifyLayout(layout repositories.Layout) string {

	table := map[string]string{}
	keys := []string{}
	builder := new(strings.Builder)

	table["spec:wkdir"] = inst.getPathString(layout.WD())
	table["spec:repo"] = inst.getPathString(layout.Repository())
	table["spec:worktree"] = inst.getPathString(layout.CurrentWorktree())

	table["file:.git"] = inst.getPathString(layout.DotGitFile())
	table["file:config"] = inst.getPathString(layout.Config())
	table["file:description"] = inst.getPathString(layout.Description())
	table["file:HEAD"] = inst.getPathString(layout.HEAD())
	table["file:index"] = inst.getPathString(layout.Index())
	table["file:packed-refs"] = inst.getPathString(layout.PackedRefs())

	table["dir:.git"] = inst.getPathString(layout.DotGitDir())
	table["dir:objects"] = inst.getPathString(layout.Objects())
	table["dir:refs"] = inst.getPathString(layout.Refs())
	table["dir:worktrees"] = inst.getPathString(layout.Worktrees())
	table["dir:modules"] = inst.getPathString(layout.Modules())
	table["dir:branches"] = inst.getPathString(layout.Branches())
	table["dir:hooks"] = inst.getPathString(layout.Hooks())
	table["dir:logs"] = inst.getPathString(layout.Logs())
	table["dir:info"] = inst.getPathString(layout.Info())

	for k := range table {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	const tabLenW = 22
	builder.WriteRune('\n')

	for _, k := range keys {
		tabLenH := tabLenW - len(k)
		v := table[k]
		for ; tabLenH > 0; tabLenH-- {
			builder.WriteRune('.')
		}
		builder.WriteString(k)
		builder.WriteString(" = ")
		builder.WriteString(v)
		builder.WriteRune('\n')
	}

	return builder.String()
}

func (inst *GitRepoLocatorTest) getPathString(p afs.Node) string {
	if p == nil {
		return ""
	}
	return p.String()
}
