package u4go

import "github.com/starter-go/units"

type Demo struct {

	//starter:component

	_as func(units.Units) //starter:as(".")
}

// Units implements units.Units.
func (inst *Demo) Units(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:    "a",
		Enabled: true,
		Test:    inst.run,
	}

	r2 := &units.Registration{
		Name:    "b",
		Enabled: true,
		Test:    inst.run,
	}

	list = append(list, r1, r2)
	return list
}

func (inst *Demo) _impl() units.Units {
	return inst
}

func (inst *Demo) run() error {
	return nil
}
