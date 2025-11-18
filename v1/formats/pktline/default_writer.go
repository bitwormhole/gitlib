package pktline

import (
	"bytes"
	"fmt"
	"io"
)

// 注意: 这个函数创建的 Writer 不是并发安全的!
func NewWriter(target io.Writer) Writer {
	w := new(innerDefaultWriter)
	w.target = target
	return w
}

////////////////////////////////////////////////////////////////////////////////

type innerDefaultWriter struct {
	target io.Writer
	buffer bytes.Buffer
}

// Write implements Writer.
func (inst *innerDefaultWriter) Write(p []byte) (n int, err error) {
	return inst.innerWrite(p)
}

// WriteLine implements Writer.
func (inst *innerDefaultWriter) WriteLine(line *Line) (*Line, error) {

	if line == nil {
		return nil, nil // ignore
	}

	if line.Special {
		return inst.innerWriteSpecial(line)
	}

	// prepare
	str := line.Text
	bin := line.Binary
	buffer := &inst.buffer
	sep := byte('?')

	// size1 := len(str)
	size2 := len(bin)
	if size2 == 0 {
		sep = '\n'
	} else {
		sep = 0
	}

	// make
	buffer.Reset()
	buffer.WriteString("len_")
	buffer.WriteString(str)
	buffer.WriteByte(sep)
	buffer.Write(bin)

	// complete
	raw := buffer.Bytes()
	err := inst.innerMakeLength(raw)
	if err != nil {
		return nil, err
	}

	// write
	_, err = inst.innerWrite(raw)
	if err != nil {
		return nil, err
	}

	// handle result
	line.Raw = raw
	return line, nil
}

func (inst *innerDefaultWriter) innerWrite(data []byte) (int, error) {
	want := len(data)
	have, err := inst.target.Write(data)
	if err != nil {
		return have, err
	}
	if have != want {
		return have, fmt.Errorf("bad write-count: want=%d, have=%d", want, have)
	}
	return have, nil
}

func (inst *innerDefaultWriter) innerMakeLength(raw []byte) error {
	const (
		min = 4
		max = 65520
	)
	total := len(raw)
	if total < min || max < total {
		return fmt.Errorf("bad length [have:%d] of pkt-line", total)
	}
	length := Length(total)
	hex := length.Hex()
	for idx := 0; idx < 4; idx++ {
		raw[idx] = hex[idx]
	}
	return nil
}

func (inst *innerDefaultWriter) innerWriteSpecial(line *Line) (*Line, error) {
	const (
		min = 0
		max = 4
	)
	length := line.Length
	if min <= length && length <= max {
		hex := length.Hex()
		data := hex[:]
		_, err := inst.innerWrite(data)
		if err != nil {
			return nil, err
		}
		return line, nil
	}
	return nil, fmt.Errorf("innerDefaultWriter.innerWriteSpecial: the line is not a special-pkt")
}
