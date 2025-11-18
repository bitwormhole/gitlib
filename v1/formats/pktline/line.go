package pktline

import "io"

type Line struct {
	Length Length
	Text   string
	Binary []byte

	// 如果 Special == 1, 表示这是一个特别的包,
	// 它的 length = ['0000','0004'],
	// 具体类型取决于 Length 字段.
	Special bool

	Raw []byte // raw-data-buffer = (length + text + binary)
}

type Writer interface {
	io.Writer

	WriteLine(line *Line) (*Line, error)
}

type Reader interface {
	io.Reader

	ReadLine(buffer *Line) (*Line, error)
}
