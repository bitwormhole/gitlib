package pktline

import (
	"fmt"
	"io"
	"time"
)

// Reader ...
type Reader interface {
	Read() (*Packet, error)
}

// ReaderCloser ...
type ReaderCloser interface {
	Reader
	io.Closer
}

// NewReaderCloser ...
func NewReaderCloser(r io.Reader, enableAutoClose bool) ReaderCloser {
	return &inputReaderCloser{
		in:              r,
		enableAutoClose: enableAutoClose,
	}
}

////////////////////////////////////////////////////////////////////////////////

// inputReaderCloser ...
type inputReaderCloser struct {
	in              io.Reader
	closed          bool
	enableAutoClose bool
}

func (inst *inputReaderCloser) _Impl() ReaderCloser {
	return inst
}

// Close ...
func (inst *inputReaderCloser) Close() error {
	i := inst.in
	inst.closed = true
	inst.in = nil
	if i != nil && inst.enableAutoClose {
		c, ok := i.(io.Closer)
		if ok && c != nil {
			return c.Close()
		}
	}
	return nil
}

// Read ...
func (inst *inputReaderCloser) Read() (*Packet, error) {

	in := inst.in
	if in == nil {
		return nil, fmt.Errorf("stream is closed")
	}

	// read total length

	const buffer1size = 4 // buffer for '0000' ,16-bits, 4-runes
	buffer1raw := [buffer1size]byte{}
	buffer1 := buffer1raw[:]

	err := inst.readFrom(in, buffer1)
	if err != nil {
		return nil, err
	}

	total, err := inst.getInt16(buffer1)
	if err != nil {
		return nil, err
	}

	pack := &Packet{Length: total}
	buffer2size := total - buffer1size
	if buffer2size <= 0 {
		// flush
		pack.Flush = true
		return pack, nil
	}

	// read head & body

	buffer2 := make([]byte, buffer2size)
	err = inst.readFrom(in, buffer2)
	if err != nil {
		return nil, err
	}

	// find '\0'

	i0 := 0
	has0 := false
	for i, b := range buffer2 {
		if b == 0 {
			has0 = true
			i0 = i
			break
		}
	}

	if has0 {
		head := buffer2[0:i0]
		body := buffer2[i0+1:]
		pack.Head = string(head)
		pack.Body = body
	} else {
		pack.Head = string(buffer2)
		pack.Body = []byte{}
	}

	return pack, nil
}

func (inst *inputReaderCloser) readFrom(src io.Reader, buf []byte) error {
	want := len(buf)
	have := 0
	for have < want {
		n, err := src.Read(buf[have:])
		if n > 0 {
			have += n
		} else {
			time.Sleep(time.Millisecond)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *inputReaderCloser) getInt16(buf []byte) (int, error) {
	n := 0
	for _, b := range buf {
		num := 0
		if '0' <= b && b <= '9' {
			num = int(b - '0')
		} else if 'a' <= b && b <= 'f' {
			num = int(b-'a') + 10
		} else if 'A' <= b && b <= 'F' {
			num = int(b-'A') + 10
		} else {
			return 0, fmt.Errorf("bad pktline-pack-size")
		}
		n = (n << 4) | num
	}
	return n, nil
}
