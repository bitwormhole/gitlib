package gitconfig

import (
	"strings"

	"github.com/bitwormhole/gitlib/git/store"
)

// ConfigIndex ... [aaa.bbb.ccc] is [<class>.<name>.<field>]
type ConfigIndex interface {
	GetConfig() store.Config

	ListClasses() []string

	ListNames(aClass string) []string

	// key like 'a.b.*'
	Contains(key string) bool
}

////////////////////////////////////////////////////////////////////////////////

type myConfigIndex struct {
	config  store.Config
	classes map[string]*myConfigClass
}

func (inst *myConfigIndex) _Impl() ConfigIndex {
	return inst
}

func (inst *myConfigIndex) GetConfig() store.Config {
	return inst.config
}

// Contains ...
func (inst *myConfigIndex) Contains(key string) bool {

	parts := strings.Split(key, ".")
	p1 := parts[0]
	p2 := ""
	hasP2 := false
	if len(parts) > 1 {
		p2 = parts[1]
		hasP2 = true
	}

	t1 := inst.classes
	if t1 == nil {
		return false
	}
	clazz := t1[p1]
	if clazz == nil {
		return false
	}
	if !hasP2 {
		return true
	}
	t2 := clazz.objects
	if t2 == nil {
		return false
	}
	obj := t2[p2]
	return obj != nil
}

// ListClasses ...
func (inst *myConfigIndex) ListClasses() []string {
	src := inst.classes
	dst := make([]string, 0)
	for _, item := range src {
		dst = append(dst, item.className)
	}
	return dst
}

// ListNames ...
func (inst *myConfigIndex) ListNames(aClass string) []string {
	dst := make([]string, 0)
	t1 := inst.classes
	if t1 == nil {
		return dst
	}
	clazz := t1[aClass]
	if clazz == nil {
		return dst
	}
	t2 := clazz.objects
	if t2 == nil {
		return dst
	}
	for _, item := range t2 {
		dst = append(dst, item.name)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type myConfigClass struct {
	className string
	objects   map[string]*myConfigObject
}

func (inst *myConfigClass) init(cName string) {
	inst.className = cName
	inst.objects = make(map[string]*myConfigObject)
}

////////////////////////////////////////////////////////////////////////////////

type myConfigObject struct {
	clazz       string
	name        string
	countFields int
}

func (inst *myConfigObject) init(clazz, name string) {
	inst.clazz = clazz
	inst.name = name
	inst.countFields = 0
}

////////////////////////////////////////////////////////////////////////////////

type myConfigIndexBuilder struct {
	index myConfigIndex
}

func (inst *myConfigIndexBuilder) init(cfg store.Config) {
	inst.index.config = cfg
	inst.index.classes = make(map[string]*myConfigClass)
	src := cfg.Export()
	for k, v := range src {
		inst.put(k, v)
	}
}

func (inst *myConfigIndexBuilder) create() *myConfigIndex {
	dst := &myConfigIndex{
		config:  inst.index.config,
		classes: inst.index.classes,
	}
	return dst
}

func (inst *myConfigIndexBuilder) put(k, v string) {
	parts := strings.Split(k, ".")
	if len(parts) < 2 {
		return // ignore
	}
	clazzStr := parts[0] // the tag
	nameStr := parts[1]  // the name
	tab := inst.index.classes
	clazz := tab[clazzStr]
	if clazz == nil {
		clazz = &myConfigClass{}
		clazz.init(clazzStr)
		tab[clazzStr] = clazz
	}
	inst.putNameToIndex(clazz, nameStr)
}

func (inst *myConfigIndexBuilder) putNameToIndex(clazz *myConfigClass, name string) {
	tab := clazz.objects
	obj := tab[name]
	if obj == nil {
		obj = &myConfigObject{}
		obj.init(clazz.className, name)
		tab[name] = obj
	}
	obj.countFields++
}
