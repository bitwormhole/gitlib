package git

// PackIndexItem ...
type PackIndexItem struct {
	OID    ObjectID
	Offset int64 // the offset in pack file
	CRC32  uint32

	// ext
	PID        PackID
	Exists     bool
	Index      int64 // the index in idx table
	Length     int64
	Type       ObjectType
	PackedType PackedObjectType // in-pack entity
}

// PackIndexHead ...
type PackIndexHead struct {
	Version int // 1 or 2
	Total   int64
	FanOut  PackIdxFanOut
}

// PackIndex ...
type PackIndex struct {
	Head  PackIndexHead
	Items []*PackIndexItem
}

// PackedObjectHeader ...
type PackedObjectHeader struct {
	Type PackedObjectType
	Size int64
}

// PackedObjectHeaderEx ...
type PackedObjectHeaderEx struct {
	PackedObjectHeader

	PID    PackID
	OID    ObjectID
	Offset int64

	DeltaRef    ObjectID
	DeltaOffset int64

	// // for debug
	// DeltaOffsetAbs int64 //
}

// GetDeltaParentOffset  计算 delta-parent 的绝对位置
func (inst *PackedObjectHeaderEx) GetDeltaParentOffset() int64 {
	if inst.Type == PackedDeltaOFS {
		return inst.Offset - inst.DeltaOffset
	}
	return 0
}

////////////////////////////////////////////////////////////////////////////////

// PackIdxFanOut ...
type PackIdxFanOut struct {
	Data [256]uint32
}

// Total 返回总条数
func (inst *PackIdxFanOut) Total() int64 {
	list := inst.Data[:]
	size := len(list)
	n := list[size-1]
	return int64(n)
}

////////////////////////////////////////////////////////////////////////////////
