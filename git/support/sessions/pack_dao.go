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

func (inst *packDaoImpl) _Impl() store.PackDAO {
	return inst
}

func (inst *packDaoImpl) init(size int) error {
	inst.cache = NewPackCacheChain(size)
	return nil
}

func (inst *packDaoImpl) Close() error {
	return nil
}

func (inst *packDaoImpl) FindPackObject(o *git.PackIndexItem) (*git.PackIndexItem, error) {

	q := &PackQuery{
		PID:     o.PID,
		OID:     o.OID,
		Session: inst.session,
	}

	ok := inst.cache.Query(q)
	item := q.ResultItem
	if !ok || item == nil {
		oid := o.OID
		oidstr := "nil"
		if oid != nil {
			oidstr = oid.String()
		}
		return nil, fmt.Errorf("no wanted item [object.id:%v]", oidstr)
	}

	return item, nil
}

func (inst *packDaoImpl) ReadPackObject(pii *git.PackIndexItem) (io.ReadCloser, error) {
	if pii == nil {
		return nil, fmt.Errorf("param is nil")
	}
	pid := pii.PID
	if pid == nil {
		return nil, fmt.Errorf("pid is nil")
	}
	q := &PackQuery{PID: pid}
	ok := inst.cache.Query(q)
	if !ok {
		return nil, fmt.Errorf("no pack with id:%v", pid.String())
	}
	holder := q.ResultHolder
	pool := inst.session.GetReaderPool()
	file := holder.pack.GetEntityFile()
	in1, err := pool.OpenReader(file, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if in1 != nil {
			in1.Close()
		}
	}()
	readerbuilder := packEntityReaderBuilder{
		input:   in1,
		file:    file,
		index:   pii,
		session: inst.session,
	}
	in2, err := readerbuilder.open()
	if err != nil {
		return nil, err
	}
	pii.Type = readerbuilder.entityType.ToObjectType()
	pii.PackedType = readerbuilder.entityType
	pii.Length = readerbuilder.entityLength
	in1 = nil
	return in2, nil
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
