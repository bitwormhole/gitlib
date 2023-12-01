package pack

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/afs"
)

type packExWrapper struct {
	inner Pack
	file  *File
}

func (inst *packExWrapper) _Impl() ComplexPack {
	return inst
}

func (inst *packExWrapper) Load() error {
	return inst.inner.Load()
}

func (inst *packExWrapper) Reload() error {
	return inst.inner.Reload()
}

func (inst *packExWrapper) Check(flags CheckFlag) error {
	return inst.inner.Check(flags)
}

func (inst *packExWrapper) GetPackID() git.PackID {
	return inst.inner.GetPackID()
}

func (inst *packExWrapper) Scan() ([]*git.PackedObjectHeaderEx, error) {
	return inst.inner.Scan()
}

func (inst *packExWrapper) OpenSimpleObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	return inst.inner.OpenSimpleObjectReader(item, pool)
}

func (inst *packExWrapper) ReadSimpleObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error) {
	return inst.inner.ReadSimpleObjectHeader(item, pool)
}

func (inst *packExWrapper) IndexToHeader(item *git.PackIndexItem) *git.PackedObjectHeaderEx {
	return inst.inner.IndexToHeader(item)
}

func (inst *packExWrapper) OpenObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	return inst.OpenComplexObjectReader(item, pool)
}

func (inst *packExWrapper) ReadObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error) {
	return inst.ReadComplexObjectHeader(item, pool)
}

func (inst *packExWrapper) OpenComplexObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	rb := packEntityReaderBuilder{}
	rb.head = *item
	rb.pack = inst
	return rb.open()
}

func (inst *packExWrapper) ReadComplexObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *packExWrapper) MakeIdx(dst afs.Path) error {
	hxlist, err := inst.Scan()
	if err != nil {
		return err
	}
	ctx := inst.file.Context
	builder := &v2IdxBuilder{context: ctx, pack: inst}
	for _, hx := range hxlist {
		err = builder.AddItem(hx)
		if err != nil {
			return err
		}
	}
	return builder.Make(dst)
}
