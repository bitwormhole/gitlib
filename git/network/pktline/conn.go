package pktline

import (
	"io"
)

// 定义 git 服务名称
const (
	ServiceGitUpload = "git-upload-objects"
)

// ConnParams ...
type ConnParams struct {
	Method  string
	URL     string
	Service string
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

	Reader() Reader

	Writer() Writer

	// 创建新的附加连接
	NewConnection(service string) (Connection, error)
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
