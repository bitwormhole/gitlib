package sessions

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type packEntityReader struct {

	// inputRaw io.ReadSeekCloser
	// compre   git.Compression
	// entityType   git.PackedObjectType
	// entityLength int64

	all  []*git.PackedObjectHeaderEx
	todo []*git.PackedObjectHeaderEx

	pack    pack.Pack
	closed  bool
	current *git.PackedObjectHeaderEx
	source  io.ReadCloser
}

func (inst *packEntityReader) _Impl() io.ReadCloser {
	return inst
}

func (inst *packEntityReader) Read(b []byte) (int, error) {
	for !inst.closed {
		n, _ := inst.tryRead(b)
		if n > 0 {
			return n, nil
		}
		err := inst.loadNextBlock()
		if err != nil {
			return 0, err
		}
	}
	return 0, fmt.Errorf("this stream is closed")
}

func (inst *packEntityReader) tryRead(b []byte) (int, error) {
	in := inst.source
	if in == nil {
		return 0, fmt.Errorf("stream is nil")
	}
	return in.Read(b)
}

func (inst *packEntityReader) loadNextBlock() error {
	var next *git.PackedObjectHeaderEx = nil
	list := inst.todo
	for i, item := range list {
		if item == nil {
			continue
		} else {
			inst.todo = list[i+1:]
			next = item
			break
		}
	}
	if next == nil {
		return io.EOF
	}

	// prepare index item
	item := &git.PackIndexItem{}
	item.Offset = next.Offset
	item.PID = next.PID
	item.OID = next.OID
	item.Type = next.Type.ToObjectType()
	item.PackedType = next.Type

	// open
	hx, in, err := inst.pack.OpenObjectReader(item, nil)
	if err != nil {
		return err
	}

	inst.current = hx
	inst.source = in
	return nil
}

func (inst *packEntityReader) Close() error {
	clist := make([]io.Closer, 0)
	clist = append(clist, inst.source)
	for _, c := range clist {
		if c == nil {
			continue
		}
		err := c.Close()
		if err != nil {
			vlog.Warn(err)
		}
	}
	inst.source = nil
	inst.all = nil
	inst.todo = nil
	inst.current = nil
	inst.closed = true
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type packEntityReaderBuilder struct {

	// file         afs.Path // the pack-*.pack file
	// input        io.ReadSeekCloser
	// index        *git.PackIndexItem
	// entityType   git.PackedObjectType
	// entityLength int64

	session store.Session
	pack    pack.Pack
	want    *git.PackIndexItem
}

func (inst *packEntityReaderBuilder) open() (*git.PackedObjectHeaderEx, io.ReadCloser, error) {

	blocks, err := inst.loadBlockList(inst.pack)
	if err != nil {
		return nil, nil, err
	}

	hx, err := inst.makeMainHeader(blocks)
	if err != nil {
		return nil, nil, err
	}

	reader := &packEntityReader{}
	reader.pack = inst.pack
	reader.all = blocks
	reader.todo = blocks

	return hx, reader, nil
}

func (inst *packEntityReaderBuilder) makeMainHeader(blocks []*git.PackedObjectHeaderEx) (*git.PackedObjectHeaderEx, error) {
	result := &git.PackedObjectHeaderEx{
		OID:    inst.want.OID,
		PID:    inst.want.PID,
		Offset: inst.want.Offset,
	}
	for _, hx := range blocks {
		if hx.Type == git.PackedDeltaOFS || hx.Type == git.PackedDeltaRef {
			// skip
		} else {
			result.Type = hx.Type
		}
		result.Size += hx.Size
	}
	return result, nil
}

func (inst *packEntityReaderBuilder) loadBlockList(p pack.Pack) ([]*git.PackedObjectHeaderEx, error) {

	index := inst.want
	want := &git.PackedObjectHeaderEx{}
	want.OID = index.OID
	want.PID = index.PID
	want.Offset = index.Offset

	list := make([]*git.PackedObjectHeaderEx, 0)

	for {
		have, err := inst.loadBlock(want, p)
		if err != nil {
			return nil, err
		}
		list = append(list, have)
		if have.Type == git.PackedDeltaOFS {
			want, err = inst.prepareWantBlockForDeltaOFS(have)
		} else if have.Type == git.PackedDeltaRef {
			want, err = inst.prepareWantBlockForDeltaRef(have)
		} else {
			break // ending
		}
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (inst *packEntityReaderBuilder) prepareWantBlockForDeltaOFS(have *git.PackedObjectHeaderEx) (*git.PackedObjectHeaderEx, error) {
	delta := have.DeltaOffset
	offset := have.Offset
	if offset < 1 || delta < 1 || delta >= offset {
		return nil, fmt.Errorf("bad offset value, offset:%v delta:%v", offset, delta)
	}
	want := &git.PackedObjectHeaderEx{}
	want.OID = have.OID
	want.PID = have.PID
	want.Offset = have.GetDeltaParentOffset()
	return want, nil
}

func (inst *packEntityReaderBuilder) prepareWantBlockForDeltaRef(d *git.PackedObjectHeaderEx) (*git.PackedObjectHeaderEx, error) {

	return nil, fmt.Errorf("no impl: prepareWantBlockForDeltaRef")
}

func (inst *packEntityReaderBuilder) loadBlock(want *git.PackedObjectHeaderEx, p pack.Pack) (*git.PackedObjectHeaderEx, error) {

	// make want-index
	wantIndex := &git.PackIndexItem{
		PID:    p.GetPackID(),
		Offset: want.Offset,
	}

	// read head
	hx, err := p.ReadObjectHeader(wantIndex, nil)
	if err != nil {
		return nil, err
	}
	hx.OID = want.OID
	hx.PID = want.PID
	hx.Offset = want.Offset

	return hx, nil
}

////////////////////////////////////////////////////////////////////////////////
