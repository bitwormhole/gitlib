package pack

import (
	"bytes"
	"fmt"
	"hash"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
)

type packSumChecker struct {
	file     afs.Path
	digest   git.Digest
	hashSize git.HashSize

	hash hash.Hash
	tail bytes.Buffer

	fileSize int64 // the total file size
	mainSize int64 // the file size without tail
	count    int64 // the count of current position
}

func (inst *packSumChecker) check(file *File) error {

	// cr := commonReader{}

	err := inst.prepare(file)
	if err != nil {
		return err
	}

	in, err := file.OpenReader()
	if err != nil {
		return err
	}
	defer func() { in.Close() }()

	err = inst.scan(in)
	if err != nil {
		return err
	}

	return inst.finish()
}

func (inst *packSumChecker) prepare(f *File) error {

	digest := f.Digest
	hashSize := digest.Size()
	file := f.Path
	fileSize := file.GetInfo().Length()
	tailSize := int64(hashSize.SizeInBytes() * 2)

	if tailSize >= fileSize {
		return fmt.Errorf("bad pack file size")
	}

	inst.digest = digest
	inst.hashSize = hashSize
	inst.file = file
	inst.fileSize = fileSize
	inst.mainSize = fileSize - tailSize
	inst.hash = digest.New()
	inst.count = 0
	return nil
}

func (inst *packSumChecker) scan(src io.Reader) error {
	const bufsize = 1024 * 64
	buffer := make([]byte, bufsize)
	for {
		cb, err := src.Read(buffer)
		if cb > 0 {
			inst.write(buffer[0:cb])
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}
	return nil
}

func (inst *packSumChecker) write(data []byte) error {

	cb1n2 := int64(len(data))
	cb1 := int64(0)
	cb2 := int64(0)

	pos1 := inst.count
	pos2 := inst.count + cb1n2
	posTail := inst.mainSize

	if pos2 < posTail {
		cb1 = cb1n2
		cb2 = 0
	} else {
		if posTail < pos1 {
			cb1 = 0
			cb2 = cb1n2
		} else {
			cb1 = posTail - pos1
			cb2 = cb1n2 - cb1
		}
	}

	if cb1 > 0 {
		// has main data
		d := data
		if cb1 < cb1n2 {
			d = data[0:cb1]
		}
		inst.hash.Write(d)
	}

	if cb2 > 0 {
		// has tail data
		d := data
		if cb2 < cb1n2 {
			d = data[cb1:]
		}
		inst.tail.Write(d)
	}

	inst.count += cb1n2
	return nil
}

func (inst *packSumChecker) getTail() ([]byte, []byte, error) {
	tail := inst.tail.Bytes()
	sumSize := inst.hashSize.SizeInBytes()
	if len(tail) != (2 * sumSize) {
		return nil, nil, fmt.Errorf("bad pack tail size")
	}
	p1 := tail[0:sumSize]
	p2 := tail[sumSize:]
	return p1, p2, nil
}

func (inst *packSumChecker) finish() error {

	want1, want2, err := inst.getTail()
	if err != nil {
		return err
	}

	h := inst.hash
	have1 := h.Sum([]byte{})

	h.Write(want1)
	have2 := h.Sum([]byte{})

	inst.log("have1", have1)
	inst.log("have2", have2)
	inst.log("want1", want1)
	inst.log("want2", want2)

	if !bytes.Equal(want2, have2) {
		path := inst.file.GetPath()
		return fmt.Errorf("bad check sum of file [%v]", path)
	}

	return nil
}

func (inst *packSumChecker) log(tag string, data []byte) {
	str := util.StringifyBytes(data)
	vlog.Warn(tag, ": ", str)
}
