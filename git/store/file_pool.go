package store

import (
	"io"

	"bitwormhole.com/starter/afs"
)

// FileReaderPool ...
type FileReaderPool interface {
	io.Closer

	OpenReader(file afs.Path, op *afs.Options) (io.ReadSeekCloser, error)
}
