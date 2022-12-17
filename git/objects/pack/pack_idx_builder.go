package pack

import (
	"fmt"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/starter/vlog"
)

// IdxBuilder ...
type IdxBuilder interface {
	AddItem(item *git.PackIndexItem)
}

////////////////////////////////////////////////////////////////////////////////

type idxBuilderContext struct {
	Builder     IdxBuilder
	Pool        afs.ReaderPool
	output      io.Writer
	compression git.Compression
	pack        Pack
	pos         int64
}

func (inst *idxBuilderContext) init(src afs.Path, dst io.Writer) {
	inst.output = dst
	inst.Builder = &idxBuilderImpl{context: inst}
	inst.Pool = &idxBuilderPool{context: inst, file: src}
	inst.pos = 4 * 3 // 4-bytes * (sign + version + count)
}

func (inst *idxBuilderContext) make() error {
	for {
		err := inst.readNextObject()
		if err != nil {
			return err
		}
	}
}

func (inst *idxBuilderContext) readNextObject() error {

	pool := inst.Pool
	item := &git.PackedObjectHeaderEx{
		OID:    nil,
		Offset: inst.pos,
	}

	h, r, err := inst.pack.OpenObjectReader(item, pool)
	if err != nil {
		return err
	}
	defer func() {
		r.Close()
	}()

	vlog.Info("in-pack-object  type:%v length:%v", h.Type, h.Size)
	io.ReadAll(r)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type idxBuilderImpl struct {
	context *idxBuilderContext
}

func (inst *idxBuilderImpl) _Impl() IdxBuilder {
	return inst
}

func (inst *idxBuilderImpl) AddItem(item *git.PackIndexItem) {
	return
}

////////////////////////////////////////////////////////////////////////////////

type idxBuilderPool struct {
	context      *idxBuilderContext
	file         afs.Path
	cachedReader io.ReadSeekCloser
	realCloser   io.Closer
}

func (inst *idxBuilderPool) _Impl() afs.ReaderPool {
	return inst
}

func (inst *idxBuilderPool) Clean() {}

func (inst *idxBuilderPool) Close() error {
	c := inst.realCloser
	inst.realCloser = nil
	if c != nil {
		return c.Close()
	}
	return nil
}

func (inst *idxBuilderPool) wrapReader(inner io.ReadSeekCloser) io.ReadSeekCloser {
	return &idxBuilderReaderProxy{
		context: inst.context,
		inner:   inner,
	}
}

func (inst *idxBuilderPool) getFilePath(file afs.Path) string {
	if file == nil {
		return ""
	}
	return file.GetPath()
}

func (inst *idxBuilderPool) OpenReader(file afs.Path, op *afs.Options) (io.ReadSeekCloser, error) {

	want := inst.getFilePath(inst.file)
	have := inst.getFilePath(file)
	if want != have {
		return nil, fmt.Errorf("bad file path, want:[%v] have:[%v]", want, have)
	}

	r := inst.cachedReader
	if r == nil {
		r2, err := file.GetIO().OpenSeekerR(op)
		if err != nil {
			return nil, err
		}
		inst.realCloser = r2
		r = inst.wrapReader(r2)
		inst.cachedReader = r
	}
	return r, nil
}

////////////////////////////////////////////////////////////////////////////////

type idxBuilderReaderProxy struct {
	context *idxBuilderContext
	inner   io.ReadSeekCloser
}

func (inst *idxBuilderReaderProxy) _Impl() io.ReadSeekCloser {
	return inst
}

func (inst *idxBuilderReaderProxy) Read(dst []byte) (int, error) {
	return inst.inner.Read(dst)
}

func (inst *idxBuilderReaderProxy) Seek(off int64, whence int) (int64, error) {
	return inst.inner.Seek(off, whence)
}

func (inst *idxBuilderReaderProxy) Close() error {
	pos, err := inst.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}
	inst.context.pos = pos
	return nil
}

////////////////////////////////////////////////////////////////////////////////
