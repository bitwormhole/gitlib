package submodules

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/utils"
	"github.com/bitwormhole/starter/collection"
	"github.com/starter-go/afs"
)

type loader struct{}

func (inst *loader) load(c *store.Core) ([]*submodule, error) {
	p, err := inst.loadProperties(c)
	if err != nil {
		return nil, err
	}
	return inst.loadModuleList(c, p)
}

func (inst *loader) loadProperties(core *store.Core) (collection.Properties, error) {

	layout := core.Layout
	dotgit := layout.DotGit()
	config := layout.Config()

	if dotgit == nil || config == nil {
		return nil, fmt.Errorf("bad git repository")
	}

	file1 := dotgit.GetParent().GetChild(".gitmodules")
	file2 := config
	filelist := []afs.Path{file1, file2}
	dst := collection.CreateProperties()

	for _, file := range filelist {
		inst.loadPropertiesFromFile(core, file, dst)
	}

	return dst, nil
}

func (inst *loader) loadPropertiesFromFile(core *store.Core, file afs.Path, dst collection.Properties) error {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}
	_, err = collection.ParseProperties(text, dst)
	return err
}

func (inst *loader) loadModule(core *store.Core, p collection.Properties, prefix string, name string) (*submodule, error) {

	g := p.Getter()
	key1 := prefix + name

	o := &submodule{}
	o.core = core
	o.name = name
	o.url = g.GetString(key1+".url", "")
	o.path = g.GetString(key1+".path", "")
	o.active = g.GetBool(key1+".active", false)

	if o.url == "" {
		return nil, fmt.Errorf("no property: " + key1 + ".url")
	}

	if o.path == "" {
		return nil, fmt.Errorf("no property: " + key1 + ".path")
	}

	if o.active {
		err := inst.makeModuleActived(o)
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

func (inst *loader) loadModuleList(core *store.Core, p collection.Properties) ([]*submodule, error) {

	const (
		prefix = "submodule."
		suffix = ".url"
	)

	list := make([]*submodule, 0)
	namelist := (&utils.PropertiesTools{}).ListNames(p, prefix, suffix)
	for _, name := range namelist {
		m, err := inst.loadModule(core, p, prefix, name)
		if err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}

func (inst *loader) makeModuleActived(m *submodule) error {

	c := m.core
	path := m.path
	layout := c.Layout
	wkdir := layout.Workspace()
	repodir := layout.Repository()

	dotgit := wkdir.GetChild("./" + path + "/.git")
	subconfig := repodir.GetChild("./modules/" + path + "/config")

	if !subconfig.IsFile() {
		return fmt.Errorf("bad git-submodule: no sub-config file")
	}
	if !dotgit.IsFile() {
		return fmt.Errorf("bad git-submodule: no .git file")
	}

	m.subconfig = subconfig
	m.dotgit = dotgit
	return nil
}
