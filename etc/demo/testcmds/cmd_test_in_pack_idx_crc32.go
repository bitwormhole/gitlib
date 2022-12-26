package testcmds

import (
	"fmt"
	"hash/crc32"
	"io"
	"sort"
	"strconv"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// TestPackIdxCRC32 ...
type TestPackIdxCRC32 struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`

	items []*git.PackIndexItem
}

func (inst *TestPackIdxCRC32) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestPackIdxCRC32) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-pack-idx-crc32",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestPackIdxCRC32) run(task *cli.Task) error {

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
		err := inst.tryPack(pid, session)
		if err != nil {
			return err
		}
	}

	// inst.log(session)

	return nil
}

func (inst *TestPackIdxCRC32) tryPack(pid git.PackID, session store.Session) error {

	repo := session.GetRepository()
	pk := session.GetRepository().Objects().GetPack(pid)
	ctx1 := &objects.Context{
		Compression: repo.Compression(),
		Digest:      repo.Digest(),
		Pool:        session.GetReaderPool(),
	}
	ctx2 := ctx1.NewPackContext(pid)

	pack1, err := pack.NewComplexPack(&pack.File{
		Context: ctx2,
		Type:    pack.FileTypePack,
		Path:    pk.GetDotPack(),
	})
	if err != nil {
		return err
	}
	if pack1 != nil {
	}

	idx1, err := pack.NewIdx(&pack.File{
		Context: ctx2,
		Type:    pack.FileTypeIdx,
		Path:    pk.GetDotIdx(),
	})
	if err != nil {
		return err
	}

	pack1path := pk.GetDotPack()

	list1, err := inst.loadIdx(idx1)
	if err != nil {
		return err
	}
	inst.sort(list1)
	size1 := len(list1)
	for i := 0; i < size1-10; i++ {
		if i < size1 {
			item1 := list1[i]
			item2 := list1[i+1]

			sum1 := item1.CRC32
			const f = "index:%v offset:%v size:%v oid:%v crc32:%v"
			crc := strconv.FormatUint(uint64(sum1), 16)
			oid := item1.OID.String()
			vlog.Info(fmt.Sprintf(f, item1.Index, item1.Offset, item1.Length, oid, crc))

			sum2, err := inst.computeItemDataCRC32(item1, item2, pack1path)
			if err != nil {
				return err
			}

			if sum1 != sum2 {
				return fmt.Errorf("bad crc32 sum, want:%v have:%v", sum1, sum2)
			}
		}
	}

	// list2, err := inst.scanPack(pack1)
	// if err != nil {
	// 	return err
	// }
	// size2 := len(list2)

	// for i := 0; i < size1 && i < size2; i++ {
	// 	item1 := list1[i]
	// 	item2 := list2[i]
	// 	vlog.Debug(item1, item2)
	// }

	return nil
}

func (inst *TestPackIdxCRC32) scanPack(p pack.Pack) ([]*git.PackedObjectHeaderEx, error) {
	return p.Scan()
}

func (inst *TestPackIdxCRC32) loadIdx(idx pack.Idx) ([]*git.PackIndexItem, error) {
	count := idx.Count()
	return idx.GetItems(0, int(count+10))
}

func (inst *TestPackIdxCRC32) sort(items []*git.PackIndexItem) {
	inst.items = items
	sort.Sort(inst)
}
func (inst *TestPackIdxCRC32) Len() int {
	return len(inst.items)
}
func (inst *TestPackIdxCRC32) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Offset < o2.Offset
}
func (inst *TestPackIdxCRC32) Swap(i1, i2 int) {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	inst.items[i1] = o2
	inst.items[i2] = o1
}

func (inst *TestPackIdxCRC32) computeItemDataCRC32(item1 *git.PackIndexItem, item2 *git.PackIndexItem, packPath afs.Path) (uint32, error) {

	pos1 := item1.Offset
	pos2 := item2.Offset

	r1, err := packPath.GetIO().OpenSeekerR(&afs.Options{Create: true, Mkdirs: true})
	if err != nil {
		return 0, err
	}
	defer func() {
		r1.Close()
	}()

	r1.Seek(pos1, io.SeekStart)
	r2 := io.LimitReader(r1, pos2-pos1)

	crcTable := crc32.MakeTable(crc32.IEEE)
	h := crc32.New(crcTable)
	_, err = io.Copy(h, r2)
	if err != nil {
		return 0, err
	}

	sum := h.Sum32()
	// crc := strconv.FormatUint(uint64(sum), 16)

	// oid := item1.OID
	const f = "compute_crc32_from_pack oid:%v crc32:%v"
	// vlog.Info(fmt.Sprintf(f, oid.String(), crc))
	// vlog.Info("")
	return sum, nil
}

func (inst *TestPackIdxCRC32) writeObjectHeadTo(dst io.Writer, head *git.PackedObjectHeaderEx) error {
	t := head.Type.ToObjectType().String()
	s := head.Size
	str := fmt.Sprintf("%v %v.", t, s)
	data := []byte(str)
	data[len(data)-1] = 0
	_, err := dst.Write(data)
	return err
}

func (inst *TestPackIdxCRC32) openOutFile(session store.Session, oid git.ObjectID, name string) (io.WriteCloser, error) {

	dir := session.GetRepository().Objects().Path()
	dir = dir.GetChild("pack/entity/" + oid.String())
	file := dir.GetChild(name)

	op := &afs.Options{Mkdirs: true, Create: true}
	return file.GetIO().OpenWriter(op)
}
