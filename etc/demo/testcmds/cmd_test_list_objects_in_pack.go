package testcmds

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

type tagListObjectsInPackItem struct {
	i git.PackIndexItem
	p git.PackedObjectHeaderEx
}

// TestListObjectsInPack ...
type TestListObjectsInPack struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`

	all       []*tagListObjectsInPackItem
	countRead int64
}

func (inst *TestListObjectsInPack) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestListObjectsInPack) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-list-objects-in-pack",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestListObjectsInPack) run(task *cli.Task) error {

	lib, err := inst.LA.GetLib()
	if err != nil {
		return err
	}

	wd := lib.FS().NewPath(inst.WD)
	repo, err := lib.RepositoryLoader().LoadWithPath(wd)
	if err != nil {
		return err
	}

	session, err := repo.OpenSession()
	if err != nil {
		return err
	}
	defer func() { session.Close() }()

	plist := repo.Objects().ListPacks()
	for _, pid := range plist {
		err := inst.loadObjectsInPack(pid, session)
		if err != nil {
			return err
		}
	}

	inst.log(session)

	return nil
}

func (inst *TestListObjectsInPack) loadObjectsInPack(pid git.PackID, session store.Session) error {

	vlog.Info("print objects info in pack:", pid.String())

	repo := session.GetRepository()
	packStore := repo.Objects().GetPack(pid)
	idx, pack, err := inst.openPack(packStore, session)
	if err != nil {
		return err
	}

	count := idx.Count()
	all, err := idx.GetItems(0, int(count))
	if err != nil {
		return err
	}

	for _, item1 := range all {
		item2 := pack.IndexToHeader(item1)
		hx, err := inst.readObjectInfo(pack, item2)
		if err != nil {
			return err
		}
		inst.add(item1, hx)
	}

	inst.sort()

	return nil
}

func (inst *TestListObjectsInPack) readObjectInfo(pack pack.Pack, item *git.PackedObjectHeaderEx) (*git.PackedObjectHeaderEx, error) {
	hx, in, err := pack.OpenObjectReader(item, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		in.Close()
	}()
	count := int64(0)
	buf := make([]byte, 1024)
	for {
		cb, err := in.Read(buf)
		if cb > 0 {
			count += int64(cb)
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}
	// hx.RealSize = count
	if count != hx.Size {
		return nil, fmt.Errorf("bad entity size")
	}
	return hx, nil
}

func (inst *TestListObjectsInPack) log(session store.Session) {

	sb := strings.Builder{}
	for i, item := range inst.all {
		sb.Reset()
		sb.WriteString(fmt.Sprintf("%v. ", i))
		sb.WriteString(item.i.OID.String())

		sb.WriteString(fmt.Sprintf(" index:%v", item.i.Index))
		sb.WriteString(fmt.Sprintf(" offset:%v", item.i.Offset))
		sb.WriteString(fmt.Sprintf(" length:%v", item.p.Size))
		// sb.WriteString(fmt.Sprintf(" length-real:%v", item.p.RealSize))

		sb.WriteString(fmt.Sprintf(" type:%v", item.p.Type.ToObjectType().String()))

		if item.p.Type == git.PackedDeltaOFS {
			abs := item.p.GetDeltaParentOffset()
			sb.WriteString(fmt.Sprintf(" deltaOffset:%v", item.p.DeltaOffset))
			sb.WriteString(fmt.Sprintf(" deltaOffset_abs:%v", abs))

		} else if item.p.Type == git.PackedDeltaRef {
		} else if item.p.Type == git.PackedCommit {
		}

		fmt.Println(sb.String())
	}

	fmt.Printf("count-read: %v bytes\n", inst.countRead)
}

func (inst *TestListObjectsInPack) add(itemIdx *git.PackIndexItem, itemPack *git.PackedObjectHeaderEx) {
	item := &tagListObjectsInPackItem{}
	item.i = *itemIdx
	item.p = *itemPack
	inst.all = append(inst.all, item)
}

func (inst *TestListObjectsInPack) openPack(p store.Pack, session store.Session) (pack.Idx, pack.Pack, error) {

	pool := session.GetReaderPool()
	repo := session.GetRepository()
	compr := repo.Compression()
	digest := repo.Digest()

	idx, err := pack.NewIdx(&pack.File{
		Compression: compr,
		Digest:      digest,
		Path:        p.GetDotIdx(), //IndexFile(),
		Pool:        pool,
		Type:        pack.FileTypeIdx,
	})
	if err != nil {
		return nil, nil, err
	}

	pk, err := pack.NewPack(&pack.File{
		Compression: compr,
		Digest:      digest,
		Path:        p.GetDotPack(), // EntityFile(),
		Pool:        pool,
		Type:        pack.FileTypePack,
	})
	if err != nil {
		return nil, nil, err
	}

	return idx, pk, nil
}

func (inst *TestListObjectsInPack) sort() {
	sort.Sort(inst)
}
func (inst *TestListObjectsInPack) Len() int {
	return len(inst.all)
}
func (inst *TestListObjectsInPack) Less(i1, i2 int) bool {
	o1 := inst.all[i1]
	o2 := inst.all[i2]
	return o1.i.Offset < o2.i.Offset
}
func (inst *TestListObjectsInPack) Swap(i1, i2 int) {
	o1 := inst.all[i1]
	o2 := inst.all[i2]
	inst.all[i1] = o2
	inst.all[i2] = o1
}
