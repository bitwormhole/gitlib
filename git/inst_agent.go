package git

import (
	"github.com/bitwormhole/gitlib/git/instructions"
	"github.com/bitwormhole/gitlib/git/store"
)

func findService(meta *instructions.Meta) instructions.Service {
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
