package testcmds

import (
	"fmt"
	"sort"
	"strings"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/starter-go/afs"
	"github.com/starter-go/vlog"
)

// TestGenIdxForPack ...
type TestGenIdxForPack struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`

	idxItems []*git.PackIndexItem
}

func (inst *TestGenIdxForPack) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestGenIdxForPack) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-gen-idx-for-pack",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestGenIdxForPack) run(task *cli.Task) error {

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

	// dao := session.GetPacks()
	plist := repo.Objects().ListPacks()
	for _, pid := range plist {
		err = inst.genIdx(session, pid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *TestGenIdxForPack) genIdx(session store.Session, pid git.PackID) error {

	repo := session.GetRepository()
	pack1 := repo.Objects().GetPack(pid)
	path := pack1.GetDotPack()

	objCtx := session.GetObjectContext()
	packCtx := objCtx.NewPackContext(pid)

	// 打开pack
	pack2, err := pack.NewComplexPack(&pack.File{
		Context: packCtx,
		Path:    path,
		Type:    pack.FileTypePack,
	})
	if err != nil {
		return err
	}

	// 加载作为参考的idx
	err = inst.loadExampleIdx(packCtx, pack1.GetDotIdx())
	if err != nil {
		return err
	}

	// 生成新的 idx
	idx := pack1.GetDotIdx()
	idx = inst.getTempIdxFile(idx)
	err = pack2.MakeIdx(idx)
	if err != nil {
		return err
	}

	return nil
}

func (inst *TestGenIdxForPack) getTempIdxFile(file afs.Path) afs.Path {
	dir := file.GetParent()
	name := file.GetName()
	return dir.GetChild(name + ".tmp")
}

func (inst *TestGenIdxForPack) loadExampleIdx(ctx *objects.PackContext, path afs.Path) error {

	idx, err := pack.NewIdx(&pack.File{
		Context: ctx,
		Path:    path,
		Type:    pack.FileTypeIdx,
	})
	if err != nil {
		return err
	}

	items, err := idx.GetItems(0, 9999)
	if err != nil {
		return err
	}
	inst.sort(items)
	count := 0

	for _, item := range items {
		if count > 9 {
			break
		}
		sb := strings.Builder{}
		sb.WriteString("[pack-idx-item ")
		sb.WriteString(fmt.Sprintf(" index:%v", item.Index))
		sb.WriteString(fmt.Sprintf(" oid:%v", item.OID.String()))
		sb.WriteString(fmt.Sprintf(" offset:%v", item.Offset))
		sb.WriteString(fmt.Sprintf(" crc32:%v", item.CRC32))
		sb.WriteString("]")
		vlog.Warn(sb.String())
		count++
	}

	return nil
}

func (inst *TestGenIdxForPack) sort(items []*git.PackIndexItem) {
	inst.idxItems = items
	sort.Sort(inst)
}
func (inst *TestGenIdxForPack) Len() int {
	return len(inst.idxItems)
}
func (inst *TestGenIdxForPack) Less(i1, i2 int) bool {
	list := inst.idxItems
	o1 := list[i1]
	o2 := list[i2]
	return o1.Offset < o2.Offset
}
func (inst *TestGenIdxForPack) Swap(i1, i2 int) {
	list := inst.idxItems
	o1 := list[i1]
	o2 := list[i2]
	list[i1] = o2
	list[i2] = o1
}
