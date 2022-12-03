package sessions

import (
	"fmt"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type packEntityReader struct {
	source       io.ReadCloser
	c            io.Closer
	file         afs.Path
	closed       bool
	entityType   git.PackedObjectType
	entityLength int64
}

func (inst *packEntityReader) _Impl() io.ReadCloser {
	return inst
}

func (inst *packEntityReader) Read(b []byte) (int, error) {
	if inst.closed {
		return 0, fmt.Errorf("this stream is closed")
	}
	return inst.source.Read(b)
}

func (inst *packEntityReader) Close() error {
	clist := make([]io.Closer, 0)
	clist = append(clist, inst.source)
	clist = append(clist, inst.c)
	for _, c := range clist {
		if c == nil {
			continue
		}
		err := c.Close()
		if err != nil {
			vlog.Warn(err)
		}
	}
	inst.closed = true
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type packEntityReaderBuilder struct {
	file         afs.Path // the pack-*.pack file
	input        io.ReadSeekCloser
	index        *git.PackIndexItem
	session      store.Session
	entityType   git.PackedObjectType
	entityLength int64
}

func (inst *packEntityReaderBuilder) open() (io.ReadCloser, error) {

	// seek
	pos1 := inst.index.Offset
	in1 := inst.input
	pos2, err := in1.Seek(pos1, io.SeekStart)
	if err != nil {
		return nil, err
	}
	if pos1 != pos2 {
		return nil, fmt.Errorf("bad seek, want:%v have:%v", pos1, pos2)
	}

	// read head
	ot, ol, err := inst.readEntityHead(in1)
	if err != nil {
		return nil, err
	}

	// unzip
	compr := inst.session.GetRepository().Compression()
	in2, err := compr.NewReader(in1)
	if err != nil {
		return nil, err
	}

	// result
	inst.entityType = ot
	inst.entityLength = ol
	return &packEntityReader{
		source:       in2,
		c:            in1,
		file:         inst.file,
		entityType:   ot,
		entityLength: ol,
	}, nil
}

func (inst *packEntityReaderBuilder) readEntityHead(in io.Reader) (git.PackedObjectType, int64, error) {
	const size = 12
	b1 := [size]byte{}
	objType := git.PackedObjectType(0)
	objLength := int64(0)
	mbits := int64(0) // bits to shift
	eoh := false
	for i := 0; i < size; i++ {
		buf := b1[i : i+1]
		n, err := in.Read(buf)
		if err != nil {
			return 0, 0, err
		} else if n != 1 {
			return 0, 0, fmt.Errorf("bad head in stream")
		}
		b := int64(buf[0])
		eoh = ((b & 0x80) == 0) // End of Head
		if i == 0 {
			objType = git.PackedObjectType((b >> 4) & 0x07)
			objLength = b & 0x0f
			mbits = 4
		} else {
			objLength |= ((b & 0x7f) << mbits)
			mbits += 7
		}
		if eoh {
			break
		}
	}
	if !eoh {
		return 0, 0, fmt.Errorf("no End of Head byte")
	}
	return git.PackedObjectType(objType), objLength, nil
}

////////////////////////////////////////////////////////////////////////////////
