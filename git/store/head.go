package store

import "github.com/bitwormhole/gitlib/git"

// HEAD ...
type HEAD interface {
	NodeLocation

	GetValue(s Session) (git.ReferenceName, error)
}
