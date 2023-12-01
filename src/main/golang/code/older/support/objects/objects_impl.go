package objects

import (
	"strings"

	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
)

// GitObjectsImpl ...
type GitObjectsImpl struct {
	Core *store.Core

	// cache
	path    afs.Path
	pathMap git.PathMapping
}

func (inst *GitObjectsImpl) _Impl() store.Objects {
	return inst
}

func (inst *GitObjectsImpl) getPathMap() git.PathMapping {
	pm := inst.pathMap
	if pm != nil {
		return pm
	}
	pm = inst.Core.PathMapping
	inst.pathMap = pm
	return pm
}

// Path ...
func (inst *GitObjectsImpl) Path() afs.Path {
	p := inst.path
	if p != nil {
		return p
	}
	layout := inst.Core.Layout
	p = layout.Objects()
	inst.path = p
	return p
}

// GetSparseObject ...
func (inst *GitObjectsImpl) GetSparseObject(oid git.ObjectID) store.SparseObject {
	base := inst.Path()
	pm := inst.getPathMap()
	file := pm.Map(base, oid)
	return &sparseObjectImpl{
		file: file,
		id:   oid,
	}
}

func (inst *GitObjectsImpl) getPackDir() afs.Path {
	base := inst.Path()
	return base.GetChild("pack")
}

// GetPack ...
func (inst *GitObjectsImpl) GetPack(pid git.PackID) store.Pack {
	const (
		fileNamePrefix     = "pack-"
		fileNameSuffixPack = ".pack"
		fileNameSuffixIdx  = ".idx"
	)
	idstr := "0000"
	if pid != nil {
		idstr = pid.String()
	}
	dir := inst.getPackDir()
	packFile := dir.GetChild(fileNamePrefix + idstr + fileNameSuffixPack)
	idxFile := dir.GetChild(fileNamePrefix + idstr + fileNameSuffixIdx)
	p := &packImpl{
		id:             pid,
		packIndexFile:  idxFile,
		packEntityFile: packFile,
	}
	return p
}

// Info ...
func (inst *GitObjectsImpl) Info() store.InfoFolder {
	panic("no impl")
}

// ListPacks ...
func (inst *GitObjectsImpl) ListPacks() []git.PackID {
	const (
		fileNamePrefix = "pack-"
		// fileNameSuffixPack = ".pack"
		fileNameSuffixIdx = ".idx"
	)
	dst := make([]git.PackID, 0)
	dir := inst.getPackDir()
	namelist := dir.ListNames()
	for _, name := range namelist {
		if strings.HasPrefix(name, fileNamePrefix) && strings.HasSuffix(name, fileNameSuffixIdx) {
			idstr := name[len(fileNamePrefix) : len(name)-len(fileNameSuffixIdx)]
			id, err := git.ParsePackID(idstr)
			if err != nil {
				vlog.Warn(err)
			} else if id != nil {
				dst = append(dst, id)
			}
		}
	}
	return dst
}
