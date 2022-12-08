package sessions

import (
	"fmt"
	"io"
)

// DeltaInstructionCode ...
type DeltaInstructionCode int

// 定义pack-delta的指令
const (
	InstructionCopyFromBase DeltaInstructionCode = 1
	InstructionAddNewData   DeltaInstructionCode = 2
)

// DeltaInstruction 表示一条 delta 指令
type DeltaInstruction struct {
	Code   DeltaInstructionCode
	Offset uint32
	Size   uint32
	Data   []byte
}

////////////////////////////////////////////////////////////////////////////////

type deltaInstructionReader struct {
}

func (inst *deltaInstructionReader) Read(in io.Reader) (*DeltaInstruction, error) {
	b, err := inst.readByte(in)
	if err != nil {
		return nil, err
	}
	if b == 0 {
		return nil, fmt.Errorf("bad git-pack-delta instruction: Reserved")
	}
	flag := (b & 0x80)
	if flag == 0 {
		return inst.readAddNewData(b, in)

	}
	return inst.readCopyFromBase(b, in)
}

func (inst *deltaInstructionReader) readCopyFromBase(b0 int, in io.Reader) (*DeltaInstruction, error) {

	// +----------+---------+---------+---------+---------+-------+-------+-------+
	// | 1xxxxxxx | offset1 | offset2 | offset3 | offset4 | size1 | size2 | size3 |
	// +----------+---------+---------+---------+---------+-------+-------+-------+
	// in byte[0] , flags : { 1 , s3 ,s2,s1,o4,o3,o2,o1 }

	count := 0
	for flags := b0 & 0x7f; flags > 0; flags >>= 1 {
		if flags&0x01 != 0 {
			count++
		}
	}
	src, err := inst.readBytes(in, count)
	if err != nil {
		return nil, err
	}
	flags := b0 & 0x7f
	mask := 0x40
	dst := [7]byte{}
	for i := len(dst) - 1; i >= 0; i-- {
		if mask&flags != 0 {
			count--
			dst[i] = src[count]
		}
		mask >>= 1
	}
	di := &DeltaInstruction{
		Code:   InstructionCopyFromBase,
		Offset: 0,
		Size:   0,
	}
	di.Offset = inst.parseUInt32(dst[0:4])
	di.Size = inst.parseUInt32(dst[4:7])

	// the only exception
	if di.Size == 0 {
		di.Size = 0x10000
	}

	return di, nil
}

func (inst *deltaInstructionReader) parseUInt32(b []byte) uint32 {
	x := uint32(0)
	size := len(b)
	for i := size - 1; i >= 0; i-- {
		n := uint32(b[i])
		x = (x << 8) | (n & 0xff)
	}
	return x
}

func (inst *deltaInstructionReader) readAddNewData(b0 int, in io.Reader) (*DeltaInstruction, error) {
	size := b0 & 0x7f
	data, err := inst.readBytes(in, size)
	if err != nil {
		return nil, err
	}
	di := &DeltaInstruction{
		Code:   InstructionAddNewData,
		Size:   uint32(size),
		Offset: 0,
		Data:   data,
	}
	return di, nil
}

func (inst *deltaInstructionReader) readByte(in io.Reader) (int, error) {
	var buf [1]byte
	cb, err := io.ReadFull(in, buf[:])
	if cb == 1 {
		b := buf[0]
		return int(b), nil
	}
	if err == nil {
		err = fmt.Errorf("bad read size:%v", cb)
	}
	return 0, err
}

func (inst *deltaInstructionReader) readBytes(in io.Reader, size int) ([]byte, error) {
	buf := make([]byte, size)
	cb, err := io.ReadFull(in, buf) // in.Read(buf)
	if cb == size {
		return buf, nil
	}
	if err == nil {
		err = fmt.Errorf("bad read size, want:%v have:%v", size, cb)
	}
	return nil, err
}
