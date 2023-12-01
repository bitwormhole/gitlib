package http4git

import (
	"strings"

	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/starter/markup"
)

// HTTPGitConnector ...
type HTTPGitConnector struct {
	markup.Component `class:"pktline-connector-registry"`
}

func (inst *HTTPGitConnector) _Impl() (pktline.ConnectorRegistry, pktline.Connector) {
	return inst, inst
}

// GetRegistration ...
func (inst *HTTPGitConnector) GetRegistration() *pktline.ConnectorRegistration {
	return &pktline.ConnectorRegistration{
		Connector: inst,
	}
}

// Connect ...
func (inst *HTTPGitConnector) Connect(p *pktline.ConnParams) (pktline.Connection, error) {
	cc := &connectionContext{}
	err := cc.init(p.URL, p.SecurityOnly)
	if err != nil {
		return nil, err
	}
	return cc.open(p)
}

// Accept ...
func (inst *HTTPGitConnector) Accept(p *pktline.ConnParams) bool {
	url := p.URL
	if strings.HasPrefix(url, "http://") {
		return true
	}
	if strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}
