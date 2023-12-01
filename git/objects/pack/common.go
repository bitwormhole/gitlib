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

func (inst *commonReader) readFull(in io.Reader, p []byte) (int, error) {
	return io.ReadFull(in, p)
}

func (inst *commonReader) read4Bytes(in io.Reader) ([4]byte, error) {
	buf := inst.buffer4b[:]
	n, err := inst.readFull(in, buf)
	if n == len(buf) && err == nil {
		return inst.buffer4b, nil
	}
	inst.buffer4b[0] = 0
	inst.buffer4b[1] = 0
	inst.buffer4b[2] = 0
	inst.buffer4b[3] = 0
	if err == nil {
		err = fmt.Errorf("bad read size, want:4 have:%v", n)
	}
	return inst.buffer4b, err
}

func (inst *commonReader) readUInt32(in io.Reader) (uint32, error) {
	value := uint32(0)
	buf := inst.buffer4b[:]
	n, err := inst.readFull(in, buf) // in.Read(buf)
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
	n, err := inst.readFull(in, buf) // in.Read(buf)
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
		return "", err
	}
	return git.CreatePackID(hex)
}

func (inst *commonReader) readObjectID(in io.Reader, size git.HashSize) (git.ObjectID, error) {
	hex, err := inst.readHexID(in, size)
	if err != nil {
		return "", err
	}
	return git.CreateObjectID(hex)
}

func (inst *commonReader) readHexID(in io.Reader, size git.HashSize) ([]byte, error) {
	sizeInBytes := size.SizeInBytes()
	buf := make([]byte, sizeInBytes)
	n, err := inst.readFull(in, buf) // in.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != sizeInBytes {
		return nil, fmt.Errorf("bad id size: %v-bits", sizeInBytes*8)
	}
	return buf, nil
}

func (inst *commonReader) readPackFileTail(f *File) (git.PackID, git.PackID, error) {

	ctx := f.Context.Parent
	file := f.Path
	digest := ctx.Digest
	pool := ctx.Pool

	// check size
	idSize1 := digest.Size()
	idSize := int64(idSize1.SizeInBytes())
	fileSize := file.GetInfo().Length()
	tailSize := (2 * idSize)

	if fileSize < tailSize {
		return "", "", fmt.Errorf("bad file size")
	}

	// open
	in, err := pool.OpenReader(file, nil)
	if err != nil {
		return "", "", err
	}
	defer func() { in.Close() }()

	// seek
	pos1 := fileSize - tailSize
	pos2, err := in.Seek(pos1, io.SeekStart)
	if err != nil {
		return "", "", err
	}
	if pos1 != pos2 {
		return "", "", fmt.Errorf("bad seek position")
	}

	// read
	pid1, err := inst.readPackID(in, idSize1)
	if err != nil {
		return "", "", err
	}
	pid2, err := inst.readPackID(in, idSize1)
	if err != nil {
		return "", "", err
	}

	return pid1, pid2, nil
}

////////////////////////////////////////////////////////////////////////////////

type commonWriter struct {
	buffer8b [8]byte
	buffer4b [4]byte
}

func (inst *commonWriter) writeInt64(n int64, dst io.Writer) error {
	value := n
	data := inst.buffer8b[:]
	size := len(data)
	for i := size - 1; i >= 0; i-- {
		data[i] = byte(value & 0x00ff)
		value >>= 8
	}
	return inst.writeBytes(data, dst)
}

func (inst *commonWriter) writeUInt32(n uint32, dst io.Writer) error {
	value := n
	data := inst.buffer4b[:]
	size := len(data)
	for i := size - 1; i >= 0; i-- {
		data[i] = byte(value & 0x00ff)
		value >>= 8
	}
	return inst.writeBytes(data, dst)
}

func (inst *commonWriter) writeBytes(p []byte, dst io.Writer) error {
	if p == nil || dst == nil {
		return fmt.Errorf("param is nil")
	}
	have, err := dst.Write(p)
	if err != nil {
		return err
	}
	want := len(p)
	if have != want {
		return fmt.Errorf("bad writing-data size")
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
