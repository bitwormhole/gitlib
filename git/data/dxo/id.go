package dxo

// HexID 表示git的一个 hash 值
type HexID string

func (id HexID) String() string {
	return string(id)
}

// Bytes ...
func (id HexID) Bytes() []byte {
	return nil
}

// Size ...
func (id HexID) Size() int {
	cb := len(id)
	return cb * 4
}

////////////////////////////////////////////////////////////////////////////////

// ObjectID 表示git的 object-ID
type ObjectID HexID

// PackID 表示git的包ID
type PackID HexID
