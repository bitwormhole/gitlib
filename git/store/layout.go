package store

import (
	"bitwormhole.com/starter/afs"
)

// RepositoryLayout 表示关于一个仓库的几个关键路径
type RepositoryLayout interface {

	// 取工作目录
	WD() afs.Path

	Workspace() afs.Path
	DotGit() afs.Path
	Repository() afs.Path
	Config() afs.Path
	Index() afs.Path
	HEAD() afs.Path
	Objects() afs.Path
	Refs() afs.Path
}

////////////////////////////////////////////////////////////////////////////////

// LayoutBuilder 仓库布局对象创建器
type LayoutBuilder struct {
	Config     afs.Path
	DotGit     afs.Path
	HEAD       afs.Path
	Index      afs.Path
	Objects    afs.Path
	Refs       afs.Path
	Repository afs.Path
	Workspace  afs.Path
}

// Create 创建仓库布局对象
func (inst *LayoutBuilder) Create() RepositoryLayout {

	res := &innerLayout{}

	res.config = inst.Config
	res.dotgit = inst.DotGit
	res.head = inst.HEAD
	res.index = inst.Index
	res.objects = inst.Objects
	res.refs = inst.Refs
	res.repo = inst.Repository
	res.workspace = inst.Workspace

	return res
}

////////////////////////////////////////////////////////////////////////////////

type innerLayout struct {
	pwd       afs.Path
	repo      afs.Path
	workspace afs.Path
	dotgit    afs.Path
	config    afs.Path
	index     afs.Path
	head      afs.Path
	objects   afs.Path
	refs      afs.Path
}

func (inst *innerLayout) _Impl() RepositoryLayout {
	return inst
}

func (inst *innerLayout) Workspace() afs.Path {
	return inst.workspace
}

func (inst *innerLayout) HEAD() afs.Path {
	return inst.head
}

func (inst *innerLayout) Objects() afs.Path {
	return inst.objects
}

func (inst *innerLayout) Refs() afs.Path {
	return inst.refs
}

func (inst *innerLayout) Index() afs.Path {
	return inst.index
}

func (inst *innerLayout) Config() afs.Path {
	return inst.config
}

func (inst *innerLayout) Repository() afs.Path {
	return inst.repo
}

func (inst *innerLayout) DotGit() afs.Path {
	return inst.dotgit
}

func (inst *innerLayout) WD() afs.Path {
	return inst.pwd
}