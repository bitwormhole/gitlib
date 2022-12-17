package pack

import (
	"bufio"

	"fmt"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

type packObjectOffsetScanner struct {
	context *idxBuilderContext
}

func (inst *packObjectOffsetScanner) scan(file afs.Path) ([]*git.PackedObjectHeaderEx, error) {

	// inst.context = &idxBuilderContext{}

	r1, err := file.GetIO().OpenReader(nil)
	if err != nil {
		return nil, err
	}
	r := inst.wrapReader(r1)
	defer func() {
		r1.Close()
	}()

	cr := commonReader{}

	sign, e1 := cr.read4Bytes(r)
	version, e2 := cr.readUInt32(r)
	count, e3 := cr.readUInt32(r)
	errlist := []error{e1, e2, e3}
	for _, err := range errlist {
		if err != nil {
			return nil, err
		}
	}

	// vlog.Warn(fmt.Sprintf("pack sign:%v ver:%v cnt:%v", sign, version, count))
	err = inst.checkFileHead(sign[:], version)
	if err != nil {
		return nil, err
	}

	items := make([]*git.PackedObjectHeaderEx, 0)

	for i := uint32(0); i < count; i++ {
		item, err := inst.readNextObject(r)
		if err != nil {
			return nil, err
		}
		// item.Index = int64(i)
		// const f = "in-pack-item index:%v offset:%v length:%v"
		// vlog.Warn(fmt.Sprintf(f, item.Index, item.Offset, item.Length))
		items = append(items, item)
	}

	// ehr.readHeader(item, r)

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

func (inst *packObjectOffsetScanner) wrapReader(r io.Reader) io.Reader {
	r = bufio.NewReaderSize(r, 1024*64)
	r = &the1byteReader{inner: r, context: inst.context}
	return r
}

func (inst *packObjectOffsetScanner) readNextObject(r io.Reader) (*git.PackedObjectHeaderEx, error) {

	ehr := entityHeaderReader{}
	item := &git.PackedObjectHeaderEx{}
	item.Offset = inst.context.pos
	hx, err := ehr.readHeader(item, r)
	if err != nil {
		return nil, err
	}

	// vlog.Warn(" offset: %v", hx.Offset)

	err = inst.unzip(hx, r)
	if err != nil {
		return nil, err
	}

	// item.Size = hx.Size
	return hx, nil
}

func (inst *packObjectOffsetScanner) unzip(hx *git.PackedObjectHeaderEx, r1 io.Reader) error {

	comp := inst.context.compression
	src, err := comp.NewReader(r1)
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

type the1byteReader struct {
	context *idxBuilderContext
	inner   io.Reader
}

func (inst *the1byteReader) Read(p1 []byte) (int, error) {
	p2 := p1[0:1]
	n, err := inst.inner.Read(p2)
	if n > 0 {
		inst.context.pos += int64(n)
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
