package sessions

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

type sparseObjectsImpl struct {
	session store.Session
}

func (inst *sparseObjectsImpl) _Impl() store.SparseObjects {
	return inst
}

func (inst *sparseObjectsImpl) init() {
	return
}

func (inst *sparseObjectsImpl) Close() error {
	return nil
}

func (inst *sparseObjectsImpl) ReadSparseObjectRaw(o store.SparseObject) (io.ReadCloser, error) {
	file := o.Path()
	pool := inst.session.GetReaderPool()
	return pool.OpenReader(file, nil)
}

func (inst *sparseObjectsImpl) WriteSparseObject(o *git.Object, data io.Reader) (*git.Object, error) {
	saver := plainSparseObjectSaver{session: inst.session}
	return saver.Save(o, data)
}

func (inst *sparseObjectsImpl) WriteSparseObjectRaw(o *git.Object, data io.Reader) (*git.Object, error) {
	saver := rawSparseObjectSaver{session: inst.session}
	return saver.Save(o, data)
}

func (inst *sparseObjectsImpl) ReadSparseObject(o store.SparseObject) (*git.Object, io.ReadCloser, error) {
	repo := inst.session.GetRepository()
	in := &sparseObjectReaderBuilder{
		so:      o,
		profile: repo,
	}
	return in.open()
}
