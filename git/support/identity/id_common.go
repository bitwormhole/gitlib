package identity

import (
	"errors"
	"strings"
)

func stringify4bits(n int, builder *strings.Builder) {
	n = n & 0x0f
	if (0 <= n) && (n <= 9) {
		builder.WriteRune(rune('0' + n))
	} else {
		builder.WriteRune(rune('a' + n - 0x0a))
	}
}

func stringifyBytes(array []byte) string {
	size := len(array)
	builder := &strings.Builder{}
	for i := 0; i < size; i++ {
		n := int(array[i])
		stringify4bits(n>>4, builder)
		stringify4bits(n, builder)
	}
	return builder.String()
}

func rune2int(r rune) (int, error) {
	if ('0' <= r) && (r <= '9') {
		return int(r - '0'), nil
	} else if ('a' <= r) && (r <= 'f') {
		return int(r - 'a' + 0x0a), nil
	} else if ('A' <= r) && (r <= 'F') {
		return int(r - 'A' + 0x0a), nil
	} else {
		return 0, errors.New("bad hex char")
	}
}

func parseID(str string, buffer []byte, size int) error {
	chs := []rune(str)
	length := len(chs)
	if length != (size * 2) {
		return errors.New("bad sum string length")
	}
	for i := 0; i < length; i += 2 {
		nh, err1 := rune2int(chs[i])
		nl, err2 := rune2int(chs[i+1])
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
		buffer[i/2] = byte((nh << 4) | nl)
	}
	return nil
}

func initID(src []byte, dst []byte, size int) error {
	len1 := len(src)
	len2 := len(dst)
	if len1 != size {
		return errors.New("bad src buffer size")
	}
	if len2 != size {
		return errors.New("bad dst buffer size")
	}
	for i := size - 1; i >= 0; i-- {
		dst[i] = src[i]
	}
	return nil
}
