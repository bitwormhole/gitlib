package pack

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
)

// idxFileV2 ...
type idxFileV2 struct {
	file    *File
	fanout  *git.PackIdxFanOut
	total   int64 // ids 条目数量
	version int
	magic   int

	checkSizeOK bool
	checkHeadOK bool
	hasOffset64 bool
}

func (inst *idxFileV2) _Impl() Idx {
	return inst
}

func (inst *idxFileV2) hasCheckFlag(a, b CheckFlag) bool {
	return (a & b) != 0
}

// Check ...
func (inst *idxFileV2) Check(flags CheckFlag) error {

	err := inst.Load()
	if err != nil {
		return err
	}

	list := make([]func() error, 0)
	if inst.hasCheckFlag(flags, CheckHead) {
		list = append(list, inst.checkHead)
	}
	if inst.hasCheckFlag(flags, CheckSize) {
		list = append(list, inst.checkSize)
	}
	if inst.hasCheckFlag(flags, CheckSum) {
		list = append(list, inst.checkSum)
	}

	for _, step := range list {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

// Load ...
func (inst *idxFileV2) Load() error {
	fo := inst.fanout
	if fo != nil {
		return nil
	}
	return inst.Reload()
}

// Reload ...
func (inst *idxFileV2) Reload() error {

	in, err := inst.file.OpenReader()
	if err != nil {
		return err
	}
	defer func() { in.Close() }()
	cr := commonReader{}

	magic, err := cr.readUInt32(in)
	if err != nil {
		return err
	}

	version, err := cr.readUInt32(in)
	if err != nil {
		return err
	}

	fo, err := cr.readIdxFanOut(in)
	if err != nil {
		return err
	}

	if fo == nil {
		return fmt.Errorf("fan-out table is nil")
	}

	inst.magic = int(magic)
	inst.version = int(version)
	inst.fanout = fo
	inst.total = fo.Total()

	return inst.Check(CheckSize | CheckHead)
}

func (inst *idxFileV2) checkHead() error {

	const wantVersion = 2
	const wantMagic = MagicNumberIdxV2

	magic := inst.magic
	version := inst.version
	fo := inst.fanout

	if magic != wantMagic {
		return fmt.Errorf("bad idx-v2 magic number:%v, want:%v", inst.magic, wantMagic)
	}

	if version != wantVersion {
		return fmt.Errorf("bad version number:%v", version)
	}

	if fo == nil {
		return fmt.Errorf("git-pack-idx fanout table is nil")
	}

	// for fan-out
	item1 := uint32(0)
	for _, item2 := range fo.Data {
		if item1 <= item2 {
			item1 = item2
		} else {
			return fmt.Errorf("bad fan-out value")
		}
	}

	inst.checkHeadOK = true
	return nil
}

func (inst *idxFileV2) checkSize() error {

	file := inst.file.Path
	digest := inst.file.Digest
	// pool := inst.file.Pool

	// check size
	idSize1 := digest.Size()
	idSize := int64(idSize1.SizeInBytes())
	fileSize := file.GetInfo().Length()
	headSize := inst.headSize()
	total := inst.total

	count := int64(0)
	count += headSize

	// ids
	// crc32
	// offset(4-bytes)
	count += ((idSize + 4 + 4) * total)

	// tail
	count += (idSize * 2)

	// compare
	count1 := count
	count2 := count + (8 * total) // add 8-bytes offset
	if (count1 != fileSize) && (count2 != fileSize) {
		return fmt.Errorf("bad file size")
	}
	inst.hasOffset64 = (count2 == fileSize)
	inst.checkSizeOK = true
	return nil
}

func (inst *idxFileV2) checkSum() error {

	return nil
}

func (inst *idxFileV2) headSize() int64 {
	return 4 * (2 + 256) // sizeof(magic) + sizeof(version) + sizeof(fanout)
}

// ReadPackID ...
func (inst *idxFileV2) ReadPackID() (git.PackID, error) {

	file := inst.file.Path
	digest := inst.file.Digest
	pool := inst.file.Pool

	// check size
	idSize1 := digest.Size()
	idSize := int64(idSize1.SizeInBytes())
	fileSize := file.GetInfo().Length()
	headSize := inst.headSize()

	if fileSize < headSize+(2*idSize) {
		return nil, fmt.Errorf("bad file size")
	}

	// open
	in, err := pool.OpenReader(file, nil)
	if err != nil {
		return nil, err
	}
	defer func() { in.Close() }()

	// seek
	pos1 := fileSize - (2 * idSize)
	pos2, err := in.Seek(pos1, io.SeekStart)
	if err != nil {
		return nil, err
	}

	if pos1 != pos2 {
		return nil, fmt.Errorf("bad seek position")
	}

	// read
	cr := commonReader{}
	return cr.readPackID(in, idSize1)
}
