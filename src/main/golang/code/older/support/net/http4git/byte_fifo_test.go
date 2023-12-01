package http4git

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"testing"
	"time"

	"github.com/bitwormhole/starter/util"
)

func TestByteFIFO(t *testing.T) {

	t0 := time.Now()

	fifo, err := OpenByteFIFO()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		fifo.Close()
	}()

	var sum1 []byte = nil
	go func() {
		h1 := sha256.New()
		buffer1 := make([]byte, 1024)
		for {
			n, err := fifo.Read(buffer1)
			if n > 0 {
				h1.Write(buffer1[0:n])
			} else if n == 0 && err == io.EOF {
				break
			}
		}
		sum1 = h1.Sum([]byte{})
	}()

	const total = 1024 * 1024 * 1
	src := rand.Reader
	h2 := sha256.New()
	buffer2 := make([]byte, 1024)
	count := 0
	for count < total {
		n, err := src.Read(buffer2)
		if n > 0 {
			x := buffer2[0:n]
			fifo.Write(x)
			h2.Write(x)
			count += n
		} else if n == 0 && err == io.EOF {
			break
		}
		time.Sleep(time.Millisecond)
	}
	sum2 := h2.Sum([]byte{})

	fifo.Close()

	// wait for sum1
	for {
		if sum1 != nil {
			break
		}
		time.Sleep(time.Second)
	}

	t1 := time.Now()

	result := bytes.Compare(sum1, sum2)
	if result != 0 {
		str1 := util.StringifyBytes(sum1)
		str2 := util.StringifyBytes(sum2)
		t.Errorf("bad hash, sum1=%v, sum2=%v", str1, str2)
	}

	diff := t0.Sub(t1)
	ms := diff.Milliseconds()
	t.Logf("hash %v bytes, cost %v ms", count, ms)
}
