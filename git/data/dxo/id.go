package dxo

import "github.com/bitwormhole/starter/util"

////////////////////////////////////////////////////////////////////////////////

// HexID 表示git的一个 hash 值
type HexID string

// ObjectID 表示git的 object-ID
type ObjectID HexID

// PackID 表示git的包ID
type PackID HexID

////////////////////////////////////////////////////////////////////////////////

func (id HexID) String() string {
	return string(id)
}

// Bytes ...
func (id HexID) Bytes() []byte {
	str := string(id)
	return string2bytes(str)
}

// Size ...
func (id HexID) Size() int {
	cb := len(id)
	return cb * 4
}

////////////////////////////////////////////////////////////////////////////////

func (id ObjectID) String() string {
	return string(id)
}

// Bytes ...
func (id ObjectID) Bytes() []byte {
	str := string(id)
	return string2bytes(str)
}

// Size ...
func (id ObjectID) Size() int {
	cb := len(id)
	return cb * 4
}

////////////////////////////////////////////////////////////////////////////////

func (id PackID) String() string {
	return string(id)
}

// Bytes ...
func (id PackID) Bytes() []byte {
	str := string(id)
	return string2bytes(str)
}

// Size ...
func (id PackID) Size() int {
	cb := len(id)
	return cb * 4
}

////////////////////////////////////////////////////////////////////////////////

func string2bytes(str string) []byte {
	hex, err := util.HexFromString(str)
	if err != nil {
		return []byte{}
	}
	return hex.Bytes()
}
