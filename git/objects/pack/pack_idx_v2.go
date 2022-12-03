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
	pid     git.PackID

	loaded      bool
	loading     bool
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
	if inst.loaded {
		return nil
	}
	err := inst.Reload()
	if err != nil {
		return err
	}
	inst.loaded = true
	return nil
}

// Reload ...
func (inst *idxFileV2) Reload() error {

	// check state
	if inst.loading {
		return nil
	}
	inst.loading = true
	defer func() {
		inst.loading = false
	}()

	// open reader
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

	// tail for id
	pid1, _, err := cr.readPackFileTail(inst.file)
	if err != nil {
		return err
	}

	// keep values
	inst.magic = int(magic)
	inst.version = int(version)
	inst.fanout = fo
	inst.total = fo.Total()
	inst.pid = pid1

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
	ch := packSumChecker{}
	return ch.check(inst.file)
}

func (inst *idxFileV2) headSize() int64 {
	return 4 * (2 + 256) // sizeof(magic) + sizeof(version) + sizeof(fanout)
}

// ReadPackID ...
func (inst *idxFileV2) GetPackID() git.PackID {
	return inst.pid
}

// Find ...
func (inst *idxFileV2) Find(oid git.ObjectID) (*git.PackIndexItem, error) {
	result := &git.PackIndexItem{}
	dao := idxFileV2dao{}
	dao.init(inst)
	e1 := dao.do(func() error {
		item, err := dao.findItemByID(oid)
		if err != nil {
			return err
		}
		items := []*git.PackIndexItem{item}
		err = dao.readItemsFields(items)
		if err != nil {
			return err
		}
		result = item
		return nil
	})
	if e1 != nil {
		return nil, e1
	} else if result != nil {
		if git.HashEqual(oid, result.OID) {
			return result, nil
		}
	}
	return nil, fmt.Errorf("git-pack-object not found, id:%v", oid.String())
}

// Count ...
func (inst *idxFileV2) Count() int64 {
	return inst.total
}

// GetItem ...
func (inst *idxFileV2) GetItem(index int64) (*git.PackIndexItem, error) {
	items, err := inst.GetItems(index, 1)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item == nil {
			continue
		}
		if item.Index == index {
			return item, nil
		}
	}
	return nil, fmt.Errorf("no object found")
}

// GetItems ...
func (inst *idxFileV2) GetItems(index int64, limit int) ([]*git.PackIndexItem, error) {
	result := make([]*git.PackIndexItem, 0)
	dao := idxFileV2dao{}
	dao.init(inst)
	e1 := dao.do(func() error {
		items, err := dao.readListID(index, limit)
		if err != nil {
			return err
		}
		err = dao.readItemsFields(items)
		if err != nil {
			return err
		}
		result = items
		return nil
	})
	if e1 != nil {
		return nil, e1
	}
	return result, nil
}

////////////////////////////////////////////////////////////////////////////////

type idxFileV2dao struct {
	parent *idxFileV2

	reader io.ReadSeeker

	idsize           git.HashSize
	total            int64
	tablePosID       int64
	tablePosCRC      int64
	tablePosOffset32 int64
	tablePosOffset64 int64
}

func (inst *idxFileV2dao) init(parent *idxFileV2) {

	total := parent.total
	pos := int64(0)
	idsize := parent.file.Digest.Size()

	// magic uint32
	// version uint32
	// fanout [256]uint32
	pos += (4 * (2 + 256))
	// ids
	pIDs := pos
	pos += (total * int64(idsize.SizeInBytes()))
	// crc32
	pCRC32 := pos
	pos += (total * 4)
	// offset32
	pOffset32 := pos
	pos += (total * 4)
	// offset64
	pOffset64 := pos
	pos += (total * 8)

	inst.parent = parent
	inst.total = total
	inst.tablePosCRC = pCRC32
	inst.tablePosID = pIDs
	inst.tablePosOffset32 = pOffset32
	inst.tablePosOffset64 = pOffset64
	inst.idsize = idsize
}

func (inst *idxFileV2dao) do(fn func() error) error {

	if fn == nil {
		return fmt.Errorf("no handler func")
	}

	f := inst.parent.file
	in, err := f.Pool.OpenReader(f.Path, nil)
	if err != nil {
		return err
	}
	defer func() {
		inst.reader = nil
		in.Close()
	}()
	inst.reader = in

	return fn()
}

func (inst *idxFileV2dao) seekToListID(index int64) error {
	itemSize := inst.idsize.SizeInBytes()
	base := inst.tablePosID
	return inst.seek(base, int64(itemSize), index)
}

func (inst *idxFileV2dao) seekToListCRC32(index int64) error {
	const itemSize = 4
	base := inst.tablePosCRC
	return inst.seek(base, itemSize, index)
}

func (inst *idxFileV2dao) seekToListOffset32(index int64) error {
	const itemSize = 4
	base := inst.tablePosOffset32
	return inst.seek(base, itemSize, index)
}

func (inst *idxFileV2dao) seekToListOffset64(index int64) error {
	const itemSize = 8
	base := inst.tablePosOffset64
	return inst.seek(base, itemSize, index)
}

func (inst *idxFileV2dao) seek(base, itemSize, index int64) error {
	total := inst.total
	if index < 0 || total <= index {
		return fmt.Errorf("this index [%v] is overflow", index)
	}
	pos1 := base + (itemSize * index)
	pos2, err := inst.reader.Seek(pos1, io.SeekStart)
	if err != nil {
		return err
	}
	if pos1 != pos2 {
		return fmt.Errorf("cannot seek to wanted position, want:%v have:%v", pos1, pos2)
	}
	return nil
}

func (inst *idxFileV2dao) readListCRC32(items []*git.PackIndexItem) error {
	wantIndex := int64(0)
	cr := commonReader{}
	in := inst.reader
	for i, item := range items {
		if item == nil {
			return fmt.Errorf("item is nil")
		}
		if i == 0 {
			wantIndex = item.Index
			err := inst.seekToListCRC32(item.Index)
			if err != nil {
				return err
			}
		} else if wantIndex != item.Index {
			return fmt.Errorf("bad index, this list must be a increasing sequence by 1")
		}
		wantIndex++
		n, err := cr.readUInt32(in)
		if err != nil {
			return err
		}
		item.CRC32 = n
	}
	return nil
}

func (inst *idxFileV2dao) readListOffset32(items []*git.PackIndexItem) error {
	wantIndex := int64(0)
	cr := commonReader{}
	in := inst.reader
	for i, item := range items {
		if item == nil {
			return fmt.Errorf("item is nil")
		}
		if i == 0 {
			wantIndex = item.Index
			err := inst.seekToListOffset32(item.Index)
			if err != nil {
				return err
			}
		} else if wantIndex != item.Index {
			return fmt.Errorf("bad index, this list must be a increasing sequence by 1")
		}
		wantIndex++
		n, err := cr.readUInt32(in)
		if err != nil {
			return err
		}
		item.Offset = int64(n)
	}
	return nil
}

func (inst *idxFileV2dao) readListOffset64(items []*git.PackIndexItem) error {
	wantIndex := int64(0)
	cr := commonReader{}
	in := inst.reader
	for i, item := range items {
		if item == nil {
			return fmt.Errorf("item is nil")
		}
		if i == 0 {
			wantIndex = item.Index
			err := inst.seekToListOffset64(item.Index)
			if err != nil {
				return err
			}
		} else if wantIndex != item.Index {
			return fmt.Errorf("bad index, this list must be a increasing sequence by 1")
		}
		wantIndex++
		n, err := cr.readUInt64(in)
		if err != nil {
			return err
		}
		item.Offset = int64(n)
	}
	return nil
}

func (inst *idxFileV2dao) readItemsFields(items []*git.PackIndexItem) error {

	hasOffset64 := inst.parent.hasOffset64
	err := inst.readListCRC32(items)
	if err != nil {
		return err
	}

	if hasOffset64 {
		err = inst.readListOffset64(items)
		if err != nil {
			return err
		}
	} else {
		err = inst.readListOffset32(items)
		if err != nil {
			return err
		}
	}

	for _, item := range items {
		item.Exists = true
	}

	return nil
}

func (inst *idxFileV2dao) readListID(index int64, limit int) ([]*git.PackIndexItem, error) {
	err := inst.seekToListID(index)
	if err != nil {
		return nil, err
	}
	cr := commonReader{}
	in := inst.reader
	idsize := inst.idsize
	pid := inst.parent.pid
	list := make([]*git.PackIndexItem, 0)
	total := inst.total
	count := 0
	i := index
	for (i < total) && (count < limit) {
		id, err := cr.readObjectID(in, idsize)
		if err != nil {
			return nil, err
		}
		item := &git.PackIndexItem{
			Index:  i,
			OID:    id,
			PID:    pid,
			Exists: true,
		}
		list = append(list, item)
		i++
		count++
	}
	return list, nil
}

func (inst *idxFileV2dao) findItemByID(want git.ObjectID) (*git.PackIndexItem, error) {

	total := inst.total
	i1 := int64(0) // todo : use fanout
	i9 := total - 1
	ids := make(map[int64]string)
	result := &git.PackIndexItem{
		PID: inst.parent.pid,
	}

	for i1 < i9 {
		if (i9 - i1) < 3 {
			break // 顺序查找
		}
		i5 := (i1 + i9) / 2
		have, err := inst.getOIDByIndex(i5)
		if err != nil {
			return nil, err
		}
		ids[i5] = have.String()
		x := git.HashCompare(want, have)
		if 0 < x {
			// want <  have
			i9 = i5
		} else if x < 0 {
			// want > have
			i1 = i5
		} else /* x==0 */ if git.HashEqual(want, have) {
			result.Index = i5
			result.OID = want
			return result, nil
		}
	}

	// 按顺序查找 [i1,i9]
	for i := i1; i <= i9; i++ {
		have, err := inst.getOIDByIndex(i)
		if err != nil {
			return nil, err
		}
		ids[i] = have.String()
		if git.HashEqual(want, have) {
			result.Index = i
			result.OID = want
			return result, nil
		}
	}

	return nil, fmt.Errorf("git-pack-object not found, id:%v", want.String())
}

func (inst *idxFileV2dao) getOIDByIndex(index int64) (git.ObjectID, error) {
	err := inst.seekToListID(index)
	if err != nil {
		return nil, err
	}
	cr := commonReader{}
	return cr.readObjectID(inst.reader, inst.idsize)
}
