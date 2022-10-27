package pktline

import "io"

// Reader ...
type Reader interface {
	Read() (*Packet, error)
}

// ReaderCloser ...
type ReaderCloser interface {
	Reader
	io.Closer
}
