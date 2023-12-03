package repositories

import "github.com/starter-go/afs"

// SystemSpace 是 SystemContext 的 facade
type SystemSpace interface {
	Lib() Lib
}

// UserSpace 是 UserContext 的 facade
type UserSpace interface {
	Home() afs.Path

	SystemSpace() SystemSpace
}
