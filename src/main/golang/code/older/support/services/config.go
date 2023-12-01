package services

import "github.com/bitwormhole/gitlib/git/store"

// Configurer ...
type Configurer struct{}

func (inst *Configurer) _Impl() store.ContextConfigurer {
	return inst
}

// Configure ...
func (inst *Configurer) Configure(c *store.Context) error {

	list := c.Services

	list = append(list, &GitAddService{})
	list = append(list, &GitCommitService{})
	list = append(list, &GitInitService{})
	list = append(list, &GitPushService{})
	list = append(list, &GitStatusService{})

	c.Services = list
	return nil
}
