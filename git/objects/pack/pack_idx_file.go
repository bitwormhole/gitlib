package pack

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
)

// Idx 表示一个 pack-*.idx 文件
type Idx interface {
	Load() error
	Reload() error
	Check(flags CheckFlag) error
	GetPackID() git.PackID
}

////////////////////////////////////////////////////////////////////////////////

// NewIdx ...
func NewIdx(file *File) (Idx, error) {
	idx := &idxFileFacade{}
	err := idx.Init(file)
	if err != nil {
		return nil, err
	}
	err = idx.Load()
	if err != nil {
		return nil, err
	}
	return idx, nil
}

////////////////////////////////////////////////////////////////////////////////

// idxFileFacade ...
type idxFileFacade struct {
	file *File
	impl Idx
}

func (inst *idxFileFacade) _Impl() Idx {
	return inst
}

// Init ...
func (inst *idxFileFacade) Init(file *File) error {
	if file == nil {
		return fmt.Errorf("param:file is nil")
	}
	inst.file = file.Clone()
	i, err := inst.loadImpl()
	if err != nil {
		return err
	}
	inst.impl = i
	return nil
}

func (inst *idxFileFacade) isV2() (bool, error) {
	in, err := inst.file.OpenReader()
	if err != nil {
		return false, err
	}
	defer func() { in.Close() }()
	in.Seek(0, io.SeekStart)
	cr := commonReader{}
	n, err := cr.readUInt32(in)
	if err != nil {
		return false, err
	}
	isv2 := (n == MagicNumberIdxV2)
	return isv2, nil
}

func (inst *idxFileFacade) loadImpl() (Idx, error) {
	isv2, err := inst.isV2()
	if err != nil {
		return nil, err
	}
	idx := inst.impl
	if isv2 {
		idx = &idxFileV2{file: inst.file}
	} else {
		idx = &idxFileV1{file: inst.file}
	}
	return idx, nil
}

// Check ...
func (inst *idxFileFacade) Check(flags CheckFlag) error {
	return inst.impl.Check(flags)
}

// Load ...
func (inst *idxFileFacade) Load() error {
	return inst.impl.Load()
}

// Reload ...
func (inst *idxFileFacade) Reload() error {
	return inst.impl.Reload()
}

// ReadPackID ...
func (inst *idxFileFacade) GetPackID() git.PackID {
	return inst.impl.GetPackID()
}
