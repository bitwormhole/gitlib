package config

import "github.com/bitwormhole/gitlib/git/repositories"

func mix(chain repositories.ConfigChain) (repositories.Config, error) {
	list := make([]repositories.Config, 0)
	p := chain
	for ; p != nil; p = p.Parent() {
		if p.Scope() == repositories.ConfigScopeMix {
			continue
		}
		list = append(list, p.Config())
	}
	mixer := configMixer{}
	mixer.init()
	for i := len(list) - 1; i >= 0; i-- {
		cfg := list[i]
		src := cfg.Export()
		mixer.putAll(src)
	}
	return mixer.create()
}

////////////////////////////////////////////////////////////////////////////////

type configMixer struct {
	table map[string]string
}

func (inst *configMixer) init() {
	inst.table = make(map[string]string)
}

func (inst *configMixer) putAll(src map[string]string) {
	dst := inst.table
	for k, v := range src {
		dst[k] = v
	}
}

func (inst *configMixer) create() (repositories.Config, error) {
	cfg := &simpleConfig{}
	cfg.properties = inst.table
	cfg.ignoreCase = true
	return cfg, nil
}
