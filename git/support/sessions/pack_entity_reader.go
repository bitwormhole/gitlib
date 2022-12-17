package sessions

import (
	"bytes"
	"fmt"
	"io"

	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

type packEntityContext struct {
	session store.Session
}

////////////////////////////////////////////////////////////////////////////////

type packEntitySource interface {
	init() error
	test() error
	writeTo(w io.Writer) (int64, error)
	// openFragment(offset, size int64) (io.ReadCloser, error)
}

////////////////////////////////////////////////////////////////////////////////

type packEntityDeltaSource struct {
	context *packEntityContext
	parent  packEntitySource
	pack    pack.Pack
	// item       git.PackIndexItem
	head       git.PackedObjectHeaderEx
	sizeBefore int64
	sizeAfter  int64
}

func (inst *packEntityDeltaSource) _Impl() packEntitySource {
	return inst
}

func (inst *packEntityDeltaSource) init() error {
	return nil
}

func (inst *packEntityDeltaSource) test() error {

	return nil
}

func (inst *packEntityDeltaSource) writeTo(dst io.Writer) (int64, error) {

	buffer := &bytes.Buffer{}
	size1, err := inst.parent.writeTo(buffer)
	if err != nil && err != io.EOF {
		return 0, err
	}
	baseData := buffer.Bytes()

	src, err := inst.openMySource()
	if err != nil {
		return 0, err
	}
	defer func() {
		src.Close()
	}()

	// for debug
	disize10000 := 0
	list := make([]*DeltaInstruction, 0)

	count := int64(0)
	diReader := deltaInstructionReader{}
	for {
		di, err := diReader.Read(src)
		if di != nil {
			list = append(list, di)
			n, err1 := inst.apply(di, baseData, dst)
			if n > 0 {
				count += int64(n)
			}
			err = err1

			if di.Size == 0x10000 {
				disize10000 += int(di.Size)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
	}
	size2 := count

	if disize10000 > 0 {
		vlog.Debug("disize10000 = ", disize10000)
	}

	// check input size
	if inst.sizeBefore != size1 {
		return 0, fmt.Errorf("bad before size, have:%v want:%v", size1, inst.sizeBefore)
	}
	// check output size
	if inst.sizeAfter != size2 {
		return 0, fmt.Errorf("bad after size, have:%v want:%v", size2, inst.sizeAfter)
	}

	return size2, nil
}

func (inst *packEntityDeltaSource) apply(di *DeltaInstruction, basedata []byte, dst io.Writer) (int, error) {
	code := di.Code
	data := di.Data
	if code == InstructionCopyFromBase {
		p1 := di.Offset
		p2 := p1 + di.Size
		p3 := len(basedata)
		if p1 <= p2 && p2 <= uint32(p3) {
			data = basedata[p1:p2]
		} else {
			return 0, fmt.Errorf("out of range")
		}
	} else if code == InstructionAddNewData {
		data = di.Data
	} else {
		return 0, fmt.Errorf("bad instruction")
	}
	n, err := dst.Write(data)
	if n != len(data) {
		return 0, fmt.Errorf("bad size to write")
	} else if err != nil {
		return 0, err
	}
	return n, nil
}

func (inst *packEntityDeltaSource) openMySource() (io.ReadCloser, error) {
	item := &inst.head
	hx, in, err := inst.pack.OpenObjectReader(item, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if in != nil {
			in.Close()
		}
	}()

	vlog.Debug("delta_data_size:", hx.Size)
	// data, err := io.ReadAll(in)
	// if err != nil {
	// 	return nil, err
	// }
	// vlog.Debug("follow_data_size:", len(data))

	// 读取两个 vlq : len(source), len(dest)
	the2size := [2]int64{}
	count := len(the2size)
	for i := 0; i < count; i++ {
		n, err := pack.ReadSimple7bitsInt(in)
		if err != nil {
			return nil, err
		}
		the2size[i] = n
	}

	inst.sizeBefore = the2size[0]
	inst.sizeAfter = the2size[1]
	in2 := in
	in = nil
	return in2, err
}

func (inst *packEntityDeltaSource) openFragment(offset, size int64) (io.ReadCloser, error) {

	reader := &packEntityDeltaReader{
		source: inst,
		offset: offset,
		size:   size,
	}

	return reader, nil
}

////////////////////////////////////////////////////////////////////////////////

type packEntityDeltaReader struct {
	source *packEntityDeltaSource
	offset int64
	size   int64
}

func (inst *packEntityDeltaReader) Read(b []byte) (int, error) {
	return 0, nil
}

func (inst *packEntityDeltaReader) Close() error {
	return nil

}

////////////////////////////////////////////////////////////////////////////////

type packEntityBaseSource struct {
	context *packEntityContext
	pack    pack.Pack
	// item    git.PackIndexItem
	head git.PackedObjectHeaderEx
}

func (inst *packEntityBaseSource) _Impl() packEntitySource {
	return inst
}

func (inst *packEntityBaseSource) init() error {
	return nil
}

func (inst *packEntityBaseSource) test() error {
	return nil
}

func (inst *packEntityBaseSource) writeTo(dst io.Writer) (int64, error) {
	item := &inst.head
	_, src, err := inst.pack.OpenObjectReader(item, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		if src != nil {
			src.Close()
		}
	}()
	return io.Copy(dst, src)
}

////////////////////////////////////////////////////////////////////////////////

type packEntityBaseReader struct {
	source *packEntityBaseSource
	offset int64
	size   int64
	in     io.Reader
	cl     io.Closer
}

func (inst *packEntityBaseReader) Read(b []byte) (int, error) {
	return 0, nil
}

func (inst *packEntityBaseReader) Close() error {
	c := inst.cl
	inst.cl = nil
	if c == nil {
		return nil
	}
	return c.Close()
}

////////////////////////////////////////////////////////////////////////////////

type packEntityInfo struct {
	context *packEntityContext
	pid     git.PackID
	oid     git.ObjectID
	idx     pack.Idx
	pack    pack.Pack
	head    *git.PackedObjectHeaderEx
	// item    *git.PackIndexItem
}

func (inst *packEntityInfo) isBaseObject() bool {
	h := inst.head
	if h == nil {
		return false
	}
	t := h.Type
	return (t != git.PackedDeltaOFS && t != git.PackedDeltaRef)
}

func (inst *packEntityInfo) hasParent() bool {
	h := inst.head
	if h == nil {
		return false
	}
	t := h.Type
	return (t == git.PackedDeltaOFS || t == git.PackedDeltaRef)
}

func (inst *packEntityInfo) loadParent() (*packEntityInfo, error) {
	h := inst.head
	if h == nil {
		return nil, fmt.Errorf("no head for this in-pack object")
	}
	t := h.Type
	if t == git.PackedDeltaOFS {
		return inst.loadParentForDeltaOFS()
	} else if t == git.PackedDeltaRef {
		return inst.loadParentForDeltaRef()
	}
	return nil, fmt.Errorf("this is a in-pack base object")
}

func (inst *packEntityInfo) loadParentForDeltaOFS() (*packEntityInfo, error) {
	t := inst.head.Type
	if t != git.PackedDeltaOFS {
		return nil, fmt.Errorf("this is not a ofs_delta object")
	}
	parentOffset := inst.head.GetDeltaParentOffset()
	parent := &packEntityInfo{
		context: inst.context,
		pid:     inst.pid,
		oid:     nil,
		pack:    inst.pack,
		idx:     inst.idx,
		head:    nil,
	}
	item := &git.PackedObjectHeaderEx{
		Offset: parentOffset,
	}
	hx, err := parent.pack.ReadObjectHeader(item, nil)
	if err != nil {
		return nil, err
	}
	parent.head = hx
	// parent.item = item
	return parent, nil
}

func (inst *packEntityInfo) loadParentForDeltaRef() (*packEntityInfo, error) {
	return nil, fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////

type packEntityReader struct {
	chain packEntitySource
}

func (inst *packEntityReader) _Impl() io.ReadCloser {
	return inst
}

func (inst *packEntityReader) Read(b []byte) (int, error) {
	return 0, fmt.Errorf("this stream is closed")
}

func (inst *packEntityReader) Close() error {

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type packEntityReaderBuilder struct {
	session store.Session
	wantPID git.PackID
	wantOID git.ObjectID
	pc      PackCache

	// pack    pack.Pack
	// want *git.PackIndexItem
}

func (inst *packEntityReaderBuilder) open() (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	leaf, err := inst.loadLeafEntity()
	if err != nil {
		return nil, nil, err
	}
	if leaf.isBaseObject() {
		return inst.openAsSimpleObject(leaf)
	}
	return inst.openAsDeltaChainObject(leaf)
}

func (inst *packEntityReaderBuilder) openAsSimpleObject(info *packEntityInfo) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	item := info.head
	return info.pack.OpenObjectReader(item, nil)
}

func (inst *packEntityReaderBuilder) openAsDeltaChainObject(leaf *packEntityInfo) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {

	list, err := inst.loadEntityList(leaf)
	if err != nil {
		return nil, nil, err
	}

	chain, err := inst.makeSourceChain(list)
	if err != nil {
		return nil, nil, err
	}

	buffer := &bytes.Buffer{}
	size, err := chain.writeTo(buffer)
	if err != nil {
		return nil, nil, err
	}

	hx, err := inst.makeMainHeader(list, size)
	if err != nil {
		return nil, nil, err
	}

	// vlog.Debug(n)
	// reader := &packEntityReader{
	// 	chain: chain,
	// }

	reader := io.NopCloser(buffer)
	return hx, reader, nil
}

func (inst *packEntityReaderBuilder) makeSourceChain(list []*packEntityInfo) (packEntitySource, error) {
	var chain packEntitySource = nil
	for i := len(list) - 1; i >= 0; i-- {
		info := list[i]
		if info.isBaseObject() {
			src := &packEntityBaseSource{
				context: info.context,
				pack:    info.pack,
				// item:    *info.item,
				head: *info.head,
			}
			chain = src
		} else {
			src := &packEntityDeltaSource{
				parent:  chain,
				context: info.context,
				pack:    info.pack,
				// item:    *info.item,
				head: *info.head,
			}
			chain = src
		}
	}
	if chain == nil {
		return nil, fmt.Errorf("no node for chain")
	}

	err := chain.init()
	if err != nil {
		return nil, err
	}

	err = chain.test()
	if err != nil {
		return nil, err
	}

	return chain, nil
}

func (inst *packEntityReaderBuilder) makeMainHeader(list []*packEntityInfo, fianlSize int64) (*git.PackedObjectHeaderEx, error) {
	result := &git.PackedObjectHeaderEx{
		OID:    nil,
		PID:    nil,
		Offset: 0,
	}
	for _, info := range list {
		hx := info.head
		if hx.Type == git.PackedDeltaOFS || hx.Type == git.PackedDeltaRef {
			// skip
		} else {
			result.Type = hx.Type
		}
	}
	result.Size = fianlSize
	return result, nil
}

func (inst *packEntityReaderBuilder) loadLeafEntity() (*packEntityInfo, error) {

	oid := inst.wantOID
	pid := inst.wantPID
	session := inst.session
	if oid == nil {
		return nil, fmt.Errorf("wanted oid is nil")
	}
	q := &PackQuery{
		Session: session,
		PID:     pid,
		OID:     oid,
	}
	ok := inst.pc.Query(q)
	if !ok {
		err := q.Error
		if err == nil {
			err = fmt.Errorf("cannot find wanted object")
		}
		return nil, err
	}
	pid = q.PID

	pack := q.ResultHolder.pack
	item := q.ResultItem
	head := pack.IndexToHeader(item)
	hx, err := pack.ReadObjectHeader(head, nil)
	if err != nil {
		return nil, err
	}

	ctx := &packEntityContext{
		session: session,
	}

	leaf := &packEntityInfo{
		context: ctx,
		pid:     pid,
		oid:     oid,
		pack:    q.ResultHolder.pack,
		idx:     q.ResultHolder.idx,
		// item:    item,
		head: hx,
	}
	return leaf, nil
}

func (inst *packEntityReaderBuilder) loadEntityList(leaf *packEntityInfo) ([]*packEntityInfo, error) {
	list := make([]*packEntityInfo, 0)
	p := leaf
	for p != nil {
		list = append(list, p)
		if !p.hasParent() {
			break
		}
		parent, err := p.loadParent()
		if err != nil {
			return nil, err
		}
		p = parent
	}
	return list, nil
}

////////////////////////////////////////////////////////////////////////////////
