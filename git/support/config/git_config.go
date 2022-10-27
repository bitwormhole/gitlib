package config

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
)

type simpleConfig struct {
	path       afs.Path
	properties map[string]string
}

func (inst *simpleConfig) _Impl() (store.Config, store.RepositoryConfiguration) {
	return inst, inst
}

func (inst *simpleConfig) NodeType() store.NodeType {
	return store.NodeConfig
}

func (inst *simpleConfig) init(path afs.Path) {
	inst.properties = make(map[string]string)
	inst.path = path
}

func (inst *simpleConfig) Import(src map[string]string) {
	dst := inst.properties
	if dst == nil {
		dst = make(map[string]string)
		inst.properties = dst
	}
	for k, v := range src {
		dst[k] = v
	}
}

func (inst *simpleConfig) Export() map[string]string {
	dst := make(map[string]string)
	src := inst.properties
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (inst *simpleConfig) GetProperty(name string) string {
	return inst.properties[name]
}

func (inst *simpleConfig) SetProperty(name, value string) {
	inst.properties[name] = value
}

func (inst *simpleConfig) Path() afs.Path {
	return inst.path
}

func (inst *simpleConfig) Clear() {
	inst.properties = make(map[string]string)
}

func (inst *simpleConfig) Save(se store.Session) error {
	return se.SaveConfig(inst)
}

func (inst *simpleConfig) Load(se store.Session) error {
	return se.LoadConfig(inst)
}
