package repository

import (
	"io"

	"github.com/bitwormhole/starter/lang"
)

// Session 访问仓库的会话
type Session interface {
	io.Closer
	GetView() View
	GetAttribute(name string) lang.Object
	SetAttribute(name string, attr lang.Object)
}

// SessionFactory 用来创建会话
type SessionFactory interface {
	Open(view View) Session
}
