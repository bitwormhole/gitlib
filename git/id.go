package git

import (
	"bytes"
	"fmt"

	"github.com/starter-go/base/lang"
)

// EmptyID 表示空的 ID 值
const EmptyID = ""

////////////////////////////////////////////////////////////////////////////////

// HashSize 表示 Hash 的长度, 单位是 byte
type HashSize int

// SizeInBytes ...
func (size HashSize) SizeInBytes() int {
	return int(size)
}

// SizeInBits ...
func (size HashSize) SizeInBits() int {
	return int(size * 8)
}

// SizeInRune ...
func (size HashSize) SizeInRune() int {
	return int(size * 2)
}

// IsValid ...
func (size HashSize) IsValid() bool {
	const minSize = 12
	return size >= minSize
}

// HashSizeInBytes ...
func HashSizeInBytes(n int) HashSize {
	return HashSize(n)
}

////////////////////////////////////////////////////////////////////////////////

// HashID 表示git的一个 hash 值
type HashID lang.Hex

func (hid HashID) hex() lang.Hex {
	return lang.Hex(hid)
}

// Size ...
func (hid HashID) Size() HashSize {
	const minSize = 6
	str := hid.String()
	size := len(str)
	if size < minSize {
		size = minSize
	}
	return HashSizeInBytes(size / 2)
}

// Bytes ...
func (hid HashID) Bytes() []byte {
	return hid.hex().Bytes()
}

func (hid HashID) String() string {
	return hid.hex().String()
}

// Equals ...
func (hid HashID) Equals(other HashID) bool {
	return HashEqual(hid, other)
}

// CreateHashID 创建 HashID
func CreateHashID(b []byte) (HashID, error) {
	if b == nil {
		z := HashZero()
		return z, fmt.Errorf("buffer is nil")
	}
	if len(b) < 1 {
		z := HashZero()
		return z, fmt.Errorf("buffer is empty")
	}
	hex := lang.HexFromBytes(b)
	return HashID(hex), nil
}

// ParseHashID 解析 HashID
func ParseHashID(s string) (HashID, error) {
	hex := lang.Hex(s)
	bin := hex.Bytes()
	id, err := CreateHashID(bin)
	if err != nil {
		err = fmt.Errorf("bad hash-id string:%s", s)
	}
	return id, err
}

////////////////////////////////////////////////////////////////////////////////

// ObjectID 表示git的 object-ID
type ObjectID HashID

func (id ObjectID) hex() HashID {
	return HashID(id)
}

// Bytes ...
func (id ObjectID) Bytes() []byte {
	return id.hex().Bytes()
}

func (id ObjectID) String() string {
	return id.hex().String()
}

// HID ...
func (id ObjectID) HID() HashID {
	return id.hex()
}

// Size ...
func (id ObjectID) Size() HashSize {
	return id.hex().Size()
}

// Equals ...
func (id ObjectID) Equals(other ObjectID) bool {
	h1 := HashID(id)
	h2 := HashID(other)
	return HashEqual(h1, h2)
}

// CompareObjectIDs 比较两个 ObjectID 的大小
func CompareObjectIDs(a, b ObjectID) int {
	h1 := a.hex()
	h2 := b.hex()
	return HashCompare(h1, h2)
}

////////////////////////////////////////////////////////////////////////////////

// PackID 表示git的包ID
type PackID HashID

func (id PackID) hex() HashID {
	return HashID(id)
}

// Bytes ...
func (id PackID) Bytes() []byte {
	return id.hex().Bytes()
}

func (id PackID) String() string {
	return id.hex().String()
}

// Size ...
func (id PackID) Size() HashSize {
	return id.hex().Size()
}

// HID ...
func (id PackID) HID() HashID {
	return id.hex()
}

// Equals ...
func (id PackID) Equals(other PackID) bool {
	h1 := HashID(id)
	h2 := HashID(other)
	return HashEqual(h1, h2)
}

////////////////////////////////////////////////////////////////////////////////

// CommonID ... 表示一个多用途的ID
type CommonID HashID

func (id CommonID) hex() HashID {
	return HashID(id)
}

// HID ...
func (id CommonID) HID() HashID {
	return id.hex()
}

// Bytes ...
func (id CommonID) Bytes() []byte {
	return id.hex().Bytes()
}

func (id CommonID) String() string {
	return id.hex().String()
}

// Size ...
func (id CommonID) Size() HashSize {
	return id.hex().Size()
}

// Equals ...
func (id CommonID) Equals(other CommonID) bool {
	h1 := HashID(id)
	h2 := HashID(other)
	return HashEqual(h1, h2)
}

////////////////////////////////////////////////////////////////////////////////

// ParseObjectID ...
func ParseObjectID(s string) (ObjectID, error) {
	id, err := ParseHashID(s)
	return ObjectID(id), err
}

// CreateObjectID ...
func CreateObjectID(b []byte) (ObjectID, error) {
	id, err := CreateHashID(b)
	return ObjectID(id), err
}

// ParsePackID ...
func ParsePackID(s string) (PackID, error) {
	id, err := ParseHashID(s)
	return PackID(id), err
}

// CreatePackID ...
func CreatePackID(b []byte) (PackID, error) {
	id, err := CreateHashID(b)
	return PackID(id), err
}

// HashEqual 判断两个ID是否相等
func HashEqual(h1, h2 HashID) bool {
	b1 := h1.Bytes()
	b2 := h2.Bytes()
	return bytes.Equal(b1, b2)
}

// HashCompare 比较两个ID的大小
func HashCompare(h1, h2 HashID) int {
	b1 := h1.Bytes()
	b2 := h2.Bytes()
	return bytes.Compare(b1, b2)
}

// HashZero 取 0
func HashZero() HashID {
	// 160b:0123456789012345678901234567890123456789
	return "0000000000000000000000000000000000000000"
}

// HashBytes 稳妥的取 bytes
func HashBytes(x HashID) []byte {
	if x == "" {
		x = HashZero()
	}
	return x.Bytes()
}

// HashString 稳妥的取 string
func HashString(x HashID) string {
	if x == "" {
		x = HashZero()
	}
	return x.String()
}

////////////////////////////////////////////////////////////////////////////////
