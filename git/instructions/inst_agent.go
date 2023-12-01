package instructions

import (
	"github.com/bitwormhole/gitlib/git/repositories"
)

func findService(meta *repositories.Meta) repositories.Service {
	lib, err := repositories.GetLib(meta.Context)
	if err != nil {
		panic(err)
	}
	csm := lib.InstructionServiceManager()
	ser, err := csm.Find(meta.Name)
	if err != nil {
		panic(err)
	}
	return ser
}
