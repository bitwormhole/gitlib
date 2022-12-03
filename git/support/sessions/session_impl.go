package sessions

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/data/gitfmt"
	"github.com/bitwormhole/gitlib/git/store"
)

type sessionImpl struct {
	repo store.Repository

	closelist []io.Closer
	pool      afs.ReaderPool
	packDao   store.PackDAO
	sparseDao store.SparseObjectLS

	tempFiles tempFileFactory
}

func (inst *sessionImpl) _Impl() store.Session {
	return inst
}

func (inst *sessionImpl) open() error {
	inst.tempFiles.init()
	return nil
}

func (inst *sessionImpl) Close() error {
	clist := inst.closelist
	inst.closelist = nil
	if clist == nil {
		return nil
	}
	errlist := make([]error, 0)
	for _, c := range clist {
		if c == nil {
			continue
		}
		err := c.Close()
		if err != nil {
			errlist = append(errlist, err)
		}
	}
	if len(errlist) > 0 {
		return errlist[0]
	}
	return nil
}

func (inst *sessionImpl) getSmallObjectSizeMax() int {
	return 8 * 1024 * 1024
}

func (inst *sessionImpl) GetRepository() store.Repository {
	return inst.repo
}

// 根据名称，取指定的组件
func (inst *sessionImpl) GetComponent(name string) (any, error) {

	return "", errors.New("no impl")
}

// 取工作目录
func (inst *sessionImpl) GetWD() afs.Path {
	return inst.repo.Layout().WD()
}

// 取仓库布局
func (inst *sessionImpl) GetLayout() store.RepositoryLayout {
	return inst.repo.Layout()
}

// objects

func (inst *sessionImpl) LoadText(id git.ObjectID) (string, *git.Object, error) {
	bin, o, err := inst.LoadBinary(id)
	if err != nil {
		return "", nil, err
	}
	return string(bin), o, nil
}

func (inst *sessionImpl) LoadBinary(id git.ObjectID) ([]byte, *git.Object, error) {
	o, reader, err := inst.ReadObject(id)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		reader.Close()
	}()
	limit := inst.getSmallObjectSizeMax()
	if o.Length > int64(limit) {
		const f = "the size of small object is over limit, size=%v limit=%v id=%v"
		return nil, nil, fmt.Errorf(f, o.Length, limit, id)
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}
	return data, o, nil
}

func (inst *sessionImpl) LoadCommit(id git.ObjectID) (*git.Commit, error) {
	text, o, err := inst.LoadText(id)
	if err != nil {
		return nil, err
	}
	if o.Type != git.ObjectTypeCommit {
		return nil, fmt.Errorf("bad object type, it's not a commit")
	}
	commit, err := gitfmt.ParseCommit(text)
	if err != nil {
		return nil, err
	}
	commit.ID = id
	return commit, nil
}

func (inst *sessionImpl) LoadTag(id git.ObjectID) (*git.Tag, error) {
	text, o, err := inst.LoadText(id)
	if err != nil {
		return nil, err
	}
	if o.Type != git.ObjectTypeTag {
		return nil, fmt.Errorf("bad object type, it's not a tag")
	}
	tag, err := gitfmt.ParseTag(text)
	if err != nil {
		return nil, err
	}
	tag.ID = id
	return tag, nil
}

func (inst *sessionImpl) LoadTree(id git.ObjectID) (*git.Tree, error) {
	bin, o, err := inst.LoadBinary(id)
	if err != nil {
		return nil, err
	}
	if o.Type != git.ObjectTypeTree {
		return nil, fmt.Errorf("bad object type, it's not a tree")
	}
	tree, err := gitfmt.ParseTree(bin)
	if err != nil {
		return nil, err
	}
	tree.ID = id
	return tree, nil
}

// HEAD ...
func (inst *sessionImpl) LoadHEAD(h store.HEAD) (*git.HEAD, error) {
	text, err := h.Path().GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return gitfmt.ParseHEAD(text)
}

// LoadRef ...
func (inst *sessionImpl) LoadRef(r store.Ref) (*git.Ref, error) {
	text, err := r.Path().GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return gitfmt.ParseRef(text)
}

func (inst *sessionImpl) ReadObject(id git.ObjectID) (*git.Object, io.ReadCloser, error) {
	so := inst.repo.Objects().GetSparseObject(id)
	if so.Exists() {
		spo := inst.GetSparseObjects()
		return spo.ReadSparseObject(so)
	}
	return inst.findAndReadObjectInPacks(id)
}

func (inst *sessionImpl) findAndReadObjectInPacks(id git.ObjectID) (*git.Object, io.ReadCloser, error) {
	// find in packs
	pii := &git.PackIndexItem{OID: id}
	packs := inst.GetPacks()
	pii, err := packs.FindPackObject(pii)
	if err != nil {
		return nil, nil, err
	}
	in, err := packs.ReadPackObject(pii)
	if err != nil {
		return nil, nil, err
	}
	return &git.Object{
		ID:     pii.OID,
		Type:   pii.Type,
		Length: pii.Length,
	}, in, nil
}

func (inst *sessionImpl) WriteObject(o *git.Object, data io.Reader) (*git.Object, error) {
	spo := inst.GetSparseObjects()
	return spo.WriteSparseObject(o, data)
}

func (inst *sessionImpl) GetPacks() store.PackDAO {
	dao1 := inst.packDao
	if dao1 != nil {
		return dao1
	}
	dao2 := &packDaoImpl{session: inst}
	dao2.init(32)
	dao1 = dao2
	inst.closelist = append(inst.closelist, dao2)
	inst.packDao = dao1
	return dao1
}

func (inst *sessionImpl) GetSparseObjects() store.SparseObjectLS {
	dao1 := inst.sparseDao
	if dao1 != nil {
		return dao1
	}
	dao2 := &sparseObjectsImpl{session: inst}
	dao2.init()
	dao1 = dao2
	inst.closelist = append(inst.closelist, dao2)
	inst.sparseDao = dao1
	return dao1
}

func (inst *sessionImpl) NewTemporaryFile(dir afs.Path) afs.Path {
	if dir == nil {
		dir = inst.repo.Layout().Objects().GetChild("info")
	}
	builder := inst.tempFiles.newBuilder()
	builder.dir = dir
	return builder.Create()
}

func (inst *sessionImpl) NewTemporaryBuffer(dir afs.Path) store.TemporaryBuffer {
	file := inst.NewTemporaryFile(dir)
	return &tempBuffer{
		tmpFile:   file,
		flushSize: 1024 * 1024,
	}
}

func (inst *sessionImpl) NewReaderPool(size int) afs.ReaderPool {
	return &afs.NopReaderPool{}
}

func (inst *sessionImpl) GetReaderPool() afs.ReaderPool {
	pool := inst.pool
	if pool != nil {
		return pool
	}
	pool = inst.NewReaderPool(32)
	inst.closelist = append(inst.closelist, pool)
	inst.pool = pool
	return pool
}
