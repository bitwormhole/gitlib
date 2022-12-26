package objects

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

// Context ... 是 objects 的相关组件专用的上下文
type Context struct {
	Compression git.Compression
	Digest      git.Digest
	Pool        afs.ReaderPool

	AllObjects    Set
	PackedObjects Set
	SparseObjects Set
}

// PackContext ... 表示一个具体包的上下文
type PackContext struct {
	Parent *Context
	PID    git.PackID
	Self   Set
}

////////////////////////////////////////////////////////////////////////////////

// NewPackContext ...
func (inst *Context) NewPackContext(pid git.PackID) *PackContext {
	pc := &PackContext{
		Parent: inst,
		PID:    pid,
	}
	return pc
}
