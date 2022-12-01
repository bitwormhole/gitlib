package pack

import (
	"fmt"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

// File ...
type File struct {
	Digest git.Digest
	Pool   afs.ReaderPool
	Path   afs.Path
	Type   FileType
}

// OpenReader ...
func (inst *File) OpenReader() (io.ReadSeekCloser, error) {
	pool := inst.Pool
	file := inst.Path
	digest := inst.Digest
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

	dst.Path = src.Path
	dst.Type = src.Type
	dst.Pool = src.Pool
	dst.Digest = src.Digest

	return dst
}
