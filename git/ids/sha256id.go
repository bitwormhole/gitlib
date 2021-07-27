package ids

import "github.com/bitwormhole/gitlib/git"

const sha256sumLengthInBit = 256
const sha256sumLengthInByte = sha256sumLengthInBit / 8

type Sha256id struct {
	data [sha256sumLengthInByte]byte
}

////////////////////////////////////////////////////////////////////////////////

func CreateSha256id(data []byte) (*Sha256id, error) {
	inst := &Sha256id{}
	err := initId(data, inst.data[:], sha256sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func ParseSha256id(str string) (*Sha256id, error) {
	inst := &Sha256id{}
	err := parseId(str, inst.data[:], sha256sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

////////////////////////////////////////////////////////////////////////////////

func (inst *Sha256id) ToPackId() git.PackId {
	return inst
}

func (inst *Sha256id) ToObjectId() git.ObjectId {
	return inst
}

func (inst *Sha256id) String() string {
	return stringifyBytes(inst.Bytes())
}

func (inst *Sha256id) Bytes() []byte {
	return inst.data[:]
}
