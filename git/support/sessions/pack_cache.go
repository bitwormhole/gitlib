package sessions

import (
	"sort"
	"sync"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

// PackCache ...
type PackCache interface {
	Query(q *PackQuery) (ok bool)
}

// PackQuery ...
type PackQuery struct {

	// in
	Session store.Session
	PID     git.PackID
	OID     git.ObjectID

	// out
	ResultHolder *PackHolder
	ResultItem   *git.PackIndexItem
	Error        error

	// mid
	sn        PackQuerySN // 序列号，表示查询的次序
	tableDone map[PackHolderKey]*PackHolder
}

func (inst *PackQuery) accept(h *PackHolder, add2done bool) bool {
	if h == nil {
		return false
	}
	pid1 := inst.PID
	pid2 := h.pid
	key := h.key
	table := inst.tableDone
	if table == nil {
		table = make(map[PackHolderKey]*PackHolder)
		inst.tableDone = table
	}
	old := table[key]
	if add2done {
		table[key] = h
	}
	if old != nil {
		return false
	}
	if pid1 != nil && pid2 != nil {
		if !git.HashEqual(pid1, pid2) {
			return false
		}
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////

type packCacheL1 struct {
	next   PackCache
	sn     PackQuerySN
	cached *PackHolder
	mu     sync.Mutex
}

func (inst *packCacheL1) _Impl() PackCache {
	return inst
}

func (inst *packCacheL1) Query(q *PackQuery) bool {

	if q == nil {
		return false
	}

	inst.mu.Lock()
	defer func() {
		inst.mu.Unlock()
	}()

	inst.sn++
	q.sn = inst.sn

	// find by cache
	h1 := inst.cached
	if q.accept(h1, false) {
		ok := h1.query(q)
		if ok {
			return ok
		}
	}

	// find by next
	ok := inst.next.Query(q)
	if ok {
		h2 := q.ResultHolder
		if h2 != nil {
			inst.cached = h2
		}
	}
	return ok
}

////////////////////////////////////////////////////////////////////////////////

type packCacheL2 struct {
	next      PackCache
	cacheSize int
	cached    []*PackHolder
}

func (inst *packCacheL2) _Impl() PackCache {
	return inst
}

func (inst *packCacheL2) Query(q *PackQuery) bool {

	// find by cache
	sort.Sort(inst) // warn todo: 这里目前可能需要调整排序方向
	list := inst.cached
	for _, h := range list {
		if q.accept(h, false) {
			ok := h.query(q)
			if ok {
				return ok
			}
		}
	}

	// find by next
	ok := inst.next.Query(q)
	if ok {
		h2 := q.ResultHolder
		if h2 != nil {
			inst.cached = append(inst.cached, h2)
		}
	}
	return ok
}

func (inst *packCacheL2) Len() int {
	list := inst.cached
	if list == nil {
		return 0
	}
	return len(list)
}
func (inst *packCacheL2) Less(a, b int) bool {
	list := inst.cached
	aa := list[a]
	bb := list[b]
	return aa.hitAt < bb.hitAt
}
func (inst *packCacheL2) Swap(a, b int) {
	list := inst.cached
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////

type packCacheL3 struct {
	// nextLayer PackCache
}

func (inst *packCacheL3) _Impl() PackCache {
	return inst
}

func (inst *packCacheL3) Query(q *PackQuery) bool {
	pid := q.PID
	if pid == nil {
		return inst.scanAll(q)
	}
	return inst.findOne(pid, q)
}

func (inst *packCacheL3) scanAll(q *PackQuery) bool {
	// find all in disc
	session := q.Session
	repo := session.GetRepository()
	objs := repo.Objects()
	ids := objs.ListPacks()
	for _, pid := range ids {
		ok := inst.findOne(pid, q)
		if ok {
			return ok
		}
	}
	return false
}

func (inst *packCacheL3) findOne(pid git.PackID, q *PackQuery) bool {
	session := q.Session
	repo := session.GetRepository()
	objs := repo.Objects()
	p := objs.GetPack(pid)
	if !p.Exists() {
		return false
	}
	h, err := inst.initPack(p, session)
	if err != nil {
		return false
	}
	if !q.accept(h, false) {
		return false
	}
	err = h.load()
	if err != nil {
		return false
	}
	return h.query(q)
}

func (inst *packCacheL3) initPack(p store.Pack, s store.Session) (*PackHolder, error) {
	h := &PackHolder{}
	err := h.init(p, s)
	if err != nil {
		return nil, err
	}
	return h, nil
}

////////////////////////////////////////////////////////////////////////////////

// NewPackCacheChain ...
func NewPackCacheChain(size int) PackCache {
	l1 := &packCacheL1{}
	l2 := &packCacheL2{cacheSize: size}
	l3 := &packCacheL3{}
	l1.next = l2
	l2.next = l3
	return l1
}
