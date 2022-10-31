package gitconfig

import (
	"strings"

	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

// Property ...
type Property struct {
	properties store.Config
	key        string
}

func (inst *Property) init(cfg store.Config, template KeyTemplate, id string) {
	temp := string(template)
	key := temp
	const np = NamePlaceholder
	if strings.Contains(temp, np) {
		key = strings.Replace(temp, np, id, 1)
	}
	inst.properties = cfg
	inst.key = key
}

func (inst *Property) setProperty(value string) {
	key := inst.key
	inst.properties.SetProperty(key, value)
}

func (inst *Property) getProperty() (string, error) {
	const empty = ""
	key := inst.key
	value := inst.properties.GetProperty(key)
	if value == empty {
		return empty, nil
	}
	return value, nil
}

////////////////////////////////////////////////////////////////////////////////

// BooleanProperty ...
type BooleanProperty interface {
	Get() bool
	Set(value bool)
}

////////////////////////////////////////////////////////////////////////////////

// IntProperty ...
type IntProperty interface {
	Get() int
	Set(value int)
}

////////////////////////////////////////////////////////////////////////////////

// StringProperty ...
type StringProperty interface {
	Get() string
	Set(value string)
}

////////////////////////////////////////////////////////////////////////////////

// FloatProperty ...
type FloatProperty interface {
	Get() float64
	Set(value float64)
}

////////////////////////////////////////////////////////////////////////////////

// PropertyAccess ...
type PropertyAccess interface {
	ForString(template KeyTemplate, id string, def string) StringProperty
	ForInt(template KeyTemplate, id string, def int) IntProperty
	ForBool(template KeyTemplate, id string, def bool) BooleanProperty
	ForFloat(template KeyTemplate, id string, def float64) FloatProperty
}
