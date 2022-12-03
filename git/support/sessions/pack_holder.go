package sessions

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
)

// PackHolderKey ...
type PackHolderKey string

// PackQuerySN  序列号，表示查询的次序
type PackQuerySN int64

// PackHolder ...
type PackHolder struct {
	key     PackHolderKey
	pid     git.PackID
	pack    store.Pack
	session store.Session
	idx     pack.Idx
	entity  pack.Pack
	hitAt   PackQuerySN
}

// FindObject ...
func (inst *PackHolder) query(q *PackQuery) (ok bool) {
	if q == nil {
		return false
	}
	if !q.accept(inst, true) {
		return false
	}
	oid := q.OID
	pid := q.PID
	if oid != nil {
		return inst.findObject(oid, q)
	} else if pid != nil {
		return inst.findPack(pid, q)
	}
	return false
}

func (inst *PackHolder) findPack(want git.PackID, q *PackQuery) (ok bool) {
	have := inst.pid
	if !git.HashEqual(want, have) {
		return false
	}
	inst.hitAt = q.sn
	q.ResultHolder = inst
	return true
}

func (inst *PackHolder) findObject(want git.ObjectID, q *PackQuery) (ok bool) {
	item, err := inst.idx.Find(want)
	if err != nil {
		return false
	}
	q.ResultItem = item
	q.ResultHolder = inst
	inst.hitAt = q.sn
	return true
}

func (inst *PackHolder) init(p store.Pack, session store.Session) error {
	if p == nil || session == nil {
		return fmt.Errorf("param is nil")
	}
	pid := p.GetID()
	if pid == nil {
		return fmt.Errorf("pack.pid is nil")
	}
	inst.session = session
	inst.pid = pid
	inst.key = PackHolderKey(pid.String())
	inst.pack = p
	return nil
}

func (inst *PackHolder) load() error {

	session := inst.session
	p := inst.pack

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
	return nil
}
