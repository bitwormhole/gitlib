package cases

import (
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/units"
)

// SystemContextTest ...
type SystemContextTest struct {

	//starter:component

	LibAgent repositories.LibAgent //starter:inject("#")

}

func (inst *SystemContextTest) _impl() units.Units {
	return inst
}

// Units ...
func (inst *SystemContextTest) Units(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:    "test-system-context",
		Enabled: true,
		Test:    inst.t,
	}
	list = append(list, r1)
	return list
}

func (inst *SystemContextTest) t() error {

	lib, err := inst.LibAgent.GetLib()
	if err != nil {
		return err
	}

	lib.Loader()

	return nil
}
