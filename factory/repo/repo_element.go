package repo

import "github.com/bitwormhole/starter/lang"

type Element interface {
	lang.Disposable
	Link() error
	Init() error
}

type ElementFactory interface {
	Make(ctx *ViewportContext) error
}
