package repositories

import "github.com/starter-go/v1/afs"

type Layout interface {

	/////////////////////////////////////
	//files

	// the '.git' file (refer to repo)
	DotGitFile() afs.File

	// the '.git/config' file
	Config() afs.File

	// the '.git/description' file
	Description() afs.File

	// the '.git/HEAD' file
	HEAD() afs.File

	// the '.git/index' file
	Index() afs.File

	PackedRefs() afs.File

	/////////////////////////////////////
	//dirs

	// the '.git/config/..'
	Repository() afs.Directory

	// the '.git/'
	DotGitDir() afs.Directory

	// the '.git/objects'
	Objects() afs.Directory

	// the '.git/refs'
	Refs() afs.Directory

	// the '.git/worktrees'
	Worktrees() afs.Directory

	// the '.git/modules'
	Modules() afs.Directory

	Branches() afs.Directory
	Hooks() afs.Directory
	Info() afs.Directory
	Logs() afs.Directory

	// the '.git/..'
	CurrentWorktree() afs.Directory

	// the '.'
	WD() afs.Directory
}

type LayoutEntity struct {
}

////////////////////////////////////////////////////////////////////////////////

type LayoutBuilder struct {

	// .git

	DotGitDir      afs.Directory
	DotGitEndpoint afs.Directory // .git 文件指向的目标位置
	DotGitFile     afs.File

	// files

	Config      afs.File
	Description afs.File
	HEAD        afs.File
	Index       afs.File
	PackedRefs  afs.File

	// dirs

	Repository      afs.Directory
	CurrentWorktree afs.Directory
	WD              afs.Directory

	Objects   afs.Directory
	Refs      afs.Directory
	Worktrees afs.Directory
	Modules   afs.Directory
	Branches  afs.Directory
	Hooks     afs.Directory
	Info      afs.Directory
	Logs      afs.Directory
}

func (inst *LayoutBuilder) Build() Layout {

	l := new(innerLayout)

	// files

	l.config = inst.Config
	l.dotGitFile = inst.DotGitFile
	l.head = inst.HEAD
	l.index = inst.Index
	l.description = inst.Description
	l.packedRefs = inst.PackedRefs

	// dirs

	l.repository = inst.Repository
	l.dotGitDir = inst.DotGitDir
	l.objects = inst.Objects
	l.refs = inst.Refs
	l.worktrees = inst.Worktrees
	l.modules = inst.Modules
	l.currentWorktree = inst.CurrentWorktree
	l.wd = inst.WD
	l.logs = inst.Logs
	l.hooks = inst.Hooks
	l.branches = inst.Branches
	l.info = inst.Info

	return l
}

////////////////////////////////////////////////////////////////////////////////

type innerLayout struct {

	//files

	config      afs.File
	dotGitFile  afs.File
	head        afs.File
	index       afs.File
	description afs.File
	packedRefs  afs.File

	// dirs

	repository      afs.Directory
	dotGitDir       afs.Directory
	objects         afs.Directory
	refs            afs.Directory
	worktrees       afs.Directory
	modules         afs.Directory
	currentWorktree afs.Directory
	wd              afs.Directory
	branches        afs.Directory
	logs            afs.Directory
	hooks           afs.Directory
	info            afs.Directory
}

// Branches implements Layout.
func (i *innerLayout) Branches() afs.Directory {
	return i.branches
}

// Hooks implements Layout.
func (i *innerLayout) Hooks() afs.Directory {
	return i.hooks
}

// Info implements Layout.
func (i *innerLayout) Info() afs.Directory {
	return i.info
}

// Logs implements Layout.
func (i *innerLayout) Logs() afs.Directory {
	return i.logs
}

// PackedRefs implements Layout.
func (i *innerLayout) PackedRefs() afs.File {
	return i.packedRefs
}

// Config implements Layout.
func (i *innerLayout) Config() afs.File {
	return i.config
}

// CurrentWorktree implements Layout.
func (i *innerLayout) CurrentWorktree() afs.Directory {
	return i.currentWorktree
}

// DotGitDir implements Layout.
func (i *innerLayout) DotGitDir() afs.Directory {
	return i.dotGitDir
}

// DotGitFile implements Layout.
func (i *innerLayout) DotGitFile() afs.File {
	return i.dotGitFile
}

// DotGitFile implements Layout.
func (i *innerLayout) Description() afs.File {
	return i.description
}

// HEAD implements Layout.
func (i *innerLayout) HEAD() afs.File {
	return i.head
}

// Index implements Layout.
func (i *innerLayout) Index() afs.File {
	return i.index
}

// Modules implements Layout.
func (i *innerLayout) Modules() afs.Directory {
	return i.modules
}

// Objects implements Layout.
func (i *innerLayout) Objects() afs.Directory {
	return i.objects
}

// Refs implements Layout.
func (i *innerLayout) Refs() afs.Directory {
	return i.refs
}

// Repository implements Layout.
func (i *innerLayout) Repository() afs.Directory {
	return i.repository
}

// WD implements Layout.
func (i *innerLayout) WD() afs.Directory {
	return i.wd
}

// Worktrees implements Layout.
func (i *innerLayout) Worktrees() afs.Directory {
	return i.worktrees
}

////////////////////////////////////////////////////////////////////////////////
