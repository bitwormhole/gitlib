package pktline

import (
	"io"
)

// ConnParams ...
type ConnParams struct {
	Method       string
	URL          string
	Service      string
	ContentType  string // the request content-type
	SecurityOnly bool
}

// ConnectionGroup ...
type ConnectionGroup interface {
	GetAttribute(name string) any

	SetAttribute(name string, value any)
}

// Connection ...
type Connection interface {
	io.Closer

	GetGroup() ConnectionGroup

	GetAttribute(name string) any

	SetAttribute(name string, value any)

	GetParams() *ConnParams

	GetService() string

	// @return (reader,contentType,error)
	OpenReader() (ReaderCloser, string, error)

	OpenWriter(contentType string) (WriterCloser, error)

	// 创建新的附加连接
	NewConnection(p *ConnParams) (Connection, error)
}

// Connector ...
// [inject:".pktline-connector"]
type Connector interface {
	Connect(p *ConnParams) (Connection, error)

	Accept(p *ConnParams) bool
}

// ConnectorRegistration ...
type ConnectorRegistration struct {
	Connector Connector
}

// ConnectorRegistry 。。。[inject:".pktline-connector-registry"]
type ConnectorRegistry interface {
	GetRegistration() *ConnectorRegistration
}
