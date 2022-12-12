package servers

import (
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
)

// Context ...
type Context struct {
	Protocol string
	Method   string
	User     string
	Alias    string
	Service  string

	Connection pktline.Connection
	Layout     store.RepositoryLayout
	Repository store.Repository
	Session    store.Session
}
