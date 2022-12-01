package store

import (
	"io"

	"bitwormhole.com/starter/afs"
)

// TemporaryBuffer ...
type TemporaryBuffer interface {
	io.Closer
	io.Writer

	GetTemporaryFile() afs.Path

	GetFlushSize() int

	SetFlushSize(size int)

	SaveToFile(dst afs.Path) error
}
