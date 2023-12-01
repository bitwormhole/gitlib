package testcmds

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// type tagListObjectsInPackItem struct {
// 	i git.PackIndexItem
// 	p git.PackedObjectHeaderEx
// }

// TestPackDeltaObjects ...
type TestPackDeltaObjects struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`

	all       []*tagListObjectsInPackItem
	countRead int64
}

func (inst *TestPackDeltaObjects) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestPackDeltaObjects) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-pack-delta-objects",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestPackDeltaObjects) run(task *cli.Task) error {

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

	return nil
}

func (inst *TestPackDeltaObjects) loadObjectsInPack(pid git.PackID, session store.Session) error {

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
	inst.log(session)

	return nil
}

func (inst *TestPackDeltaObjects) readObjectInfo(pack pack.Pack, item *git.PackedObjectHeaderEx) (*git.PackedObjectHeaderEx, error) {
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

func (inst *TestPackDeltaObjects) log(session store.Session) {

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
			_, err := session.GetPacks().FindPackObject(item.i.OID)
			if err != nil {
				vlog.Error(err)
			}
		}

		err := inst.parseDeltaRawContent(item, session)
		if err != nil {
			vlog.Error(err)
		}

		fmt.Println(sb.String())
	}

	fmt.Printf("count-read: %v bytes\n", inst.countRead)
}

func (inst *TestPackDeltaObjects) parseDeltaRawContent(item *tagListObjectsInPackItem, session store.Session) error {

	pid := item.p.PID
	// oid := item.p.OID
	repo := session.GetRepository()
	p1 := repo.Objects().GetPack(pid)

	objCtx := session.GetObjectContext()
	packCtx := objCtx.NewPackContext(pid)

	p2, err := pack.NewPack(&pack.File{
		Context: packCtx,
		Path:    p1.GetDotPack(), //  EntityFile(),
		Type:    pack.FileTypePack,
	})
	if err != nil {
		return err
	}

	for i := 5; i > 0; i-- {
		hx, in, err := p2.OpenObjectReader(&item.p, nil)
		if err != nil {
			return err
		}
		defer func() {
			in.Close()
		}()
		vlog.Debug("offset:", hx.Offset)
		data, _ := ioutil.ReadAll(in)
		count := len(data)
		inst.countRead += int64(count)
	}

	return nil
}

func (inst *TestPackDeltaObjects) add(itemIdx *git.PackIndexItem, itemPack *git.PackedObjectHeaderEx) {
	item := &tagListObjectsInPackItem{}
	item.i = *itemIdx
	item.p = *itemPack
	inst.all = append(inst.all, item)
}

func (inst *TestPackDeltaObjects) openPack(p store.Pack, session store.Session) (pack.Idx, pack.Pack, error) {

	pid := p.GetID()
	objCtx := session.GetObjectContext()
	packCtx := objCtx.NewPackContext(pid)

	idx, err := pack.NewIdx(&pack.File{
		Context: packCtx,
		Path:    p.GetDotIdx(), //IndexFile(),
		Type:    pack.FileTypeIdx,
	})
	if err != nil {
		return nil, nil, err
	}

	pk, err := pack.NewPack(&pack.File{
		Context: packCtx,
		Path:    p.GetDotPack(), // EntityFile(),
		Type:    pack.FileTypePack,
	})
	if err != nil {
		return nil, nil, err
	}

	return idx, pk, nil
}

func (inst *TestPackDeltaObjects) sort() {
	sort.Sort(inst)
}
func (inst *TestPackDeltaObjects) Len() int {
	return len(inst.all)
}
func (inst *TestPackDeltaObjects) Less(i1, i2 int) bool {
	o1 := inst.all[i1]
	o2 := inst.all[i2]
	return o1.i.Index < o2.i.Index
}
func (inst *TestPackDeltaObjects) Swap(i1, i2 int) {
	o1 := inst.all[i1]
	o2 := inst.all[i2]
	inst.all[i1] = o2
	inst.all[i2] = o1
}
