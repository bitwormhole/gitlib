package sessions

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
)

type packDaoImpl struct {
	session store.Session
	cache   PackCache
}

func (inst *packDaoImpl) _Impl() store.Packs {
	return inst
}

func (inst *packDaoImpl) init(size int) error {
	inst.cache = NewPackCacheChain(size)
	return nil
}

func (inst *packDaoImpl) Close() error {
	return nil
}

func (inst *packDaoImpl) FindPackObject(oid git.ObjectID) (store.PackObject, error) {
	return inst.find(nil, oid)
}

func (inst *packDaoImpl) FindPackObjectInPack(pid git.PackID, oid git.ObjectID) (store.PackObject, error) {
	return inst.find(pid, oid)
}

// pid 是可选的, oid 必须有
func (inst *packDaoImpl) find(pid git.PackID, oid git.ObjectID) (store.PackObject, error) {

	if oid == nil {
		return nil, fmt.Errorf("required param:oid is nil")
	}

	q := &PackQuery{
		PID:     pid,
		OID:     oid,
		Session: inst.session,
	}

	ok := inst.cache.Query(q)
	item := q.ResultItem
	if !ok || item == nil {
		oidstr := "nil"
		if oid != nil {
			oidstr = oid.String()
		}
		return nil, fmt.Errorf("no wanted item [object.id:%v]", oidstr)
	}

	pid = item.PID
	result := inst.session.GetRepository().Objects().GetPack(pid).GetObject(oid)
	return result, nil
}

func (inst *packDaoImpl) ReadPackObject(po store.PackObject) (*git.Object, io.ReadCloser, error) {

	session := inst.session
	pid := po.Container().GetID()
	oid := po.GetID()
	q := &PackQuery{
		Session: session,
		PID:     pid,
		OID:     oid,
	}
	ok := inst.cache.Query(q)
	err := q.Error
	if !ok {
		if err == nil {
			err = fmt.Errorf("in-pack object not found")
		}
	}
	if err != nil {
		return nil, nil, err
	}

	objCtx := session.GetObjectContext()
	packCtx := objCtx.NewPackContext(pid)
	packFile := po.Container().GetDotPack()
	cp, err := pack.NewComplexPack(&pack.File{
		Context: packCtx,
		Type:    pack.FileTypePack,
		Path:    packFile,
	})
	if err != nil {
		return nil, nil, err
	}

	hx := cp.IndexToHeader(q.ResultItem)
	hx, reader, err := cp.OpenComplexObjectReader(hx, nil)
	if err != nil {
		return nil, nil, err
	}

	o := &git.Object{
		Type:   hx.Type.ToObjectType(),
		Length: hx.Size,
		ID:     oid,
	}
	return o, reader, nil
}

// func (inst *packDaoImpl) ReadPackObject(po store.PackObject) (*git.Object, io.ReadCloser, error) {
// 	if po == nil {
// 		return nil, nil, fmt.Errorf("param is nil")
// 	}
// 	pid := po.Container().GetID()
// 	oid := po.GetID()
// 	if pid == nil {
// 		return nil, nil, fmt.Errorf("pid is nil")
// 	}
// 	if oid == nil {
// 		return nil, nil, fmt.Errorf("oid is nil")
// 	}
// 	readerbuilder := packEntityReaderBuilder{
// 		session: inst.session,
// 		wantPID: pid,
// 		wantOID: oid,
// 		pc:      inst.cache,
// 	}
// 	hx, in2, err := readerbuilder.open()
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	resultObject := &git.Object{
// 		ID:     po.GetID(),
// 		Type:   hx.Type.ToObjectType(),
// 		Length: hx.Size,
// 	}
// 	return resultObject, in2, nil
// }

func (inst *packDaoImpl) CheckPack(pid git.PackID, flags pack.CheckFlag) error {

	q := &PackQuery{
		PID:     pid,
		Session: inst.session,
	}
	ok := inst.cache.Query(q)
	if !ok {
		return fmt.Errorf("pack.object not found")
	}
	holder := q.ResultHolder
	if holder == nil {
		return fmt.Errorf("pack-holder is nil")
	}

	idx := holder.idx
	ent := holder.pack

	err := idx.Check(flags)
	if err != nil {
		return err
	}

	err = ent.Check(flags)
	if err != nil {
		return err
	}

	return nil
}

func (inst *packDaoImpl) ListPacks() ([]git.PackID, error) {
	repo := inst.session.GetRepository()
	list := repo.Objects().ListPacks()
	return list, nil
}

func (inst *packDaoImpl) NewPackBuilder() store.PackBuilder {
	panic("no impl : packDaoImpl.NewPackBuilder")
}

func (inst *packDaoImpl) ImportPack(p *store.ImportPackParams) (*store.ImportPackResult, error) {
	imp := packImporter{session: inst.session}
	return imp.Import(p)
}

////////////////////////////////////////////////////////////////////////////////
