package git

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/util"
)

////////////////////////////////////////////////////////////////////////////////

// HashID 表示git的一个 hash 值
type HashID interface {
	GetFactory() IdentityFactory

	// size in bits
	Size() int

	Bytes() []byte

	String() string
}

////////////////////////////////////////////////////////////////////////////////

// ObjectID 表示git的 object-ID
type ObjectID interface {
	HashID
}

// PackID 表示git的包ID
type PackID interface {
	HashID
}

////////////////////////////////////////////////////////////////////////////////

// IdentityFactory 表示git的一个 hash 值
type IdentityFactory interface {

	// size in bits
	Size() int

	Zero() HashID

	Create(b []byte) HashID

	Parse(s string) HashID

	TryCreate(b []byte) (HashID, error)

	TryParse(s string) (HashID, error)
}

////////////////////////////////////////////////////////////////////////////////

var theDefautIdentityFactory IdentityFactory = nil

// DefaultIdentityFactory ...
func DefaultIdentityFactory() IdentityFactory {
	f := theDefautIdentityFactory
	if f == nil {
		f = &commonIDFactory{size: 160} // default.size = 160-bits (sha-1)
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
	return id, nil
}

// CreateObjectID ...
func CreateObjectID(b []byte) (ObjectID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryCreate(b)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// ParsePackID ...
func ParsePackID(s string) (PackID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryParse(s)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// CreatePackID ...
func CreatePackID(b []byte) (PackID, error) {
	f := DefaultIdentityFactory()
	id, err := f.TryCreate(b)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// HashEqual 判断两个ID是否相等
func HashEqual(x, y HashID) bool {
	if x == nil || y == nil {
		return false
	}
	b1 := x.Bytes()
	b2 := y.Bytes()
	return bytes.Equal(b1, b2)
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

type commonIDFactory struct {
	size int // in bits

	// cache
	zeroID HashID
}

func (inst *commonIDFactory) _Impl() IdentityFactory {
	return inst
}

func (inst *commonIDFactory) Size() int {
	return inst.size
}

func (inst *commonIDFactory) sizeInByte() int {
	return inst.size / 8
}

func (inst *commonIDFactory) sizeInRune() int {
	return inst.size / 4
}

func (inst *commonIDFactory) Zero() HashID {
	z := inst.zeroID
	if z != nil {
		return z
	}
	size := inst.sizeInByte()
	data := make([]byte, size)
	z = &commonID{data: data, factory: inst}
	inst.zeroID = z
	return z
}

func (inst *commonIDFactory) Create(b []byte) HashID {
	id, err := inst.TryCreate(b)
	if err != nil {
		return inst.Zero()
	}
	return id
}

func (inst *commonIDFactory) Parse(s string) HashID {
	id, err := inst.TryParse(s)
	if err != nil {
		return inst.Zero()
	}
	return id
}

func (inst *commonIDFactory) TryCreate(src []byte) (HashID, error) {
	if src == nil {
		return nil, fmt.Errorf("id data is nil")
	}
	wantSize := inst.sizeInByte()
	haveSize := len(src)
	if wantSize != haveSize {
		return nil, fmt.Errorf("bad id size, have:%v want: %v", haveSize, wantSize)
	}
	dst := make([]byte, wantSize)
	copy(dst, src)
	id := &commonID{factory: inst, data: dst}
	return id, nil
}

func (inst *commonIDFactory) TryParse(s string) (HashID, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	data, err := util.ParseHexString(s)
	if err != nil {
		return nil, err
	}
	haveSize := len(data)
	wantSize := inst.sizeInByte()
	if wantSize != haveSize {
		return nil, fmt.Errorf("bad id size, have:%v want: %v", haveSize, wantSize)
	}
	id := &commonID{factory: inst, data: data}
	return id, nil
}

/////////////////////////////////////////////////////

type commonID struct {
	data    []byte
	factory *commonIDFactory
}

func (inst *commonID) _Impl() HashID {
	return inst
}

func (inst *commonID) GetFactory() IdentityFactory {
	return inst.factory
}

func (inst *commonID) Bytes() []byte {
	size := inst.factory.sizeInByte()
	src := inst.data
	dst := make([]byte, size)
	copy(dst, src)
	return dst
}

func (inst *commonID) String() string {
	return util.StringifyBytes(inst.data)
}

// size in bits
func (inst *commonID) Size() int {
	return inst.factory.size
}
