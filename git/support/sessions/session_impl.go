package sessions

import (
	"errors"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/data/dxo"
	"github.com/bitwormhole/gitlib/git/store"
)

type sessionImpl struct {
	profile store.RepositoryProfile
}

func (inst *sessionImpl) _Impl() store.Session {
	return inst
}

func (inst *sessionImpl) Close() error {
	return nil
}

func (inst *sessionImpl) GetRepository() store.RepositoryProfile {
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
func (inst *sessionImpl) LoadCommit(id dxo.ObjectID) (*dxo.Commit, error) {
	return nil, errors.New("no impl")
}

func (inst *sessionImpl) LoadTag(id dxo.ObjectID) (*dxo.Tag, error) {
	return nil, errors.New("no impl")
}

func (inst *sessionImpl) LoadTree(id dxo.ObjectID) (*dxo.Tree, error) {
	return nil, errors.New("no impl")
}

// HEAD
func (inst *sessionImpl) LoadHEAD(head store.HEAD) (dxo.ReferenceName, error) {
	return "", errors.New("no impl")
}
