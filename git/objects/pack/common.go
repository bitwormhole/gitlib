package pack

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
)

type commonReader struct {
	buffer8b [8]byte
	buffer4b [4]byte
}

func (inst *commonReader) readUInt32(in io.Reader) (uint32, error) {
	value := uint32(0)
	buf := inst.buffer4b[:]
	n, err := in.Read(buf)
	if n == len(buf) {
		for _, b := range buf {
			value = (value << 8) | (0x00ff & uint32(b))
		}
	}
	return value, err
}

func (inst *commonReader) readUInt64(in io.Reader) (uint64, error) {
	value := uint64(0)
	buf := inst.buffer8b[:]
	n, err := in.Read(buf)
	if n == len(buf) {
		for _, b := range buf {
			value = (value << 8) | (0x00ff & uint64(b))
		}
	}
	return value, err
}

func (inst *commonReader) readIdxFanOut(in io.Reader) (*git.PackIdxFanOut, error) {
	fanout := &git.PackIdxFanOut{}
	data := fanout.Data[:]
	for i := range data {
		n, err := inst.readUInt32(in)
		if err != nil {
			return nil, err
		}
		data[i] = n
	}
	return fanout, nil
}

func (inst *commonReader) readPackID(in io.Reader, size git.HashSize) (git.PackID, error) {
	hex, err := inst.readHexID(in, size)
	if err != nil {
		return nil, err
	}
	return git.CreatePackID(hex)
}

func (inst *commonReader) readObjectID(in io.Reader, size git.HashSize) (git.ObjectID, error) {
	hex, err := inst.readHexID(in, size)
	if err != nil {
		return nil, err
	}
	return git.CreateObjectID(hex)
}

func (inst *commonReader) readHexID(in io.Reader, size git.HashSize) ([]byte, error) {
	sizeInBytes := size.SizeInBytes()
	buf := make([]byte, sizeInBytes)
	n, err := in.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != sizeInBytes {
		return nil, fmt.Errorf("bad id size: %v-bits", sizeInBytes*8)
	}
	return buf, nil
}
