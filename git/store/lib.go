package store

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/services"
)

// Lib ...
type Lib interface {
	FS() afs.FS

	RepositoryLoader() RepositoryLoader

	RepositoryFinder() RepositoryFinder

	RepositoryLocator() RepositoryLocator

	ServiceManager() services.ServiceManager
}
