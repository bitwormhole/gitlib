package pktline

import (
	"errors"
	"io"
	"time"
)

// Reader 是 pktline 包的读取接口
type Reader interface {
	Read() (*Packet, error)
}

// ReadCloser 是 pktline 包的读取接口(支持close方法)
type ReadCloser interface {
	Reader
	io.Closer
}

////////////////////////////////////////////////////////////////////////////////

const pktlenFieldSize = 4

// StreamReadCloser 提供从流中读取 pktline 包的方法
type StreamReadCloser struct {
	r io.Reader
	c io.Closer
	// buffer1 [4]byte
	// buffer2 bytes.Buffer
}

func (inst *StreamReadCloser) _Impl() ReadCloser {
	return inst
}

// InitWithReader 初始化
func (inst *StreamReadCloser) InitWithReader(r io.Reader) error {
	inst.r = r
	return nil
}

// InitWithReadCloser 初始化
func (inst *StreamReadCloser) InitWithReadCloser(rc io.ReadCloser) error {
	inst.r = rc
	inst.c = rc
	return nil
}

func (inst *StreamReadCloser) readInBytes(buf []byte, src io.Reader) error {
	haveSize := 0
	wantSize := len(buf)
	for haveSize < wantSize {
		cnt, err := src.Read(buf)
		if err != nil {
			return err
		} else if cnt > 0 {
			haveSize += cnt
			buf = buf[cnt:]
		} else {
			time.Sleep(time.Millisecond * 10)
		}
	}
	return nil
}

func (inst *StreamReadCloser) readPacketLength(src io.Reader) (int, error) {
	var lengthBuffer [pktlenFieldSize]byte
	err := inst.readInBytes(lengthBuffer[:], src)
	if err != nil {
		return 0, err
	}
	n := 0
	for i := 0; i < pktlenFieldSize; i++ {
		ch := int(lengthBuffer[i])
		digit := 0
		if '0' <= ch && ch <= '9' {
			digit = (ch - '0')
		} else if 'a' <= ch && ch <= 'f' {
			digit = 0x0a + (ch - 'a')
		} else if 'A' <= ch && ch <= 'F' {
			digit = 0x0a + (ch - 'A')
		} else {
			return 0, errors.New("bad pktline packet format")
		}
		n = (n << 4) | (digit)
	}
	return n, nil
}

// readNextPack 读取下一个包，结果中不包含开头4bytes的长度字段
func (inst *StreamReadCloser) readNextPack() ([]byte, int, error) {
	src := inst.r
	// read length
	pktlen, err := inst.readPacketLength(src)
	if err != nil {
		return nil, 0, err
	}
	if pktlen < 0 || pktlen > 0xffff {
		return nil, 0, errors.New("bad pktline-packet-size")
	} else if pktlen <= pktlenFieldSize {
		data := []byte{}
		return data, pktlen, nil
	}
	// read head+body
	datalen := pktlen - pktlenFieldSize
	data := make([]byte, datalen)
	err = inst.readInBytes(data, src)
	if err != nil {
		return nil, 0, err
	}
	return data, pktlen, nil
}

// Read 读取包
func (inst *StreamReadCloser) Read() (*Packet, error) {

	data, pktlen, err := inst.readNextPack()
	if err != nil {
		return nil, err
	}

	// find end of head
	ending := len(data)
	ok := false
	i0 := 0 // index of '\0' ,the end of head
	for i := 0; i < ending; i++ {
		if data[i] == 0 {
			i0 = i
			ok = true
			break
		}
	}

	// make packet
	p := &Packet{}
	p.Length = pktlen
	if ok {
		p.Head = string(data[0:i0])
		p.Body = data[i0+1:]
	} else {
		p.Head = string(data)
		p.Body = []byte{}
	}
	return p, nil
}

// Close 关闭流
func (inst *StreamReadCloser) Close() error {
	c := inst.c
	if c == nil {
		return nil
	}
	return c.Close()
}
