package store

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/services"
)

// Lib ...
type Lib interface {
	// Context() context.Context

	FS() afs.FS

	RepositoryFinder() RepositoryFinder

	RepositoryLocator() RepositoryLocator

	ServiceManager() services.ServiceManager
}
