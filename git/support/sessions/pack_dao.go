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
	if po == nil {
		return nil, nil, fmt.Errorf("param is nil")
	}
	pid := po.Container().GetID()
	if pid == nil {
		return nil, nil, fmt.Errorf("pid is nil")
	}
	q := &PackQuery{PID: pid}
	ok := inst.cache.Query(q)
	if !ok {
		return nil, nil, fmt.Errorf("no pack with id:%v", pid.String())
	}
	holder := q.ResultHolder
	readerbuilder := packEntityReaderBuilder{
		session: inst.session,
		pack:    holder.entity,
		want:    q.ResultItem,
	}
	hx, in2, err := readerbuilder.open()
	if err != nil {
		return nil, nil, err
	}
	resultObject := &git.Object{
		ID:     po.GetID(),
		Type:   hx.Type.ToObjectType(),
		Length: hx.Size,
	}
	return resultObject, in2, nil
}

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
	ent := holder.entity

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
