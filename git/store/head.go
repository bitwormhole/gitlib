package store

import "github.com/bitwormhole/gitlib/git/data/dxo"

// HEAD ...
type HEAD interface {
	NodeLocation

	GetValue(s Session) (dxo.ReferenceName, error)
}
