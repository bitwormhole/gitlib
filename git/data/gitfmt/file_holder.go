package gitfmt

import (
	"io"

	"github.com/starter-go/afs"
)

// FileHolder ...
type FileHolder struct {
	File afs.Path

	io.ReadSeekCloser
}
