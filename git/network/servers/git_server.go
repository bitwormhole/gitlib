package servers

// ServerRegistration ...
type ServerRegistration struct {
	Server Server
}

// ServerRegistry ... [inject:".git-server-registry"]
type ServerRegistry interface {
	GetServerRegistration() *ServerRegistration
}

// Server ...
type Server interface {
	Accept(c *Context) bool
	Execute(c *Context) error
}

// MainServer ... [inject:"#git-main-server"]
type MainServer interface {
	Server
}
