package instructions

import (
	"github.com/bitwormhole/gitlib/git/store"
)

func findService(meta *store.Meta) store.Service {
	lib, err := store.GetLib(meta.Context)
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
