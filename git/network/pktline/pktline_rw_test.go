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
