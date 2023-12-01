package gitconfig

import (
	"strconv"
	"strings"

	"github.com/bitwormhole/gitlib/git/repositories"
)

// ConfigLoader ...
type ConfigLoader struct {
}

func (inst *ConfigLoader) getString(cfg repositories.Config, name string, key1 KeyTemplate) string {
	const namePh = NamePlaceholder
	key2 := key1.String()
	temp := key2
	if strings.Contains(temp, namePh) {
		name = strings.TrimSpace(name)
		key2 = strings.Replace(temp, namePh, name, 1)
	}
	return cfg.GetProperty(key2)
}

func (inst *ConfigLoader) getInt(cfg repositories.Config, name string, key KeyTemplate) int {
	str := inst.getString(cfg, name, key)
	if str == "" {
		return 0
	}
	n, _ := strconv.Atoi(str)
	return n
}

func (inst *ConfigLoader) getBool(cfg repositories.Config, name string, key KeyTemplate) bool {
	str := inst.getString(cfg, name, key)
	if str == "" {
		return false
	}
	b, _ := strconv.ParseBool(str)
	return b
}

func (inst *ConfigLoader) exists(value string) bool {
	return value != ""
}

// GetIndex ...
func (inst *ConfigLoader) GetIndex(cfg repositories.Config) ConfigIndex {
	builder := myConfigIndexBuilder{}
	builder.init(cfg)
	return builder.create()
}

/////////////////////////////////////////////////////

// LoadCore ...
func (inst *ConfigLoader) LoadCore(cfg repositories.Config) *Core {
	const name = ""
	res := &Core{
		Bare:                    inst.getBool(cfg, name, CoreBare),
		RepositoryFormatVersion: inst.getInt(cfg, name, CoreRepositoryFormatVersion),
	}
	ex := inst.getString(cfg, name, CoreRepositoryFormatVersion)
	res.Exists = inst.exists(ex)
	return res
}

// LoadRemote ...
func (inst *ConfigLoader) LoadRemote(cfg repositories.Config, name string) *Remote {
	res := &Remote{
		Name:  name,
		URL:   inst.getString(cfg, name, RemoteNameURL),
		Fetch: inst.getString(cfg, name, RemoteNameFetch),
	}
	res.Exists = inst.exists(res.URL)
	return res
}

// LoadBranch ...
func (inst *ConfigLoader) LoadBranch(cfg repositories.Config, name string) *Branch {
	res := &Branch{
		Name:        name,
		Merge:       inst.getString(cfg, name, BranchNameMerge),
		Remote:      inst.getString(cfg, name, BranchNameRemote),
		PushRemote:  inst.getString(cfg, name, BranchNamePushRemote),
		Description: inst.getString(cfg, name, BranchNameDescription),
	}
	res.Exists = inst.exists(res.Merge)
	return res
}
