package pktline

import "io"

// Writer ...
type Writer interface {
	Write(p *Packet) error
}

// WriterCloser ...
type WriterCloser interface {
	io.Closer
	Writer
}
