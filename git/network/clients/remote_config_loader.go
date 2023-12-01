package clients

import (
	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/repositories"
)

// RemoteConfigLoader ...
type RemoteConfigLoader struct {
	context *Context
}

// Load ...
func (inst *RemoteConfigLoader) Load() error {
	cfg := inst.context.Repository.Config()
	cfg = cfg.FindByScope(repositories.ConfigScopeMix)

	err := inst.loadRemotes(cfg)
	if err != nil {
		return err
	}

	err = inst.loadBranches(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (inst *RemoteConfigLoader) loadRemotes(cc repositories.ConfigChain) error {

	cfg := cc.Config()
	loader := gitconfig.ConfigLoader{}
	idx := loader.GetIndex(cfg)
	namelist := idx.ListNames("remote")
	dst := make(map[string]*gitconfig.Remote)

	for _, name := range namelist {
		remote := loader.LoadRemote(cfg, name)
		dst[name] = remote
	}

	inst.context.Remotes = dst
	return nil
}

func (inst *RemoteConfigLoader) loadBranches(cc repositories.ConfigChain) error {

	cfg := cc.Config()
	loader := gitconfig.ConfigLoader{}
	idx := loader.GetIndex(cfg)
	namelist := idx.ListNames("branch")
	dst := make(map[string]*gitconfig.Branch)

	for _, name := range namelist {
		branch := loader.LoadBranch(cfg, name)
		dst[name] = branch
	}

	inst.context.Branches = dst
	return nil
}
