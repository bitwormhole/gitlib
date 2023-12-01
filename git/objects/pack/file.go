package pack

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git/objects"
	"github.com/starter-go/afs"
)

// File ...
type File struct {
	Context *objects.PackContext
	Type    FileType
	Path    afs.Path
}

// OpenReader ...
func (inst *File) OpenReader() (io.ReadSeekCloser, error) {
	ctx := inst.Context.Parent
	pool := ctx.Pool
	file := inst.Path
	digest := ctx.Digest
	if digest == nil {
		return nil, fmt.Errorf("digest is nil")
	}
	if pool == nil {
		return nil, fmt.Errorf("pool is nil")
	}
	if file == nil {
		return nil, fmt.Errorf("file path is nil")
	}
	size := digest.Size()
	if !size.IsValid() {
		return nil, fmt.Errorf("bad hash size:%v bit(s)", size.SizeInBits())
	}
	r, err := pool.OpenReader(file, &afs.Options{})
	if err != nil {
		return nil, err
	}
	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Clone ...
func (inst *File) Clone() *File {

	dst := &File{}
	src := inst

	dst.Context = src.Context
	dst.Path = src.Path
	dst.Type = src.Type

	return dst
}
