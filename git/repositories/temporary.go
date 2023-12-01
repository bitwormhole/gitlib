package repositories

import (
	"io"

	"github.com/starter-go/afs"
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
