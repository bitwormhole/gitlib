package pktline

import "io"

func NewReader(target io.Reader) Reader {
	rd := new(innerDefaultReader)
	rd.target = target
	return rd
}

////////////////////////////////////////////////////////////////////////////////

type innerDefaultReader struct {
	target io.Reader
}

// Read implements Reader.
func (i *innerDefaultReader) Read(p []byte) (n int, err error) {
	return i.target.Read(p)
}

// ReadLine implements Reader.
func (i *innerDefaultReader) ReadLine(buffer *Line) (*Line, error) {
	panic("unimplemented")
}
