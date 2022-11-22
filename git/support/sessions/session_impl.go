package sessions

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/data/gitfmt"
	"github.com/bitwormhole/gitlib/git/store"
)

type sessionImpl struct {
	profile store.RepositoryProfile
	// core    *store.Core
}

func (inst *sessionImpl) _Impl() store.Session {
	return inst
}

func (inst *sessionImpl) Close() error {
	return nil
}

func (inst *sessionImpl) getSmallObjectSizeMax() int {
	return 8 * 1024 * 1024
}

func (inst *sessionImpl) GetRepository() store.Repository {
	return inst.profile
}

// 根据名称，取指定的组件
func (inst *sessionImpl) GetComponent(name string) (any, error) {

	return "", errors.New("no impl")
}

// 取工作目录
func (inst *sessionImpl) GetWD() afs.Path {
	return inst.profile.Layout().WD()
}

// 取仓库布局
func (inst *sessionImpl) GetLayout() store.RepositoryLayout {
	return inst.profile.Layout()
}

// objects

func (inst *sessionImpl) LoadText(id git.ObjectID) (string, error) {
	bin, err := inst.LoadBinary(id)
	if err != nil {
		return "", err
	}
	return string(bin), nil
}

func (inst *sessionImpl) LoadBinary(id git.ObjectID) ([]byte, error) {
	reader, o, err := inst.ReadObject(id)
	if err != nil {
		return nil, err
	}
	defer func() {
		reader.Close()
	}()
	limit := inst.getSmallObjectSizeMax()
	if o.Length > int64(limit) {
		const f = "the size of small object is over limit, size=%v limit=%v id=%v"
		return nil, fmt.Errorf(f, o.Length, limit, id)
	}
	return ioutil.ReadAll(reader)
}

func (inst *sessionImpl) LoadCommit(id git.ObjectID) (*git.Commit, error) {
	text, err := inst.LoadText(id)
	if err != nil {
		return nil, err
	}
	commit, err := gitfmt.ParseCommit(text)
	if err != nil {
		return nil, err
	}
	commit.ID = id
	return commit, nil
}

func (inst *sessionImpl) LoadTag(id git.ObjectID) (*git.Tag, error) {
	text, err := inst.LoadText(id)
	if err != nil {
		return nil, err
	}
	tag, err := gitfmt.ParseTag(text)
	if err != nil {
		return nil, err
	}
	tag.ID = id
	return tag, nil
}

func (inst *sessionImpl) LoadTree(id git.ObjectID) (*git.Tree, error) {
	bin, err := inst.LoadBinary(id)
	if err != nil {
		return nil, err
	}
	tree, err := gitfmt.ParseTree(bin)
	if err != nil {
		return nil, err
	}
	tree.ID = id
	return tree, nil
}

// HEAD ...
func (inst *sessionImpl) LoadHEAD(h store.HEAD) (*git.HEAD, error) {
	text, err := h.Path().GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return gitfmt.ParseHEAD(text)
}

// LoadRef ...
func (inst *sessionImpl) LoadRef(r store.Ref) (*git.Ref, error) {
	text, err := r.Path().GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return gitfmt.ParseRef(text)
}

func (inst *sessionImpl) ReadPackObject(o store.PackObject) (io.ReadCloser, *store.Object, error) {
	return nil, nil, errors.New("no impl")
}

func (inst *sessionImpl) ReadSparseObject(o store.SparseObject) (io.ReadCloser, *store.Object, error) {
	in := &sparseObjectReaderBuilder{
		so:      o,
		profile: inst.profile,
	}
	return in.open()
}

func (inst *sessionImpl) ReadSparseObjectRaw(o store.SparseObject) (io.ReadCloser, error) {
	return nil, errors.New("no impl")
}

func (inst *sessionImpl) WriteSparseObject(o *store.Object, data io.Reader) (*store.Object, error) {
	return nil, errors.New("no impl")
}

func (inst *sessionImpl) WriteSparseObjectRaw(o *store.Object, data io.Reader) (*store.Object, error) {
	return nil, errors.New("no impl")
}

func (inst *sessionImpl) ReadObject(id git.ObjectID) (io.ReadCloser, *store.Object, error) {
	so := inst.profile.Objects().GetSparseObject(id)
	return inst.ReadSparseObject(so)
}
