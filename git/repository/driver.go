package repository

import "github.com/bitwormhole/starter/lang"

// Driver 仓库驱动器
type Driver interface {
	Supports(uri lang.URI) bool
	Open(uri lang.URI) (View, error)
}
