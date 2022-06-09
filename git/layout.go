package git

import "github.com/bitwormhole/gitlib/git/files"

// Layout 表示关于一个仓库的几个关键路径
type Layout interface {
	PWD() files.Path
	Workspace() files.Path
	DotGit() files.Path
	Repository() files.Path
	Config() files.Path
	Index() files.Path
	HEAD() files.Path
	Objects() files.Path
	Refs() files.Path
}

////////////////////////////////////////////////////////////////////////////////

// LayoutBuilder 仓库布局对象创建器
type LayoutBuilder struct {
	Config     files.Path
	DotGit     files.Path
	HEAD       files.Path
	Index      files.Path
	Objects    files.Path
	Refs       files.Path
	Repository files.Path
	Workspace  files.Path
}

// Create 创建仓库布局对象
func (inst *LayoutBuilder) Create() Layout {

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
	pwd       files.Path
	repo      files.Path
	workspace files.Path
	dotgit    files.Path
	config    files.Path
	index     files.Path
	head      files.Path
	objects   files.Path
	refs      files.Path
}

func (inst *innerLayout) _Impl() Layout {
	return inst
}

func (inst *innerLayout) Workspace() files.Path {
	return inst.workspace
}

func (inst *innerLayout) HEAD() files.Path {
	return inst.head
}

func (inst *innerLayout) Objects() files.Path {
	return inst.objects
}

func (inst *innerLayout) Refs() files.Path {
	return inst.refs
}

func (inst *innerLayout) Index() files.Path {
	return inst.index
}

func (inst *innerLayout) Config() files.Path {
	return inst.config
}

func (inst *innerLayout) Repository() files.Path {
	return inst.repo
}

func (inst *innerLayout) DotGit() files.Path {
	return inst.dotgit
}

func (inst *innerLayout) PWD() files.Path {
	return inst.pwd
}
