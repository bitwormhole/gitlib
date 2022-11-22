package config

import (
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/data/gitfmt"
	"github.com/bitwormhole/gitlib/git/store"
)

type simpleConfig struct {
	path       afs.Path
	properties map[string]string
	ignoreCase bool
}

func (inst *simpleConfig) _Impl() (store.Config, store.RepositoryConfig) {
	return inst, inst
}

func (inst *simpleConfig) NodeType() store.NodeType {
	return store.NodeConfig
}

func (inst *simpleConfig) init(p *store.ConfigChainParams) {
	inst.properties = make(map[string]string)
	inst.path = p.File
	inst.ignoreCase = p.IgnoreCase
}

func (inst *simpleConfig) prepareKey(k string) string {
	if inst.ignoreCase {
		return strings.ToLower(k)
	}
	return k
}

func (inst *simpleConfig) Import(src map[string]string) {
	dst := inst.properties
	if dst == nil {
		dst = make(map[string]string)
		inst.properties = dst
	}
	for k, v := range src {
		k = inst.prepareKey(k)
		dst[k] = v
	}
}

func (inst *simpleConfig) Export() map[string]string {
	dst := make(map[string]string)
	src := inst.properties
	for k, v := range src {
		k = inst.prepareKey(k)
		dst[k] = v
	}
	return dst
}

func (inst *simpleConfig) GetProperty(name string) string {
	name = inst.prepareKey(name)
	return inst.properties[name]
}

func (inst *simpleConfig) SetProperty(name, value string) {
	name = inst.prepareKey(name)
	inst.properties[name] = value
}

func (inst *simpleConfig) Path() afs.Path {
	return inst.path
}

func (inst *simpleConfig) Clear() {
	inst.properties = make(map[string]string)
}

func (inst *simpleConfig) Save() error {
	file := inst.path
	if file == nil {
		return nil
	}
	props := &git.Properties{}
	props.Import(inst.Export())
	text := gitfmt.FormatPropertiesWithSegment(props)
	opt := &afs.Options{}
	err := file.GetIO().WriteText(text, opt)
	if err != nil {
		return err
	}
	return nil
}

func (inst *simpleConfig) Load() error {
	file := inst.path
	if file == nil {
		return nil
	}
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}
	src, err := gitfmt.ParseProperties(text, nil)
	if err != nil {
		return err
	}
	inst.properties = nil
	inst.Import(src.Export(nil))
	return nil
}
