package pack

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
)

type idxFileV1 struct {
	file   *File
	fanout *git.PackIdxFanOut
	total  int64
	pid    git.PackID
}

func (inst *idxFileV1) _Impl() Idx {
	return inst
}

func (inst *idxFileV1) Check(flags CheckFlag) error {

	err := inst.Load()
	if err != nil {
		return err
	}

	return nil
}

// Load ...
func (inst *idxFileV1) Load() error {
	fo := inst.fanout
	if fo != nil {
		return nil
	}
	return inst.Reload()
}

// Reload ...
func (inst *idxFileV1) Reload() error {
	return fmt.Errorf("no impl: idxFileV1.Reload()")
}

// ReadPackID ...
func (inst *idxFileV1) GetPackID() git.PackID {
	return inst.pid
}

// Find ...
func (inst *idxFileV1) Find(oid git.ObjectID) (*git.PackIndexItem, error) {
	return nil, fmt.Errorf("no impl: Find")
}

// Count ...
func (inst *idxFileV1) Count() int64 {
	return 0
}

// GetItem ...
func (inst *idxFileV1) GetItem(index int64) (*git.PackIndexItem, error) {
	return nil, fmt.Errorf("no impl: GetItem")
}

// GetItems ...
func (inst *idxFileV1) GetItems(index int64, limit int) ([]*git.PackIndexItem, error) {
	return nil, fmt.Errorf("no impl: GetItems")
}
