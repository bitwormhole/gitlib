package com4session

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects"
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// SessionFacadeRegistry ...
type SessionFacadeRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *SessionFacadeRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *SessionFacadeRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:          true,
		OnInitForSession: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *SessionFacadeRegistry) create(ctx *repositories.SessionContext) (any, error) {
	facade := new(sessionFacade)
	facade.context = ctx
	ctx.Facade = facade
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type sessionFacade struct {
	context *repositories.SessionContext
}

func (inst *sessionFacade) _impl() repositories.Session { return inst }

func (inst *sessionFacade) noImpl() error {
	return fmt.Errorf("no impl")
}

func (inst *sessionFacade) Close() error {
	return inst.noImpl()
}

func (inst *sessionFacade) GetRepository() repositories.Repository {
	return inst.context.Parent.Facade
}

// 根据名称，取指定的组件
func (inst *sessionFacade) GetComponent(name string) (any, error) {
	err := inst.noImpl()
	return nil, err
}

// 取工作目录
func (inst *sessionFacade) GetWD() afs.Path {
	panic("no impl")
}

// 取仓库布局
func (inst *sessionFacade) GetLayout() repositories.Layout {
	panic("no impl")
}

// 新建临时文件
func (inst *sessionFacade) NewTemporaryFile(dir afs.Path) afs.Path {
	panic("no impl")
}

func (inst *sessionFacade) NewTemporaryBuffer(dir afs.Path) repositories.TemporaryBuffer {
	panic("no impl")
}

func (inst *sessionFacade) NewReaderPool(size int) afs.ReaderPool {
	panic("no impl")
}

func (inst *sessionFacade) GetReaderPool() afs.ReaderPool {
	panic("no impl")
}

func (inst *sessionFacade) GetObjectContext() *objects.Context {
	panic("no impl")
}

/////////////////////////////////////////////////

// config
// SaveConfig(cfg Config) error
// LoadConfig(cfg Config) error

// objects
func (inst *sessionFacade) ReadObject(id git.ObjectID) (*git.Object, io.ReadCloser, error) {
	panic("no impl")
}

func (inst *sessionFacade) WriteObject(o *git.Object, data io.Reader) (*git.Object, error) {
	panic("no impl")
}

func (inst *sessionFacade) GetSparseObjects() repositories.SparseObjects {
	panic("no impl")
}

func (inst *sessionFacade) GetPacks() repositories.Packs {
	panic("no impl")
}

func (inst *sessionFacade) LoadText(id git.ObjectID) (string, *git.Object, error) {
	panic("no impl")
}

func (inst *sessionFacade) LoadBinary(id git.ObjectID) ([]byte, *git.Object, error) {
	panic("no impl")
}

// commit, tag, tree
func (inst *sessionFacade) LoadCommit(id git.ObjectID) (*git.Commit, error) {
	panic("no impl")
}

func (inst *sessionFacade) LoadTag(id git.ObjectID) (*git.Tag, error) {
	panic("no impl")
}

func (inst *sessionFacade) LoadTree(id git.ObjectID) (*git.Tree, error) {
	panic("no impl")
}

// refs
func (inst *sessionFacade) LoadRef(r repositories.Ref) (*git.Ref, error) {
	panic("no impl")
}

func (inst *sessionFacade) SaveRef(r *git.Ref) error {
	panic("no impl")
}

// HEAD
func (inst *sessionFacade) LoadHEAD(head repositories.HEAD) (*git.HEAD, error) {
	panic("no impl")
}

func (inst *sessionFacade) SaveHEAD(h *git.HEAD) error {
	panic("no impl")
}
