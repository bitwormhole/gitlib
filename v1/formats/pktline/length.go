package pktline

import "fmt"

////////////////////////////////////////////////////////////////////////////////

// LengthU16 表示 pkt-line 的长度, 是一个16位 uint
type LengthU16 uint16

// LengthHex 表示长度为 4-bytes 的 hex 字符串
type LengthHex [4]byte

type Length = LengthU16

////////////////////////////////////////////////////////////////////////////////

func (num LengthU16) String() string {
	hex := num.Hex()
	return hex.String()
}

func (num LengthU16) Hex() LengthHex {
	var bin LengthHex

	bin[0] = num.innerConvert4bits(num >> 12)
	bin[1] = num.innerConvert4bits(num >> 8)
	bin[2] = num.innerConvert4bits(num >> 4)
	bin[3] = num.innerConvert4bits(num)

	return bin
}

func (num LengthU16) innerConvert4bits(num4bits LengthU16) byte {
	n4b := uint(num4bits & 0x0f)
	if n4b <= 9 {
		return '0' + byte(n4b)
	} else if 0x0a <= n4b && n4b <= 0x0f {
		return 'a' + byte(n4b-0x0a)
	}
	return '0'
}

////////////////////////////////////////////////////////////////////////////////

func (array LengthHex) String() string {
	return string(array[:])
}

func (array LengthHex) Uint() LengthU16 {
	value := array.innerChar2uint(0)
	value = (value << 4) | array.innerChar2uint(1)
	value = (value << 4) | array.innerChar2uint(2)
	value = (value << 4) | array.innerChar2uint(3)
	return LengthU16(value)
}

func (array LengthHex) Normalize() (dst LengthHex, err error) {
	for i, b := range array {
		if '0' <= b && b <= '9' {
			// ok
		} else if 'a' <= b && b <= 'f' {
			// ok
		} else if 'A' <= b && b <= 'F' {
			// to lower
			tmp := b - 'A'
			b = tmp + 'a'
		} else {
			b = '0'
			str := array.String()
			err = fmt.Errorf("bad format of pkt-line::length : [%s]", str)
		}
		dst[i] = b
	}
	return
}

func (array LengthHex) innerChar2uint(index int) uint16 {
	b := array[index]
	if ('0' <= b) && (b <= '9') {
		return uint16(b - '0')

	} else if ('a' <= b) && (b <= 'f') {
		return uint16(b-'a') + 0x0a

	} else if ('A' <= b) && (b <= 'F') {
		return uint16(b-'A') + 0x0a
	}
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// EOF
