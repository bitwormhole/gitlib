package gitfmt

import (
	"io"

	"bitwormhole.com/starter/afs"
)

type FileHolder struct {
	File afs.Path

	io.ReadSeekCloser
}
