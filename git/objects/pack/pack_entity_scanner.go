package pack

import (
	"bufio"
	"hash"
	"hash/crc32"
	"strconv"
	"strings"

	"fmt"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects"
)

type packObjectOffsetScanner struct {
	context *objects.Context
	pid     git.PackID // 保存扫描所得的PID
}

func (inst *packObjectOffsetScanner) scan(file afs.Path) ([]*git.PackedObjectHeaderEx, error) {

	// 扫描简单对象，计算CRC32
	solist, err := inst.scanSimpleObjects(file)
	if err != nil {
		return nil, err
	}

	// 扫描复杂对象，计算SHA-1
	return inst.scanComplexObjects(file, solist)
}

func (inst *packObjectOffsetScanner) scanComplexObjects(file afs.Path, items []*git.PackedObjectHeaderEx) ([]*git.PackedObjectHeaderEx, error) {

	pid := inst.pid
	oc := inst.context
	pc := oc.NewPackContext(pid)

	pack1, err := NewComplexPack(&File{
		Context: pc,
		Path:    file,
		Type:    FileTypePack,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		oid, err := inst.computeComplexObjectSumID(item, pack1)
		if err != nil {
			return nil, err
		}
		item.OID = oid
		item.PID = pid
	}

	return items, nil
}

func (inst *packObjectOffsetScanner) computeComplexObjectSumID(item *git.PackedObjectHeaderEx, p ComplexPack) (git.ObjectID, error) {

	hx, src, err := p.OpenComplexObjectReader(item, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		src.Close()
	}()

	// make head
	sizeWant := hx.Size
	hBuilder := strings.Builder{}
	objType := hx.Type.ToObjectType()
	hBuilder.WriteString(objType.String())
	hBuilder.WriteRune(' ')
	hBuilder.WriteString(strconv.FormatInt(sizeWant, 10))
	hBuilder.WriteRune('.')
	head := []byte(hBuilder.String())
	head[len(head)-1] = 0

	// write head
	dst := inst.context.Digest.New()
	dst.Write(head)

	// write body (data)
	sizeHave, err := io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	// check type
	if objType == git.ObjectTypeBLOB {
	} else if objType == git.ObjectTypeTree {
	} else if objType == git.ObjectTypeCommit {
	} else if objType == git.ObjectTypeTag {
	} else {
		return nil, fmt.Errorf("bad complex-in-pack-object type: %v", objType.String())
	}

	// check size
	if sizeWant != sizeHave {
		return nil, fmt.Errorf("bad complex-in-pack-object size, want:%v have:%v", sizeWant, sizeHave)
	}

	sum := dst.Sum([]byte{})
	return git.CreateObjectID(sum)
}

func (inst *packObjectOffsetScanner) scanSimpleObjects(file afs.Path) ([]*git.PackedObjectHeaderEx, error) {

	r1, err := file.GetIO().OpenSeekerR(nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		r1.Close()
	}()

	// 读取头部
	cr := commonReader{}
	sign, e1 := cr.read4Bytes(r1)
	version, e2 := cr.readUInt32(r1)
	count, e3 := cr.readUInt32(r1)
	errlist := []error{e1, e2, e3}
	for _, err := range errlist {
		if err != nil {
			return nil, err
		}
	}

	// 校验头部
	err = inst.checkFileHead(sign[:], version)
	if err != nil {
		return nil, err
	}

	// 读取简单对象
	items := make([]*git.PackedObjectHeaderEx, 0)
	for i := uint32(0); i < count; i++ {
		item, err := inst.readNextObject(r1)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// 读取 pid
	pid, err := inst.readPackSumID(r1)
	if err != nil {
		return nil, err
	}
	inst.pid = pid

	return items, nil
}

func (inst *packObjectOffsetScanner) checkFileHead(sign []byte, ver uint32) error {
	s2 := string(sign)
	if s2 != "PACK" {
		return fmt.Errorf("bad .pack file sign:%v", s2)
	}
	if ver != 2 && ver != 3 {
		return fmt.Errorf("bad .pack file version:%v", ver)
	}
	return nil
}

// func (inst *packObjectOffsetScanner) wrapReader(r io.Reader) io.Reader {
// 	r = bufio.NewReaderSize(r, 1024*64)
// 	r = &the1byteReader{inner: r, context: inst.context}
// 	return r
// }

func (inst *packObjectOffsetScanner) readPackSumID(r io.ReadSeeker) (git.PackID, error) {

	digest := inst.context.Digest
	idsize := digest.Size()
	sum := make([]byte, idsize.SizeInBytes())
	sizeWant := len(sum)

	_, err := r.Seek(int64(0-sizeWant), io.SeekEnd)
	if err != nil {
		return nil, err
	}

	sizeHave, err := r.Read(sum)
	if err != nil {
		return nil, err
	}

	if sizeHave != sizeWant {
		return nil, fmt.Errorf("bad pack-id size")
	}

	return git.CreatePackID(sum)
}

func (inst *packObjectOffsetScanner) getPosition(r io.ReadSeeker) int64 {
	p, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0
	}
	return p
}

func (inst *packObjectOffsetScanner) readNextObject(r0 io.ReadSeeker) (*git.PackedObjectHeaderEx, error) {

	pos0 := inst.getPosition(r0)

	// prepare stream
	r1byte := &packObjectOffsetScanner1byteReader{}
	rCount := &packObjectOffsetScannerCountReader{}
	rCRC32 := &packObjectOffsetScannerCRC32Reader{}

	r1 := bufio.NewReader(r0)
	r := r1byte.wrap(r1)
	r = rCount.wrap(r)
	r = rCRC32.wrap(r)

	// read head
	ehr := entityHeaderReader{}
	item := &git.PackedObjectHeaderEx{}
	item.Offset = pos0
	hx, err := ehr.readHeader(item, r)
	if err != nil {
		return nil, err
	}

	// read body
	err = inst.unzip(hx, r)
	if err != nil {
		return nil, err
	}

	pos2 := pos0 + rCount.count
	pos3, err := r0.Seek(pos2, io.SeekStart)
	if err != nil {
		return nil, err
	}
	if pos2 != pos3 {
		return nil, fmt.Errorf("bad in-pack object position")
	}

	hx.Offset = pos0
	hx.CRC32 = rCRC32.Result()
	return hx, nil
}

func (inst *packObjectOffsetScanner) unzip(hx *git.PackedObjectHeaderEx, r io.Reader) error {

	// unzip
	comp := inst.context.Compression
	src, err := comp.NewReader(r)
	if err != nil {
		return err
	}

	dst := &theNopWriter{}
	size1 := hx.Size
	size2, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	if size1 != size2 {
		return fmt.Errorf("bad size")
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type packObjectOffsetScanner1byteReader struct {
	inner io.Reader
}

func (inst *packObjectOffsetScanner1byteReader) wrap(r io.Reader) io.Reader {
	inst.inner = r
	return inst
}

func (inst *packObjectOffsetScanner1byteReader) Read(p1 []byte) (int, error) {
	p2 := p1[0:1]
	return inst.inner.Read(p2)
}

////////////////////////////////////////////////////////////////////////////////

type packObjectOffsetScannerCRC32Reader struct {
	input io.Reader
	h32   hash.Hash32
}

func (inst *packObjectOffsetScannerCRC32Reader) wrap(r io.Reader) io.Reader {

	const n = crc32.IEEE
	t := crc32.MakeTable(n)
	h := crc32.New(t)

	inst.h32 = h
	inst.input = r
	return inst
}

func (inst *packObjectOffsetScannerCRC32Reader) Read(p []byte) (int, error) {
	n, err := inst.input.Read(p)
	if n > 0 {
		data := p[0:n]
		inst.h32.Write(data)
	}
	return n, err
}

func (inst *packObjectOffsetScannerCRC32Reader) Result() int64 {

	// sb := strings.Builder{}
	// sb.WriteString("crc32")

	h := inst.h32
	n := h.Sum32()

	// sb.WriteString("  ")
	// sb.WriteString(strconv.FormatUint(uint64(n), 16))
	// vlog.Warn("crc32.result: ", sb.String())

	return int64(n) // todo ...
}

////////////////////////////////////////////////////////////////////////////////

type packObjectOffsetScannerCountReader struct {
	input io.Reader
	count int64
}

func (inst *packObjectOffsetScannerCountReader) wrap(r io.Reader) io.Reader {
	inst.input = r
	inst.count = 0
	return inst
}

func (inst *packObjectOffsetScannerCountReader) Read(p []byte) (int, error) {
	n, err := inst.input.Read(p)
	if n > 0 {
		inst.count += int64(n)
	}
	return n, err
}

////////////////////////////////////////////////////////////////////////////////

type theNopWriter struct {
}

func (inst *theNopWriter) Write(p []byte) (int, error) {
	n := len(p)
	return n, nil
}

////////////////////////////////////////////////////////////////////////////////
