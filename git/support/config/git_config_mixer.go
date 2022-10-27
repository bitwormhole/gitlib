package config

import "github.com/bitwormhole/gitlib/git/store"

func mix(chain store.ConfigChain) (store.Config, error) {
	list := make([]store.Config, 0)
	p := chain
	for ; p != nil; p = p.Parent() {
		if p.Scope() == store.ConfigScopeMix {
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

func (inst *configMixer) create() (store.Config, error) {
	cfg := &simpleConfig{}
	cfg.properties = inst.table
	return cfg, nil
}
