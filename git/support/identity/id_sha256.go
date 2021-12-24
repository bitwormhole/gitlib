package identity

import (
	"github.com/bitwormhole/gitlib/git"
)

const sha256sumLengthInBit = 256
const sha256sumLengthInByte = sha256sumLengthInBit / 8

// sha256id 表示一个采用 sha-256 算法的 git-id
type sha256id struct {
	data [sha256sumLengthInByte]byte
}

////////////////////////////////////////////////////////////////////////////////

// CreateSha256id 创建一个sha256id
func createSha256id(data []byte) (*sha256id, error) {
	inst := &sha256id{}
	err := initID(data, inst.data[:], sha256sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

// ParseSha256id 解析一个 sha256 的 git-id
func parseSha256id(str string) (*sha256id, error) {
	inst := &sha256id{}
	err := parseID(str, inst.data[:], sha256sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

////////////////////////////////////////////////////////////////////////////////

// ToPackID 转型为 git.PackID 对象
func (inst *sha256id) ToPackID() git.PackID {
	return inst
}

// ToObjectID 转型为 git.ObjectID 对象
func (inst *sha256id) ToObjectID() git.ObjectID {
	return inst
}

// String 把ID转换为字符串形式
func (inst *sha256id) String() string {
	return stringifyBytes(inst.data[:])
}

// Bytes 把ID转换为字节形式
func (inst *sha256id) Bytes() []byte {
	return inst.GetBytes(nil)
}

// GetBytes 把ID转换为字节形式
func (inst *sha256id) GetBytes(buffer []byte) []byte {
	return inst.data[:]
}

// GetFactory 取ID工厂
func (inst *sha256id) GetFactory() git.IdentityFactory {
	return GetSha256IDFactory()
}

////////////////////////////////////////////////////////////////////////////////

// sha256IDFactory 是一个 git-id 工厂（in sha256）
type sha256IDFactory struct {
}

func (inst *sha256IDFactory) _Impl() git.IdentityFactory {
	return inst
}

// Algorithm 取算法
func (inst *sha256IDFactory) Algorithm() string {
	return "SHA-256"
}

// Size in bits
func (inst *sha256IDFactory) Size() int {
	return sha256sumLengthInBit
}

// Parse 解析ID
func (inst *sha256IDFactory) Parse(s string) (git.ID, error) {
	return parseSha256id(s)
}

// Create 创建ID
func (inst *sha256IDFactory) Create(b []byte) (git.ID, error) {
	return createSha256id(b)
}
