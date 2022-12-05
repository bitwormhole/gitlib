package git

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/util"
)

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
type HashID interface {
	GetFactory() IdentityFactory

	Size() HashSize

	Bytes() []byte

	String() string

	GetByte(index int) byte
}

////////////////////////////////////////////////////////////////////////////////

// ObjectID 表示git的 object-ID
type ObjectID interface {
	HashID
	OID() ObjectID
}

// PackID 表示git的包ID
type PackID interface {
	HashID
	PID() PackID
}

////////////////////////////////////////////////////////////////////////////////

// CommonID ... 表示一个多用途的ID
type CommonID interface {
	HashID
	ObjectID
	PackID
}

////////////////////////////////////////////////////////////////////////////////

// IdentityFactory 表示git的一个 hash 值
type IdentityFactory interface {
	Size() HashSize

	Zero() CommonID

	Create(b []byte) CommonID

	Parse(s string) CommonID

	TryCreate(b []byte) (CommonID, error)

	TryParse(s string) (CommonID, error)
}

////////////////////////////////////////////////////////////////////////////////

var theDefautIdentityFactory IdentityFactory = nil

// DefaultIdentityFactory ...
func DefaultIdentityFactory() IdentityFactory {
	f := theDefautIdentityFactory
	if f == nil {
		size := HashSizeInBytes(160 / 8) // default.size = 160-bits (sha-1)
		f = &baseCommonIDFactory{size: size}
		theDefautIdentityFactory = f
	}
	return f
}

// SetDefaultIdentityFactory ...
func SetDefaultIdentityFactory(f IdentityFactory) {
	if f == nil {
		return
	}
	theDefautIdentityFactory = f
}

// ParseObjectID ...
func ParseObjectID(s string) (ObjectID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryParse(s)
	if err != nil {
		return nil, err
	}
	return id.(ObjectID), nil
}

// CreateObjectID ...
func CreateObjectID(b []byte) (ObjectID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryCreate(b)
	if err != nil {
		return nil, err
	}
	return id.(ObjectID), nil
}

// ParsePackID ...
func ParsePackID(s string) (PackID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryParse(s)
	if err != nil {
		return nil, err
	}
	return id.(PackID), nil
}

// CreatePackID ...
func CreatePackID(b []byte) (PackID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryCreate(b)
	if err != nil {
		return nil, err
	}
	return id.(PackID), nil
}

// HashEqual 判断两个ID是否相等
func HashEqual(h1, h2 HashID) bool {
	if h1 == nil || h2 == nil {
		return false
	}
	size1 := h1.Size()
	size2 := h2.Size()
	if size1 != size2 {
		return false
	}
	size := size1.SizeInBytes()
	for i := 0; i < size; i++ {
		digit1 := h1.GetByte(i)
		digit2 := h2.GetByte(i)
		if digit1 != digit2 {
			return false
		}
	}
	return true
}

// HashCompare 比较两个ID的大小
func HashCompare(h1, h2 HashID) int {
	if h1 == nil || h2 == nil {
		return 0
	}
	size1 := h1.Size()
	size2 := h2.Size()
	if size1 < size2 {
		return 1
	} else if size1 > size2 {
		return -1
	}
	size := size1.SizeInBytes()
	for i := 0; i < size; i++ {
		digit1 := h1.GetByte(i)
		digit2 := h2.GetByte(i)
		if digit1 < digit2 {
			return 1
		} else if digit1 > digit2 {
			return -1
		}
	}
	return 0
}

// HashZero 取 0
func HashZero() HashID {
	f := DefaultIdentityFactory()
	return f.Zero()
}

// HashBytes 稳妥的取 bytes
func HashBytes(x HashID) []byte {
	if x == nil {
		x = HashZero()
	}
	return x.Bytes()
}

// HashString 稳妥的取 string
func HashString(x HashID) string {
	if x == nil {
		x = HashZero()
	}
	return x.String()
}

////////////////////////////////////////////////////////////////////////////////

type baseCommonIDFactory struct {
	size HashSize

	// cache
	zeroID CommonID
}

func (inst *baseCommonIDFactory) _Impl() IdentityFactory {
	return inst
}

func (inst *baseCommonIDFactory) Size() HashSize {
	return inst.size
}

func (inst *baseCommonIDFactory) Zero() CommonID {
	z := inst.zeroID
	if z != nil {
		return z
	}
	size := inst.size.SizeInBytes()
	data := make([]byte, size)
	z = &baseCommonID{data: data, factory: inst}
	inst.zeroID = z
	return z
}

func (inst *baseCommonIDFactory) Create(b []byte) CommonID {
	id, err := inst.TryCreate(b)
	if err != nil {
		return inst.Zero()
	}
	return id
}

func (inst *baseCommonIDFactory) Parse(s string) CommonID {
	id, err := inst.TryParse(s)
	if err != nil {
		return inst.Zero()
	}
	return id
}

func (inst *baseCommonIDFactory) TryCreate(src []byte) (CommonID, error) {
	if src == nil {
		return nil, fmt.Errorf("id data is nil")
	}
	wantSize := inst.size.SizeInBytes()
	haveSize := len(src)
	if wantSize != haveSize {
		return nil, fmt.Errorf("bad id size, have:%v want: %v", haveSize, wantSize)
	}
	dst := make([]byte, wantSize)
	copy(dst, src)
	id := &baseCommonID{factory: inst, data: dst}
	return id, nil
}

func (inst *baseCommonIDFactory) TryParse(s string) (CommonID, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	data, err := util.ParseHexString(s)
	if err != nil {
		return nil, err
	}
	haveSize := len(data)
	wantSize := inst.size.SizeInBytes()
	if wantSize != haveSize {
		return nil, fmt.Errorf("bad id size, have:%v want: %v", haveSize, wantSize)
	}
	id := &baseCommonID{factory: inst, data: data}
	return id, nil
}

/////////////////////////////////////////////////////

type baseCommonID struct {
	data    []byte
	factory *baseCommonIDFactory
}

func (inst *baseCommonID) _Impl() CommonID {
	return inst
}

func (inst *baseCommonID) GetFactory() IdentityFactory {
	return inst.factory
}

func (inst *baseCommonID) GetByte(index int) byte {
	data := inst.data
	size := len(data)
	if 0 <= index && index < size {
		return data[index]
	}
	return 0
}

func (inst *baseCommonID) Bytes() []byte {
	size := inst.factory.size.SizeInBytes()
	src := inst.data
	dst := make([]byte, size)
	copy(dst, src)
	return dst
}

func (inst *baseCommonID) String() string {
	return util.StringifyBytes(inst.data)
}

func (inst *baseCommonID) Size() HashSize {
	return inst.factory.size
}

func (inst *baseCommonID) OID() ObjectID {
	return inst
}

func (inst *baseCommonID) PID() PackID {
	return inst
}
