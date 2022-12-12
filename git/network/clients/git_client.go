package clients

// ClientRegistration ...
type ClientRegistration struct {
	Client Client
}

// ClientRegistry ... [inject:".git-client-registry"]
type ClientRegistry interface {
	GetClientRegistration() *ClientRegistration
}

// Client ...
type Client interface {
	Accept(c *Context) bool
	Execute(c *Context) error
}

// MainClient ... [inject:"#git-main-client"]
type MainClient interface {
	Client
}
