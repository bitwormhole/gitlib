package pack

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
)

// Pack 表示一个 pack-*.pack 文件
type Pack interface {
	Load() error
	Reload() error
	Check(flags CheckFlag) error
	GetPackID() git.PackID
}

////////////////////////////////////////////////////////////////////////////////

// NewPack ...
func NewPack(file *File) (Pack, error) {
	p := &EntityFile{}
	err := p.Init(file)
	if err != nil {
		return nil, err
	}
	err = p.Load()
	if err != nil {
		return nil, err
	}
	return p, nil
}

////////////////////////////////////////////////////////////////////////////////

// EntityFile ...
type EntityFile struct {
	file *File
	pid  git.PackID

	signature [4]byte
	version   uint32
	count     uint32

	loaded  bool
	loading bool
}

func (inst *EntityFile) _Impl() Pack {
	return inst
}

// Init ...
func (inst *EntityFile) Init(file *File) error {
	inst.file = file
	return nil
}

func (inst *EntityFile) hasCheckFlag(a, b CheckFlag) bool {
	return (a & b) != 0
}

// Check ...
func (inst *EntityFile) Check(flags CheckFlag) error {

	err := inst.Load()
	if err != nil {
		return err
	}

	list := make([]func() error, 0)
	list = append(list, inst.Load)

	if inst.hasCheckFlag(flags, CheckHead) {
		list = append(list, inst.checkHead)
	}
	if inst.hasCheckFlag(flags, CheckSize) {
		list = append(list, inst.checkSize)
	}
	if inst.hasCheckFlag(flags, CheckSum) {
		list = append(list, inst.checkSum)
	}

	for _, step := range list {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

// func (inst *EntityFile) loadHead() error {
// 	return nil
// }

func (inst *EntityFile) checkSize() error {
	return nil
}

func (inst *EntityFile) checkSum() error {
	ch := packSumChecker{}
	return ch.check(inst.file)
}

func (inst *EntityFile) checkHead() error {

	signature := string(inst.signature[:])
	version := inst.version

	if signature != "PACK" {
		return fmt.Errorf("bad pack.signature, want:'PACK'")
	}

	if version != 2 && version != 3 {
		return fmt.Errorf("bad pack.version, want: 2 or 3")
	}

	return nil
}

// GetPackID ...
func (inst *EntityFile) GetPackID() git.PackID {
	return inst.pid
}

// Load ...
func (inst *EntityFile) Load() error {
	if inst.loaded {
		return nil
	}
	err := inst.Reload()
	if err != nil {
		return err
	}
	inst.loaded = true
	return nil
}

// Reload ...
func (inst *EntityFile) Reload() error {

	// state
	if inst.loading {
		return nil
	}
	inst.loading = true
	defer func() {
		inst.loading = false
	}()

	// open
	in, err := inst.file.OpenReader()
	if err != nil {
		return err
	}
	defer func() {
		in.Close()
	}()
	cr := commonReader{}

	// signature string
	signature, err := cr.read4Bytes(in)
	if err != nil {
		return err
	}

	// version   uint32
	version, err := cr.readUInt32(in)
	if err != nil {
		return err
	}

	// count     uint32
	count, err := cr.readUInt32(in)
	if err != nil {
		return err
	}

	// tail for id
	_, pid2, err := cr.readPackFileTail(inst.file)
	if err != nil {
		return err
	}

	// keep values
	inst.signature = signature
	inst.version = version
	inst.count = count
	inst.pid = pid2

	return inst.Check(CheckSize | CheckHead)
}
