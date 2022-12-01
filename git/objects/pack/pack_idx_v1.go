package pack

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
)

type idxFileV1 struct {
	file   *File
	fanout *git.PackIdxFanOut
	total  int64
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
func (inst *idxFileV1) ReadPackID() (git.PackID, error) {
	return nil, fmt.Errorf("no impl: idxFileV1.ReadPackID()")
}
