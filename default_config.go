package gitlib

import (
	"bitwormhole.com/starter/cli/config"
	"github.com/bitwormhole/gitlib/git/commands"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support"
	"github.com/bitwormhole/gitlib/git/support/services"
)

// GetDefaultConfiguration 取默认配置
func GetDefaultConfiguration() *store.ContextConfiguration {
	loader := defaultConfigurationLoader{}
	return loader.load()
}

////////////////////////////////////////////////////////////////////////////////

type defaultConfigurationLoader struct {
}

func (inst *defaultConfigurationLoader) load() *store.ContextConfiguration {

	cfg := &store.ContextConfiguration{}
	cfg.Factory = &support.DefaultContextFactory{}

	inst.fillContextConfigurers(cfg)
	inst.fillCoreConfigurers(cfg)
	inst.fillCLI(cfg)

	return cfg
}

func (inst *defaultConfigurationLoader) fillCoreConfigurers(cfg *store.ContextConfiguration) {
	list := cfg.CoreConfigurers
	list = append(list, &support.BaseCoreConfigurer{})
	cfg.CoreConfigurers = list
}

func (inst *defaultConfigurationLoader) fillContextConfigurers(cfg *store.ContextConfiguration) {
	list := cfg.ContextConfigurers
	list = append(list, &support.BaseContextConfigurer{})
	list = append(list, &services.Configurer{})
	cfg.ContextConfigurers = list
}

func (inst *defaultConfigurationLoader) fillCLI(cfg *store.ContextConfiguration) {

	clicfg := config.GetDefaultConfiguration()
	hlist := clicfg.Handlers

	hlist = append(hlist, &commands.GitAdd{})
	hlist = append(hlist, &commands.GitCommit{})
	hlist = append(hlist, &commands.GitInit{})
	hlist = append(hlist, &commands.GitPush{})
	hlist = append(hlist, &commands.GitStatus{})

	hlist = append(hlist, &commands.GitCmd{})
	clicfg.Handlers = hlist

	cfg.UseCLI = true
	cfg.CLIConfig = clicfg
}
