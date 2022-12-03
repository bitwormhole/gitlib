package git

// PackIndexItem ...
type PackIndexItem struct {
	PID        PackID
	OID        ObjectID
	Exists     bool
	Index      int64 // the index in idx table
	Offset     int64 // the offset in pack file
	CRC32      uint32
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
