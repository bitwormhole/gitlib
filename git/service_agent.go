package git

import (
	"github.com/bitwormhole/gitlib/git/services"
	"github.com/bitwormhole/gitlib/git/store"
)

func findServiceForCommand(cmd *services.Command) services.Service {
	lib := store.GetLib(cmd.Context)
	csm := lib.ServiceManager()
	ser, err := csm.Find(cmd.Name)
	if err != nil {
		panic(err)
	}
	return ser
}
