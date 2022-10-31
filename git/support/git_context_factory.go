package support

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/config"
	"github.com/bitwormhole/gitlib/git/store"
)

// DefaultContextFactory ...
type DefaultContextFactory struct {
}

func (inst *DefaultContextFactory) _Impl() store.ContextFactory {
	return inst
}

func (inst *DefaultContextFactory) String() string {
	return "DefaultContextFactory"
}

// Create ...
func (inst *DefaultContextFactory) Create(cfg *store.ContextConfiguration) (*store.Context, error) {
	ctx := &store.Context{}

	err := inst.config(ctx, cfg)
	if err != nil {
		return nil, err
	}

	err = inst.initCLI(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ctx.CoreConfigurers = cfg.CoreConfigurers
	return ctx, nil
}

func (inst *DefaultContextFactory) config(ctx *store.Context, cfg *store.ContextConfiguration) error {
	confs := cfg.ContextConfigurers
	for _, conf := range confs {
		err := conf.Configure(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *DefaultContextFactory) initCLI(ctx *store.Context, cfg *store.ContextConfiguration) error {
	if !cfg.UseCLI {
		return nil
	}
	theCliConfig := cfg.CLIConfig
	theCli := cfg.CLI
	if theCli == nil {
		if theCliConfig == nil {
			theCliConfig = config.GetDefaultConfiguration()
		}
		theCli = cli.New(theCliConfig)
	}
	ctx.CLI = theCli
	return nil
}
