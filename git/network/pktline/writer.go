package pktline

import (
	"bytes"
	"errors"
	"io"
	"time"
)

// Writer 是 pktline 包的写入接口
type Writer interface {
	Write(p *Packet) error
}

// WriteCloser 是 pktline 包的写入接口(支持close方法)
type WriteCloser interface {
	Writer
	io.Closer
}

////////////////////////////////////////////////////////////////////////////////

// StreamWriteCloser ...
type StreamWriteCloser struct {
	c io.Closer
	w io.Writer
}

func (inst *StreamWriteCloser) _Impl() WriteCloser {
	return inst
}

// InitWithWriteCloser ...
func (inst *StreamWriteCloser) InitWithWriteCloser(w io.WriteCloser) {
	inst.c = w
	inst.w = w
}

// InitWithWriter ...
func (inst *StreamWriteCloser) InitWithWriter(w io.Writer) {
	inst.w = w
}

// Close ...
func (inst *StreamWriteCloser) Close() error {
	c := inst.c
	if c == nil {
		return nil
	}
	return c.Close()
}

func (inst *StreamWriteCloser) Write(p *Packet) error {
	if inst.isEmptyPacket(p) {
		return inst.writeEmptyPacket(p)
	}
	body := p.Body
	buffer := bytes.Buffer{}
	buffer.WriteString("0000")
	buffer.WriteString(p.Head)
	buffer.WriteByte(0)
	if body != nil {
		buffer.Write(body)
	}
	data := buffer.Bytes()
	pktlen := len(data)
	err := inst.preparePktlinePacketLengthField(pktlen, data)
	if err != nil {
		return err
	}
	return inst.writeTo(inst.w, data)
}

func (inst *StreamWriteCloser) isEmptyPacket(p *Packet) bool {
	if p.Head != "" {
		return false
	}
	body := p.Body
	if body != nil {
		if len(body) > 0 {
			return false
		}
	}
	return true
}

func (inst *StreamWriteCloser) writeEmptyPacket(p *Packet) error {
	buffer := [pktlenFieldSize]byte{'0', '0', '0', '0'}
	plen := p.Length
	if 0 <= plen && plen <= pktlenFieldSize {
		buffer[pktlenFieldSize-1] = byte('0' + plen)
	}
	return inst.writeTo(inst.w, buffer[:])
}

func (inst *StreamWriteCloser) preparePktlinePacketLengthField(length int, buffer []byte) error {
	if length < 4 || length > 0xffff {
		return errors.New("bad pktline packet size")
	}
	n := length
	for i := 3; i >= 0; i-- {
		digit := 0x0f & n
		n >>= 4
		if 0 <= digit && digit <= 9 {
			buffer[i] = byte('0' + digit)
		} else if 0x0a <= digit && digit <= 0x0f {
			buffer[i] = byte(digit - 0x0a + 'a')
		}
	}
	return nil
}

func (inst *StreamWriteCloser) writeTo(dst io.Writer, data []byte) error {
	wantSize := len(data)
	doneSize := 0
	for doneSize < wantSize {
		cnt, err := dst.Write(data)
		if err != nil {
			return err
		} else if cnt > 0 {
			doneSize += cnt
			data = data[cnt:]
		} else {
			time.Sleep(time.Millisecond * 10)
		}
	}
	return nil
}
