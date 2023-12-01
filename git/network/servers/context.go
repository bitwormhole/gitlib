package servers

import (
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/repositories"
)

// Context ...
type Context struct {
	Protocol string
	Method   string
	User     string
	Alias    string
	Service  string

	Connection pktline.Connection
	Layout     repositories.Layout
	Repository repositories.Repository
	Session    repositories.Session
}
