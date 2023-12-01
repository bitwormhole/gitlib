package repositories

import (
	"io"

	"github.com/starter-go/afs"
)

// FileReaderPool ...
type FileReaderPool interface {
	io.Closer

	OpenReader(file afs.Path, op *afs.Options) (io.ReadSeekCloser, error)
}
