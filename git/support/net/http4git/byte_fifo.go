package http4git

import (
	"bytes"
	"io"
	"sync"
)

// ByteFIFO ...
type ByteFIFO interface {
	io.Writer
	io.Reader
	io.Closer
}

// OpenByteFIFO ...
func OpenByteFIFO() (ByteFIFO, error) {
	fifo := &innerByteFIFO{}
	err := fifo.open()
	if err != nil {
		return nil, err
	}
	return fifo, nil
}

////////////////////////////////////////////////////////////////////////////////

// innerByteFIFO ...
type innerByteFIFO struct {
	channel chan int
	buffer  bytes.Buffer
	mutex   sync.Mutex
	closed  bool
}

func (inst *innerByteFIFO) _Impl() ByteFIFO {
	return inst
}

func (inst *innerByteFIFO) open() error {
	chl := make(chan int, 64)
	inst.channel = chl
	return nil
}

func (inst *innerByteFIFO) writeSync(b []byte) (int, error) {
	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()
	return inst.buffer.Write(b)
}

func (inst *innerByteFIFO) readSync(b []byte) (int, error) {
	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()
	return inst.buffer.Read(b)
}

func (inst *innerByteFIFO) Write(b []byte) (int, error) {
	n, err := inst.writeSync(b)
	if err != nil {
		return n, err
	}
	inst.channel <- n
	return n, nil
}

func (inst *innerByteFIFO) Read(b []byte) (int, error) {
	for {
		n, _ := inst.readSync(b)
		if n > 0 {
			return n, nil
		} else if n == 0 {
			if inst.closed {
				return 0, io.EOF
			}
		}
		n, ok := <-inst.channel
		if ok {
			// NOP
		}
	}
}

// Close ...
func (inst *innerByteFIFO) Close() error {
	if inst.closed {
		return nil
	}
	inst.closed = true
	close(inst.channel)
	return nil
}
