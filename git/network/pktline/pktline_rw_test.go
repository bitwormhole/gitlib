package pktline

import (
	"bytes"
	"io"
	"testing"
)

func TestReaderWriter(t *testing.T) {

	packs1 := []*Packet{}
	packs2 := []*Packet{}

	packs1 = append(packs1, &Packet{})
	packs1 = append(packs1, &Packet{Special: true, Length: 0})
	packs1 = append(packs1, &Packet{Special: true, Length: 1})
	packs1 = append(packs1, &Packet{Special: true, Length: 2})
	packs1 = append(packs1, &Packet{Special: true, Length: 3})
	packs1 = append(packs1, &Packet{Head: "example1"})
	packs1 = append(packs1, &Packet{Head: "example2", Body: []byte{'a', 'b', 'c'}})
	packs1 = append(packs1, &Packet{Body: []byte{'x', 'y', 'z'}})

	buffer := &bytes.Buffer{}
	w := NewWriterCloser(buffer, true)
	for _, p := range packs1 {
		err := w.Write(p)
		if err != nil {
			t.Error(err)
			return
		}
	}
	w.Close()

	r := NewReaderCloser(buffer, true)
	for {
		p, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Error(err)
				return
			}
		}
		packs2 = append(packs2, p)
	}

	t.Logf("%v", len(packs2))
}

func TestReaderRaw(t *testing.T) {
	buf := &bytes.Buffer{}
	r1 := NewReaderCloser(buf, false)
	raw := r1.(ReaderRaw)
	r2 := raw.GetReader()
	r2.Read([]byte{0, 1, 2, 3})
}

func TestWriterRaw(t *testing.T) {
	buf := &bytes.Buffer{}
	w1 := NewWriterCloser(buf, false)
	raw := w1.(WriterRaw)
	w2 := raw.GetWriter()
	w2.Write([]byte{0, 1, 2, 3})
}
