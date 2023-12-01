package repositories

import "github.com/bitwormhole/gitlib/git"

// HEAD ...
type HEAD interface {
	Node

	GetValue(s Session) (git.ReferenceName, error)
}
