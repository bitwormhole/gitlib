package lib

import (
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support/net/http4git"
	"github.com/bitwormhole/starter/markup"
)

// ConfigConnectors ...
type ConfigConnectors struct {
	markup.Component `class:"git-context-configurer"`

	Connectors []pktline.ConnectorRegistry `inject:".pktline-connector-registry"`
}

func (inst *ConfigConnectors) _Impl() store.ContextConfigurer {
	return inst
}

// Configure ...
func (inst *ConfigConnectors) Configure(c *store.Context) error {

	src := inst.Connectors
	dst := make([]pktline.Connector, 0)
	crlist := make([]*pktline.ConnectorRegistration, 0)

	for _, item := range src {
		r := item.GetRegistration()
		dst = append(dst, r.Connector)
		crlist = append(crlist, r)
	}

	c.Connectors = dst
	c.ConnectorManager = pktline.NewConnectorManager(crlist)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// HTTPGitConnectorReg ...
type HTTPGitConnectorReg struct {
	markup.Component `class:"pktline-connector-registry"`
	http4git.HTTPGitConnector
}
