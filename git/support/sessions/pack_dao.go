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

	cacheSize int
	cache     map[string]*packPairHolder // map[ stringify(pid) ]
}

func (inst *packDaoImpl) _Impl() store.PackDAO {
	return inst
}

func (inst *packDaoImpl) init(size int) error {
	inst.cache = make(map[string]*packPairHolder)
	inst.cacheSize = size
	return nil
}

func (inst *packDaoImpl) Close() error {
	return nil
}

func (inst *packDaoImpl) FindPackObject(o *git.PackIndexItem) (*git.PackIndexItem, error) {
	panic("no impl : packDaoImpl.FindPackObject")
}

func (inst *packDaoImpl) ReadPackObject(o *git.PackIndexItem) (io.ReadCloser, error) {
	panic("no impl : packDaoImpl.ReadPackObject")
}

func (inst *packDaoImpl) CheckPack(pid git.PackID, flags pack.CheckFlag) error {
	holder, err := inst.getPack(pid)
	if err != nil {
		return err
	}

	idx := holder.idx
	ent := holder.entity

	err = idx.Check(flags)
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

func (inst *packDaoImpl) getPack(pid git.PackID) (*packPairHolder, error) {
	if pid == nil {
		return nil, fmt.Errorf("param:pid is nil")
	}
	key := pid.String()
	table := inst.cache
	holder := table[key]
	if holder != nil {
		return holder, nil
	}
	// make new holder
	repo := inst.session.GetRepository()
	p := repo.Objects().GetPack(pid)
	if !p.Exists() {
		return nil, fmt.Errorf("no pack with id:%v", key)
	}
	holder = &packPairHolder{}
	err := holder.init(p, inst.session)
	if err != nil {
		return nil, err
	}
	table[key] = holder
	return holder, nil
}

////////////////////////////////////////////////////////////////////////////////

type packPairHolder struct {
	pid git.PackID

	session store.Session

	idx    pack.Idx
	entity pack.Pack
}

func (inst *packPairHolder) init(p store.Pack, session store.Session) error {

	if p == nil || session == nil {
		return fmt.Errorf("param is nil")
	}

	pool := session.GetReaderPool()
	digest := session.GetRepository().Digest()
	pathIdx := p.GetIndexFile()
	pathPack := p.GetEntityFile()

	// for .idx
	idx, err := pack.NewIdx(&pack.File{
		Pool:   pool,
		Path:   pathIdx,
		Type:   pack.FileTypeIdx,
		Digest: digest,
	})
	if err != nil {
		return err
	}

	// for .pack
	ent, err := pack.NewPack(&pack.File{
		Pool:   pool,
		Path:   pathPack,
		Type:   pack.FileTypePack,
		Digest: digest,
	})
	if err != nil {
		return err
	}

	// done
	inst.idx = idx
	inst.entity = ent
	inst.session = session
	inst.pid = p.GetID()
	return nil
}
