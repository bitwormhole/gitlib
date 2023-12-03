package cases

import (
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/units"
)

// ContextLoadersTest ...
type ContextLoadersTest struct {

	//starter:component

	LibAgent repositories.LibAgent //starter:inject("#")

	Paths RepositoryPathProvider //starter:inject("#")

}

func (inst *ContextLoadersTest) _impl() units.Units {
	return inst
}

// Units ...
func (inst *ContextLoadersTest) Units(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:    "test-context-loaders",
		Enabled: true,
		Test:    inst.t,
	}
	list = append(list, r1)
	return list
}

func (inst *ContextLoadersTest) t() error {

	lib, err := inst.LibAgent.GetLib()
	if err != nil {
		return err
	}
	// system-context
	sysCtx := lib.SystemContext()

	// user-context
	userCtx, err := sysCtx.UserContextLoader.Load(&repositories.UserParams{
		Parent: sysCtx,
		Home:   nil, // use default
	})
	if err != nil {
		return err
	}

	// repo-context
	path, err := inst.Paths.GetRepoPath("3")
	if err != nil {
		return err
	}
	layout, err := sysCtx.Lib.Locator().Locate(path)
	if err != nil {
		return err
	}
	repoCtx, err := sysCtx.RepositoryContextLoader.Load(&repositories.RepositoryParams{
		Parent: userCtx,
		Layout: layout,
	})
	if err != nil {
		return err
	}

	repoCtx.Closer.Close()

	return nil
}
