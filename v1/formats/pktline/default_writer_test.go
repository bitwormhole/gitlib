package pktline

import (
	"bytes"
	"crypto/sha1"
	"testing"
)

func TestDefaultWriter(t *testing.T) {

	out := new(bytes.Buffer)
	wtr := NewWriter(out)
	linelist := make([]*Line, 0)

	// make lines

	line := new(Line)
	linelist = append(linelist, line)

	line = &Line{
		Length:  0,
		Special: true,
	}
	linelist = append(linelist, line)

	line = &Line{
		Length:  1,
		Special: true,
	}
	linelist = append(linelist, line)

	line = &Line{
		Length:  2,
		Special: true,
	}
	linelist = append(linelist, line)

	line = &Line{
		Length:  3,
		Special: true,
	}
	linelist = append(linelist, line)

	line = &Line{
		Length:  4,
		Special: true,
	}
	linelist = append(linelist, line)

	text := "abc123"
	sum := sha1.Sum([]byte(text))

	line = &Line{}
	line.Text = text
	linelist = append(linelist, line)

	line = &Line{}
	line.Binary = sum[:]
	linelist = append(linelist, line)

	line = &Line{}
	line.Text = text
	line.Binary = sum[:]
	linelist = append(linelist, line)

	line = &Line{}
	line.Text = "EOF"
	linelist = append(linelist, line)

	line = &Line{}
	line.Special = true
	linelist = append(linelist, line)

	// write lines

	for _, li := range linelist {
		_, err := wtr.WriteLine(li)
		if err != nil {
			t.Error(err)
			return
		}
	}

	all := out.Bytes()
	str := string(all)
	t.Log(str)
}
