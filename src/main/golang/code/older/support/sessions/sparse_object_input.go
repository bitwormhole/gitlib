package sessions

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

type sparseObjectReader struct {
	inner io.Reader
	clist []io.Closer
}

func (inst *sparseObjectReader) _Impl() io.ReadCloser {
	return inst
}

func (inst *sparseObjectReader) Read(buffer []byte) (int, error) {
	return inst.inner.Read(buffer)
}

func (inst *sparseObjectReader) Close() error {
	all := inst.clist
	inst.clist = nil
	return store.CloseAll(all, true)
}

////////////////////////////////////////////////////////////////////////////////

type sparseObjectReaderBuilder struct {
	so      store.SparseObject
	profile store.Repository
	clist   []io.Closer
}

func (inst *sparseObjectReaderBuilder) appendToCList(r io.Reader) {
	c, ok := r.(io.Closer)
	if ok && c != nil {
		inst.clist = append(inst.clist, c)
	}
}

func (inst *sparseObjectReaderBuilder) closeAll() {
	all := inst.clist
	inst.clist = nil
	err := store.CloseAll(all, true)
	if err != nil {
		vlog.Error(err)
	}
}

func (inst *sparseObjectReaderBuilder) readHead(r io.Reader) (*git.Object, error) {
	// read
	buf := make([]byte, 1)
	builder := bytes.Buffer{}
	for {
		cb, err := r.Read(buf)
		if err != nil {
			return nil, err
		} else if cb != 1 {
			return nil, fmt.Errorf("cb != 1")
		}
		b := buf[0]
		if b == 0 {
			break // end of head
		} else {
			builder.WriteByte(b)
		}
	}
	// parse
	bin := builder.Bytes()
	str := string(bin)
	i := strings.IndexRune(str, ' ')
	if i < 1 {
		return nil, fmt.Errorf("bad object head: " + str)
	}
	p1 := strings.TrimSpace(str[0:i])
	p2 := strings.TrimSpace(str[i+1:])
	size, err := strconv.ParseInt(p2, 10, 0)
	if err != nil {
		return nil, err
	}
	obj := &git.Object{
		Type:   git.ObjectType(p1),
		Length: size,
	}
	return obj, nil
}

func (inst *sparseObjectReaderBuilder) open() (*git.Object, io.ReadCloser, error) {

	var reader io.Reader = nil
	defer func() {
		inst.closeAll()
	}()

	// file
	reader, err := inst.so.Path().GetIO().OpenReader(nil)
	if err != nil {
		return nil, nil, err
	}
	inst.appendToCList(reader)

	// buffer-1
	reader = bufio.NewReaderSize(reader, 1024*8)
	inst.appendToCList(reader)

	// compress
	reader, err = inst.profile.Compression().NewReader(reader)
	if err != nil {
		return nil, nil, err
	}
	inst.appendToCList(reader)

	// buffer-2
	reader = bufio.NewReaderSize(reader, 1024*8)
	inst.appendToCList(reader)

	// read head
	head, err := inst.readHead(reader)
	if err != nil {
		return nil, nil, err
	}
	head.ID = inst.so.GetID()

	//final
	r9 := &sparseObjectReader{}
	r9.clist = inst.clist
	r9.inner = reader
	inst.clist = nil
	return head, r9, nil
}
