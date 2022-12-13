package pktline

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

// Writer ...
type Writer interface {
	Write(p *Packet) error
}

// WriterCloser ...
type WriterCloser interface {
	io.Closer
	Writer
}

// NewWriterCloser ...
func NewWriterCloser(w io.Writer, enableAutoClose bool) WriterCloser {
	return &outputWriterCloser{
		out:             w,
		enableAutoClose: enableAutoClose,
	}
}

////////////////////////////////////////////////////////////////////////////////

// outputWriterCloser ...
type outputWriterCloser struct {
	out             io.Writer
	closed          bool
	enableAutoClose bool
}

func (inst *outputWriterCloser) _Impl() WriterCloser {
	return inst
}

// Close ...
func (inst *outputWriterCloser) Close() error {
	o := inst.out
	inst.closed = true
	inst.out = nil
	if o != nil && inst.enableAutoClose {
		c, ok := o.(io.Closer)
		if ok && c != nil {
			return c.Close()
		}
	}
	return nil
}

func (inst *outputWriterCloser) makeSpecialData(p *Packet) ([]byte, error) {
	const (
		z   = '0'
		min = 0
		max = 3
	)
	n := p.Length
	if min <= n && n <= max {
		z3 := byte('0' + n)
		data := []byte{z, z, z, z3}
		return data, nil
	}
	return nil, fmt.Errorf("bad Special pktline-length:%v", n)
}

// Write ...
func (inst *outputWriterCloser) Write(p *Packet) error {

	o := inst.out
	if o == nil {
		return fmt.Errorf("stream is closed")
	}

	if p == nil {
		return nil
	}

	if p.Special {
		data, err := inst.makeSpecialData(p)
		if err != nil {
			return err
		}
		return inst.writeTo(o, data)
	}

	// else, a normal packet
	head := p.Head
	body := p.Body
	builder := bytes.Buffer{}
	builder.Write([]byte{0, 0, 0, 0})
	builder.WriteString(head)
	if body != nil {
		builder.WriteByte(0)
		builder.Write(body)
	}

	data := builder.Bytes()
	size := len(data)
	err := inst.setPackSize(data, size)
	if err != nil {
		return err
	}

	// flush all data
	return inst.writeTo(o, data)
}

func (inst *outputWriterCloser) writeTo(dst io.Writer, data []byte) error {
	want := len(data)
	have := 0
	for have < want {
		n, err := dst.Write(data[have:])
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

func (inst *outputWriterCloser) setPackSize(data []byte, size int) error {
	if size > 0xffff || size < 0 {
		return fmt.Errorf("bad pktline pack size:%v", size)
	}
	total := size
	for i := 3; i >= 0; i-- {
		num := total & 0x0f
		if 0 <= num && num <= 9 {
			data[i] = byte('0' + num)
		} else if 0x0a <= num && num <= 0x0f {
			data[i] = byte('a' + num - 0x0a)
		}
		total >>= 4
	}
	return nil
}
