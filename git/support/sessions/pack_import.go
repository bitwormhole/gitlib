package sessions

import (
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
)

type packImporter struct {
	session store.Session
}

func (inst *packImporter) Import(p *store.ImportPackParams) (*store.ImportPackResult, error) {

	r1, err := inst.checkImportPack(p)
	if err != nil {
		return nil, err
	}

	pid := r1.ID
	pack := r1.Pack
	idxDst := pack.GetDotIdx()
	packDst := pack.GetDotPack()
	idxSrc := p.IdxFile
	packSrc := p.PackFile
	errlist := make([]error, 0)

	// copy or move
	if p.MoveFiles {
		e1 := idxSrc.MoveTo(idxDst, nil)
		e2 := packSrc.MoveTo(packDst, nil)
		errlist = append(errlist, e1)
		errlist = append(errlist, e2)
	} else {
		op := &afs.Options{Create: true, Mkdirs: true}
		e1 := idxSrc.CopyTo(idxDst, op)
		e2 := packSrc.CopyTo(packDst, op)
		errlist = append(errlist, e1)
		errlist = append(errlist, e2)
	}

	for _, err := range errlist {
		if err != nil {
			return nil, err
		}
	}
	result := &store.ImportPackResult{
		Params: p,
		ID:     pid,
		Pack:   pack,
	}
	return result, nil

}

func (inst *packImporter) checkImportPack(p *store.ImportPackParams) (*store.ImportPackResult, error) {

	// check params

	if p == nil {
		return nil, fmt.Errorf("param:ImportPackParams is nil")
	}

	idxSrc := p.IdxFile
	packSrc := p.PackFile
	pid := p.ID

	if idxSrc == nil {
		return nil, fmt.Errorf("param:ImportPackParams.IdxFile is nil")
	}
	if packSrc == nil {
		return nil, fmt.Errorf("param:ImportPackParams.PackFile is nil")
	}

	// check files
	pid1, err := inst.checkPackFile(packSrc)
	if err != nil {
		return nil, err
	}

	pid2, err := inst.checkPackIdxFile(idxSrc)
	if err != nil {
		return nil, err
	}

	if git.HashEqual(pid1, pid2) {
		pid = pid1
	} else {
		s1 := pid1.String()
		s2 := pid2.String()
		return nil, fmt.Errorf("pid_1 != pid_2, pid_1:%v pid_2:%v", s1, s2)
	}

	repo := inst.session.GetRepository()
	pack := repo.Objects().GetPack(pid)

	result := &store.ImportPackResult{
		Params: p,
		ID:     pid,
		Pack:   pack,
	}
	return result, nil
}

func (inst *packImporter) checkPackIdxFile(file afs.Path) (git.PackID, error) {
	objCtx := inst.session.GetObjectContext()
	packCtx := objCtx.NewPackContext(nil)
	idx, err := pack.NewIdx(&pack.File{
		Context: packCtx,
		Path:    file,
		Type:    pack.FileTypeIdx,
	})
	if err != nil {
		return nil, err
	}
	err = idx.Check(pack.CheckAll)
	if err != nil {
		return nil, err
	}
	pid := idx.GetPackID()
	if pid == nil {
		return nil, fmt.Errorf("pid is nil")
	}
	return pid, nil
}

func (inst *packImporter) checkPackFile(file afs.Path) (git.PackID, error) {
	objCtx := inst.session.GetObjectContext()
	packCtx := objCtx.NewPackContext(nil)
	pp, err := pack.NewPack(&pack.File{
		Context: packCtx,
		Path:    file,
		Type:    pack.FileTypePack,
	})
	if err != nil {
		return nil, err
	}
	err = pp.Check(pack.CheckAll)
	if err != nil {
		return nil, err
	}
	pid := pp.GetPackID()
	if pid == nil {
		return nil, fmt.Errorf("pid is nil")
	}
	return pid, nil
}
