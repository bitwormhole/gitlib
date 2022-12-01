package pack

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
)

// Pack 表示一个 pack-*.pack 文件
type Pack interface {
	Check(flags CheckFlag) error
	ReadPackID() (git.PackID, error)
}

////////////////////////////////////////////////////////////////////////////////

// NewPack ...
func NewPack(file *File) (Pack, error) {
	p := &EntityFile{}
	err := p.Init(file)
	if err != nil {
		return nil, err
	}
	return p, nil
}

////////////////////////////////////////////////////////////////////////////////

// EntityFile ...
type EntityFile struct {
	file *File
}

func (inst *EntityFile) _Impl() Pack {
	return inst
}

// Init ...
func (inst *EntityFile) Init(file *File) error {
	return nil
}

// Check ...
func (inst *EntityFile) Check(flags CheckFlag) error {
	return nil
}

// ReadPackID ...
func (inst *EntityFile) ReadPackID() (git.PackID, error) {
	return nil, fmt.Errorf("no impl: EntityFile.ReadPackID()")
}
