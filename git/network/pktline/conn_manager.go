package pktline

import "fmt"

// ConnectorManager ... [inject:"#pktline-connector-manager"]
type ConnectorManager interface {
	Connector
}

////////////////////////////////////////////////////////////////////////////////

// NewConnectorManager 新建连接器管理器
func NewConnectorManager(crlist []*ConnectorRegistration) ConnectorManager {

	cm := &connectorManagerImpl{}
	dst := cm.connectors
	src := crlist

	for _, cr := range src {
		if cr == nil {
			continue
		}
		conn := cr.Connector
		if conn == nil {
			continue
		}
		dst = append(dst, conn)
	}

	cm.connectors = dst
	return cm
}

////////////////////////////////////////////////////////////////////////////////

type connectorManagerImpl struct {
	connectors []Connector
}

func (inst *connectorManagerImpl) _Impl() ConnectorManager {
	return inst
}

func (inst *connectorManagerImpl) Connect(p *ConnParams) (Connection, error) {
	for _, c := range inst.connectors {
		if c.Accept(p) {
			return c.Connect(p)
		}
	}
	url := p.URL
	return nil, fmt.Errorf("no connector support the url [%v]", url)
}

func (inst *connectorManagerImpl) Accept(p *ConnParams) bool {
	for _, c := range inst.connectors {
		if c.Accept(p) {
			return true
		}
	}
	return false
}
