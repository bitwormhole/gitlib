package ids

import (
	"github.com/bitwormhole/gitlib/git"
)

const sha1sumLengthInBit = 160
const sha1sumLengthInByte = sha1sumLengthInBit / 8

type Sha1id struct {
	data [sha1sumLengthInByte]byte
}

////////////////////////////////////////////////////////////////////////////////

func CreateSha1id(data []byte) (*Sha1id, error) {
	inst := &Sha1id{}
	err := initId(data, inst.data[:], sha1sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func ParseSha1id(str string) (*Sha1id, error) {
	inst := &Sha1id{}
	err := parseId(str, inst.data[:], sha1sumLengthInByte)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

////////////////////////////////////////////////////////////////////////////////

func (inst *Sha1id) ToPackId() git.PackId {
	return inst
}

func (inst *Sha1id) ToObjectId() git.ObjectId {
	return inst
}

func (inst *Sha1id) String() string {
	return stringifyBytes(inst.Bytes())
}

func (inst *Sha1id) Bytes() []byte {
	return inst.data[:]
}
