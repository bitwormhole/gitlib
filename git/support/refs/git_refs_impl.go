package refs

import (
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

// GitRefsImpl ...
type GitRefsImpl struct {
	Core *store.Core

	//cache
	path afs.Path
}

func (inst *GitRefsImpl) _Impl() store.Refs {
	return inst
}

// Path ...
func (inst *GitRefsImpl) Path() afs.Path {
	p := inst.path
	if p != nil {
		return p
	}
	p = inst.Core.Layout.Refs()
	inst.path = p
	return p
}

// GetRef ...
func (inst *GitRefsImpl) GetRef(name git.ReferenceName) store.Ref {
	const prefix = "refs/"
	name = name.Normalize()
	href := name.String()
	if strings.HasPrefix(href, prefix) {
		href = href[len(prefix):]
	}
	dir := inst.Path()
	file := dir.GetChild(href)
	return &referenceImpl{name: name, path: file}
}

// List ...
func (inst *GitRefsImpl) List() []git.ReferenceName {
	f := finder{}
	return f.find(inst.path)
}
