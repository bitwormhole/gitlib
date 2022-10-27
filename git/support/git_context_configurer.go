package support

import (
	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support/config"
	"github.com/bitwormhole/gitlib/git/support/others"
)

////////////////////////////////////////////////////////////////////////////////

// BaseContextConfigurer ...
type BaseContextConfigurer struct{}

func (inst *BaseContextConfigurer) _Impl() store.ContextConfigurer {
	return inst
}

// Configure ...
func (inst *BaseContextConfigurer) Configure(c *store.Context) error {

	c.FS = files.FS()

	c.Locator = &others.RepositoryLocatorImpl{}

	c.Finder = &others.RepositoryFinderImpl{}

	c.ConfigChainFactory = &config.ChainFactory{}

	c.ProfileFactory = &others.RepositoryProfileFactoryImpl{}

	c.Lib = &others.LibImpl{Context: c}

	c.ServiceManager = &others.CommandServiceManagerImpl{Context: c}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
