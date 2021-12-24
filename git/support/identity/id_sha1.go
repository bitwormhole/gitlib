package identity

import (
	"github.com/bitwormhole/gitlib/git"
)

const sha1sumLengthInBit = 160
const sha1sumLengthInByte = sha1sumLengthInBit / 8

// sha1id 表示一个采用 sha-1 算法的 git-id
type sha1id struct {
	data [sha1sumLengthInByte]byte
}

////////////////////////////////////////////////////////////////////////////////

// CreateSha1id 创建一个sha1id
func createSha1id(data []byte) (*sha1id, error) {
	inst := &sha1id{}
	err := initID(data, inst.data[:], sha1sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

// ParseSha1id 解析一个 sha1 的 git-id
func parseSha1id(str string) (*sha1id, error) {
	inst := &sha1id{}
	err := parseID(str, inst.data[:], sha1sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

////////////////////////////////////////////////////////////////////////////////

// ToPackID 转型为 git.PackID 对象
func (inst *sha1id) ToPackID() git.PackID {
	return inst
}

// ToObjectID 转型为 git.ObjectID 对象
func (inst *sha1id) ToObjectID() git.ObjectID {
	return inst
}

// String 把ID转换为字符串形式
func (inst *sha1id) String() string {
	return stringifyBytes(inst.data[:])
}

// Bytes 把ID转换为字节形式
func (inst *sha1id) Bytes() []byte {
	return inst.GetBytes(nil)
}

// GetBytes 把ID转换为字节形式
func (inst *sha1id) GetBytes(buffer []byte) []byte {
	return inst.data[:]
}

// GetFactory 取ID工厂
func (inst *sha1id) GetFactory() git.IdentityFactory {
	return GetSha1IDFactory()
}

////////////////////////////////////////////////////////////////////////////////

// sha1IDFactory 是一个 git-id 工厂（in sha-1）
type sha1IDFactory struct {
}

func (inst *sha1IDFactory) _Impl() git.IdentityFactory {
	return inst
}

// Algorithm 取算法
func (inst *sha1IDFactory) Algorithm() string {
	return "SHA-1"
}

// Size in bits
func (inst *sha1IDFactory) Size() int {
	return sha1sumLengthInBit
}

// Parse 解析ID
func (inst *sha1IDFactory) Parse(s string) (git.ID, error) {
	return parseSha1id(s)
}

// Create 创建ID
func (inst *sha1IDFactory) Create(b []byte) (git.ID, error) {
	return createSha1id(b)
}
