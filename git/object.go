package git

import (
	"io"
)

// Object 接口是对一个git对象的抽象
type Object interface {
	GetType() string
	GetLength() int64
	GetID() ObjectID
	Read() (io.ReadCloser, error)
}
