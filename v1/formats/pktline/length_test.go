package pktline

import (
	"fmt"
	"testing"
)

func TestLengthConvert(t *testing.T) {

	const (
		min   = 0x0000
		max   = 0xffff
		step  = 135
		lower = min + step
		upper = max - step
	)

	count := 0

	for i := min; i <= max; i++ {

		n1 := LengthU16(i)
		hx1 := n1.Hex()
		n2 := hx1.Uint()

		msg := fmt.Sprintf("number[%d]: n1=%d, n2=%d, bin=0x%s", i, n1, n2, hx1.String())

		if n1 != n2 || n1 != LengthU16(i) {
			t.Error(msg)
			break
		}

		if (i < lower) || (i > upper) || ((i % step) == 0) {
			t.Log(msg)
			count++
		}

	}

	t.Logf("log.count = %d", count)
}
