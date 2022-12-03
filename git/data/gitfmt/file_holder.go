package gitfmt

import (
	"io"

	"bitwormhole.com/starter/afs"
)

// FileHolder ...
type FileHolder struct {
	File afs.Path

	io.ReadSeekCloser
}
