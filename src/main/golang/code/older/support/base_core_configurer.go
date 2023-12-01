package support

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support/index"
	"github.com/bitwormhole/gitlib/git/support/objects"
	"github.com/bitwormhole/gitlib/git/support/others"
	"github.com/bitwormhole/gitlib/git/support/refs"
	"github.com/bitwormhole/gitlib/git/support/sessions"
	"github.com/bitwormhole/gitlib/git/support/submodules"
	"github.com/bitwormhole/gitlib/git/support/worktrees"
)

////////////////////////////////////////////////////////////////////////////////

// BaseCoreConfigurer ...
type BaseCoreConfigurer struct{}

func (inst *BaseCoreConfigurer) _Impl() store.CoreConfigurer {
	return inst
}

// Configure ...
func (inst *BaseCoreConfigurer) Configure(c *store.Core) error {

	config, err := inst.loadConfig(c)
	if err != nil {
		return err
	}

	// c.WD = nil
	// c.Layout = nil
	// c.Context = nil

	c.Config = config

	c.Objects = &objects.GitObjectsImpl{Core: c}
	c.Refs = &refs.GitRefsImpl{Core: c}
	c.Repository = &others.GitRepositoryImpl{Core: c}
	c.Workspace = &others.GitWorkspaceFacade{Core: c}
	c.Head = &others.GitHeadImpl{Core: c}
	c.Index = &index.GitIndexImpl{Core: c}
	c.Worktrees = &worktrees.Facade{Core: c}
	c.Submodules = &submodules.Facade{Core: c}
	c.SessionFactory = &sessions.Factory{}

	c.Digest = nil
	c.PathMapping = nil
	c.Compression = nil

	return nil
}

func (inst *BaseCoreConfigurer) loadConfig(c *store.Core) (store.ConfigChain, error) {

	factory := c.Context.ConfigChainFactory
	layout := c.Layout
	configFile := layout.Config()

	rootConfig := factory.Root()
	params := &store.ConfigChainParams{
		File:       configFile,
		Parent:     rootConfig,
		Scope:      store.ConfigScopeRepository,
		Required:   true,
		IgnoreCase: true,
	}
	repoConfig := factory.Create(params)

	err := repoConfig.Load()
	if err != nil {
		return nil, err
	}

	err = inst.checkConfig(repoConfig)
	if err != nil {
		return nil, err
	}

	mixConfig := repoConfig.Mix()
	return mixConfig, nil
}

func (inst *BaseCoreConfigurer) checkConfig(cfg store.ConfigChain) error {

	const empty = ""
	cfg4repo := cfg.FindByScope(store.ConfigScopeRepository)
	config := cfg4repo.Config()
	count := 0

	names := []string{
		"core.logallrefupdates",
		"core.symlinks",
		"core.ignorecase",
		"core.bare",
		"core.repositoryformatversion",
		"core.filemode",
	}

	for _, name := range names {
		value := config.GetProperty(name)
		if value != empty {
			count++
		}
	}

	if count < 2 {
		path := config.Path().GetPath()
		return errors.New("bad repository config file, path=" + path)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////