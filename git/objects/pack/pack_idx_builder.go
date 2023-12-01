package pack

import (
	"bufio"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"sort"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects"
	"github.com/starter-go/afs"
)

// IdxBuilder ...
type IdxBuilder interface {
	AddItem(item *git.PackedObjectHeaderEx) error
	Make(dst afs.Path) error
}

////////////////////////////////////////////////////////////////////////////////

// type idxBuilderContext struct {
// 	Builder     IdxBuilder
// 	Pool        afs.ReaderPool
// 	output      io.Writer
// 	compression git.Compression
// 	pack        Pack
// 	// pos         int64
// }

// func (inst *idxBuilderContext) init(src afs.Path, dst io.Writer) {
// 	// inst.output = dst
// 	// inst.Builder = &v2IdxBuilder{context: inst}
// 	// inst.Pool = &idxBuilderPool{context: inst, file: src}
// 	// inst.pos = 4 * 3 // 4-bytes * (sign + version + count)
// }

// func (inst *idxBuilderContext) make() error {
// 	for {
// 		err := inst.readNextObject()
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func (inst *idxBuilderContext) readNextObject() error {

// 	pool := inst.Pool
// 	item := &git.PackedObjectHeaderEx{
// 		OID:    nil,
// 		// Offset: inst.pos,
// 	}

// 	h, r, err := inst.pack.OpenObjectReader(item, pool)
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		r.Close()
// 	}()

// 	vlog.Info("in-pack-object  type:%v length:%v", h.Type, h.Size)
// 	io.ReadAll(r)

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////

type v2IdxBuilder struct {
	context        *objects.PackContext
	pack           ComplexPack
	fanout         git.PackIdxFanOut
	items          []*git.PackIndexItem
	offset64values []int64
	hash           hash.Hash
	// hCRC32         hash.Hash32
	bufferdOut *bufio.Writer
}

func (inst *v2IdxBuilder) _Impl() IdxBuilder {
	return inst
}

func (inst *v2IdxBuilder) newCRC32() hash.Hash32 {
	t := crc32.MakeTable(crc32.Koopman)
	return crc32.New(t)
}

func (inst *v2IdxBuilder) AddItem(item1 *git.PackedObjectHeaderEx) error {

	// hx, r, err := inst.pack.OpenComplexObjectReader(item, nil)
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	r.Close()
	// }()

	// head := inst.makeObjectHead(hx)
	// digest := inst.context.Parent.Digest
	// hash1 := digest.New()
	// hash2 := inst.newCRC32()
	// hash1.Write(head)
	// w := io.MultiWriter(hash1, hash2)
	// // w.Write(head)

	// cb, err := io.Copy(w, r)
	// if err != nil {
	// 	return err
	// }
	// if cb != hx.Size {
	// 	return fmt.Errorf("bad in-pack object size")
	// }

	// sum1 := hash1.Sum(nil)
	// sum32 := hash2.Sum32()
	// oid := git.DefaultIdentityFactory().Create(sum1[:])

	// item2 := &git.PackIndexItem{
	// 	Offset:     hx.Offset,
	// 	OID:        oid,
	// 	CRC32:      sum32,
	// 	Type:       hx.Type.ToObjectType(),
	// 	PackedType: hx.Type,
	// }
	// inst.items = append(inst.items, item2)
	// return nil

	item2 := &git.PackIndexItem{
		PID:    item1.PID,
		OID:    item1.OID,
		Offset: item1.Offset,
		CRC32:  uint32(item1.CRC32),
	}
	inst.items = append(inst.items, item2)
	return nil
}

func (inst *v2IdxBuilder) makeObjectHead(item *git.PackedObjectHeaderEx) []byte {
	t := item.Type.ToObjectType().String()
	size := item.Size
	str := fmt.Sprintf("%v %v.", t, size)
	data := []byte(str)
	data[len(data)-1] = 0
	return data
}

func (inst *v2IdxBuilder) Len() int {
	return len(inst.items)
}

func (inst *v2IdxBuilder) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	n := git.CompareObjectIDs(o1.OID, o2.OID)
	return (n > 0)
}

func (inst *v2IdxBuilder) Swap(i1, i2 int) {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	inst.items[i1] = o2
	inst.items[i2] = o1
}

func (inst *v2IdxBuilder) getByte0FromOID(oid git.ObjectID) byte {
	data := oid.Bytes()
	if data == nil {
		return 0
	}
	if len(data) < 1 {
		return 0
	}
	return data[0]
}

func (inst *v2IdxBuilder) makeFanout() error {
	dst := &git.PackIdxFanOut{}
	src := inst.items
	count := uint32(0)
	b1 := byte(0)
	for _, item := range src {
		oid := item.OID
		b0 := inst.getByte0FromOID(oid) // oid.GetByte(0)
		if b1 != b0 {
			i := int(b1)
			dst.Data[i] = count
			b1 = b0
		}
		count++
	}
	const size = 256
	value := dst.Data[0]
	for i := 0; i < size; i++ {
		v := dst.Data[i]
		if v > 0 {
			value = v
		} else {
			dst.Data[i] = value
		}
	}
	dst.Data[size-1] = count
	inst.fanout = *dst
	return nil
}

func (inst *v2IdxBuilder) writeToStream(dst io.Writer) error {

	steps := make([]func(dst io.Writer) error, 0)

	steps = append(steps, inst.writeHead) // magic + version
	steps = append(steps, inst.writeFanout)
	steps = append(steps, inst.writeIDs)
	steps = append(steps, inst.writeCRC32)
	steps = append(steps, inst.writeOffset32)
	steps = append(steps, inst.writeOffset64)
	steps = append(steps, inst.writePID)
	steps = append(steps, inst.writeSum)

	for _, step := range steps {
		err := step(dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *v2IdxBuilder) writeHead(dst io.Writer) error {
	// head & fanout
	const magic = MagicNumberIdxV2
	const version = 0x02
	numlist := []uint32{magic, version}
	return inst.writeUInt32array(numlist, dst)
}

func (inst *v2IdxBuilder) writeFanout(dst io.Writer) error {
	numlist := inst.fanout.Data[:]
	return inst.writeUInt32array(numlist, dst)
}

func (inst *v2IdxBuilder) writeIDs(dst io.Writer) error {
	for _, item := range inst.items {
		id := item.OID
		err := inst.writeHashID(id.HID(), dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *v2IdxBuilder) writeOffset32(dst io.Writer) error {
	inst.offset64values = nil
	o64values := inst.offset64values
	o64index := uint32(0)
	for _, item := range inst.items {
		offset := item.Offset
		value := uint32(0)
		if inst.isOffset64mode(offset) {
			// as offset64
			value = (0x80000000) | (o64index & 0x7fffffff)
			o64index++
			o64values = append(o64values, offset)
		} else {
			// as offset32
			value = uint32(offset)
		}
		err := inst.writeUInt32(value, dst)
		if err != nil {
			return err
		}
	}
	inst.offset64values = o64values
	return nil
}

func (inst *v2IdxBuilder) writeOffset64(dst io.Writer) error {
	return inst.writeInt64array(inst.offset64values, dst)
}

func (inst *v2IdxBuilder) writeCRC32(dst io.Writer) error {
	for _, item := range inst.items {
		crc := item.CRC32
		err := inst.writeUInt32(crc, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *v2IdxBuilder) writePID(dst io.Writer) error {
	pid := inst.pack.GetPackID()
	return inst.writeHashID(pid.HID(), dst)
}

func (inst *v2IdxBuilder) writeSum(dst io.Writer) error {
	data, err := inst.computeSum()
	if err != nil {
		return err
	}
	return inst.writeBytes(data, dst)
}

func (inst *v2IdxBuilder) writeUInt32(n uint32, dst io.Writer) error {
	cw := commonWriter{}
	return cw.writeUInt32(n, dst)
}

func (inst *v2IdxBuilder) writeInt64(n int64, dst io.Writer) error {
	cw := commonWriter{}
	return cw.writeInt64(n, dst)
}

func (inst *v2IdxBuilder) writeUInt32array(list []uint32, dst io.Writer) error {
	for _, n := range list {
		err := inst.writeUInt32(n, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *v2IdxBuilder) writeInt64array(list []int64, dst io.Writer) error {
	for _, n := range list {
		err := inst.writeInt64(n, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *v2IdxBuilder) writeHashID(id git.HashID, dst io.Writer) error {
	if id == "" {
		return fmt.Errorf("param is nil")
	}
	data := id.Bytes()
	return inst.writeBytes(data, dst)
}

func (inst *v2IdxBuilder) isOffset64mode(offset int64) bool {
	const max = 0x7fffffff
	return offset > max
}

func (inst *v2IdxBuilder) computeSum() ([]byte, error) {
	h := inst.hash
	if h == nil {
		return nil, fmt.Errorf("no hash")
	}
	sum := h.Sum([]byte{})
	return sum, nil
}

func (inst *v2IdxBuilder) writeBytes(p []byte, dst io.Writer) error {
	cw := commonWriter{}
	return cw.writeBytes(p, dst)
}

func (inst *v2IdxBuilder) writeToFile(dst afs.Path) error {

	dir := dst.GetParent()
	filename := dst.GetName()
	tmp := dir.GetChild(filename + ".tmp~")

	opt := afs.Todo().Create(true).Mkdirs(true)
	tmp.MakeParents(nil)
	w, err := tmp.GetIO().OpenWriter(opt)
	if err != nil {
		return err
	}
	defer func() {
		if w != nil {
			w.Close()
		}
		if tmp.Exists() {
			tmp.Delete()
		}
	}()

	w2 := inst.wrapWriter(w)
	err = inst.writeToStream(w2)
	if err != nil {
		return err
	}

	inst.bufferdOut.Flush()
	w.Close()
	w2 = nil
	w = nil

	return tmp.MoveTo(dst, nil)
}

func (inst *v2IdxBuilder) wrapWriter(dst io.Writer) io.Writer {
	w1 := inst.context.Parent.Digest.New()
	w2 := bufio.NewWriter(dst)
	w3 := io.MultiWriter(w1, w2)
	inst.hash = w1
	inst.bufferdOut = w2
	return w3
}

func (inst *v2IdxBuilder) Make(dst afs.Path) error {

	sort.Sort(inst)

	err := inst.makeFanout()
	if err != nil {
		return err
	}

	return inst.writeToFile(dst)
}

////////////////////////////////////////////////////////////////////////////////

// type idxBuilderPool struct {
// 	context      *idxBuilderContext
// 	file         afs.Path
// 	cachedReader io.ReadSeekCloser
// 	realCloser   io.Closer
// }

// func (inst *idxBuilderPool) _Impl() afs.ReaderPool {
// 	return inst
// }

// func (inst *idxBuilderPool) Clean() {}

// func (inst *idxBuilderPool) Close() error {
// 	c := inst.realCloser
// 	inst.realCloser = nil
// 	if c != nil {
// 		return c.Close()
// 	}
// 	return nil
// }

// // func (inst *idxBuilderPool) wrapReader(inner io.ReadSeekCloser) io.ReadSeekCloser {
// // 	return &idxBuilderReaderProxy{
// // 		context: inst.context,
// // 		inner:   inner,
// // 	}
// // }

// func (inst *idxBuilderPool) getFilePath(file afs.Path) string {
// 	if file == nil {
// 		return ""
// 	}
// 	return file.GetPath()
// }

// func (inst *idxBuilderPool) OpenReader(file afs.Path, op *afs.Options) (io.ReadSeekCloser, error) {

// 	want := inst.getFilePath(inst.file)
// 	have := inst.getFilePath(file)
// 	if want != have {
// 		return nil, fmt.Errorf("bad file path, want:[%v] have:[%v]", want, have)
// 	}

// 	r := inst.cachedReader
// 	if r == nil {
// 		r2, err := file.GetIO().OpenSeekerR(op)
// 		if err != nil {
// 			return nil, err
// 		}
// 		inst.realCloser = r2
// 		r = inst.wrapReader(r2)
// 		inst.cachedReader = r
// 	}
// 	return r, nil
// }

////////////////////////////////////////////////////////////////////////////////

// type idxBuilderReaderProxy struct {
// 	context *idxBuilderContext
// 	inner   io.ReadSeekCloser
// }

// func (inst *idxBuilderReaderProxy) _Impl() io.ReadSeekCloser {
// 	return inst
// }

// func (inst *idxBuilderReaderProxy) Read(dst []byte) (int, error) {
// 	return inst.inner.Read(dst)
// }

// func (inst *idxBuilderReaderProxy) Seek(off int64, whence int) (int64, error) {
// 	return inst.inner.Seek(off, whence)
// }

// func (inst *idxBuilderReaderProxy) Close() error {
// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////
