package pack

import (
	"fmt"
	"io"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/starter/vlog"
)

// Pack 表示一个 pack-*.pack 文件
type Pack interface {
	Load() error
	Reload() error
	Check(flags CheckFlag) error
	GetPackID() git.PackID
	Scan() ([]*git.PackedObjectHeaderEx, error)

	OpenSimpleObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error)
	ReadSimpleObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error)

	// 如果 pool 参数为 nil, 则使用内部的 pool 提供数据来源
	OpenObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error)

	// 如果 pool 参数为 nil, 则使用内部的 pool 提供数据来源
	ReadObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error)

	// 把 PackIndexItem 转换为 PackedObjectHeaderEx
	IndexToHeader(item *git.PackIndexItem) *git.PackedObjectHeaderEx
}

// ComplexPack 表示一个 pack-*.pack 文件, 支持对象重建
type ComplexPack interface {
	Pack

	// 打开 in-pack 对象读者，如果是 delta 对象，执行重建
	OpenComplexObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error)

	// 读取 in-pack 对象头，如果是 delta 对象，执行重建
	ReadComplexObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error)

	// 生成 .idx 文件
	MakeIdx(idxFile afs.Path) error
}

////////////////////////////////////////////////////////////////////////////////

// NewPack ...
func NewPack(file *File) (Pack, error) {
	p := &EntityFile{}
	err := p.Init(file)
	if err != nil {
		return nil, err
	}
	err = p.Load()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// NewComplexPack ...
func NewComplexPack(file *File) (ComplexPack, error) {
	inner, err := NewPack(file)
	if err != nil {
		return nil, err
	}
	facade := &packExWrapper{
		inner: inner,
		file:  file,
	}
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

// EntityFile ...
type EntityFile struct {
	file *File
	pid  git.PackID

	signature [4]byte
	version   uint32
	count     uint32

	loaded  bool
	loading bool
}

func (inst *EntityFile) _Impl() Pack {
	return inst
}

// Init ...
func (inst *EntityFile) Init(file *File) error {
	inst.file = file
	return nil
}

func (inst *EntityFile) hasCheckFlag(a, b CheckFlag) bool {
	return (a & b) != 0
}

// Check ...
func (inst *EntityFile) Check(flags CheckFlag) error {

	err := inst.Load()
	if err != nil {
		return err
	}

	list := make([]func() error, 0)
	list = append(list, inst.Load)

	if inst.hasCheckFlag(flags, CheckHead) {
		list = append(list, inst.checkHead)
	}
	if inst.hasCheckFlag(flags, CheckSize) {
		list = append(list, inst.checkSize)
	}
	if inst.hasCheckFlag(flags, CheckSum) {
		list = append(list, inst.checkSum)
	}

	for _, step := range list {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

// func (inst *EntityFile) loadHead() error {
// 	return nil
// }

func (inst *EntityFile) checkSize() error {
	return nil
}

func (inst *EntityFile) checkSum() error {
	ch := packSumChecker{}
	return ch.check(inst.file)
}

func (inst *EntityFile) checkHead() error {

	signature := string(inst.signature[:])
	version := inst.version

	if signature != "PACK" {
		return fmt.Errorf("bad pack.signature, want:'PACK'")
	}

	if version != 2 && version != 3 {
		return fmt.Errorf("bad pack.version, want: 2 or 3")
	}

	return nil
}

// GetPackID ...
func (inst *EntityFile) GetPackID() git.PackID {
	return inst.pid
}

// Load ...
func (inst *EntityFile) Load() error {
	if inst.loaded {
		return nil
	}
	err := inst.Reload()
	if err != nil {
		return err
	}
	inst.loaded = true
	return nil
}

// Reload ...
func (inst *EntityFile) Reload() error {

	// state
	if inst.loading {
		return nil
	}
	inst.loading = true
	defer func() {
		inst.loading = false
	}()

	// open
	in, err := inst.file.OpenReader()
	if err != nil {
		return err
	}
	defer func() {
		in.Close()
	}()
	cr := commonReader{}

	// signature string
	signature, err := cr.read4Bytes(in)
	if err != nil {
		return err
	}

	// version   uint32
	version, err := cr.readUInt32(in)
	if err != nil {
		return err
	}

	// count     uint32
	count, err := cr.readUInt32(in)
	if err != nil {
		return err
	}

	// tail for id
	_, pid2, err := cr.readPackFileTail(inst.file)
	if err != nil {
		return err
	}

	// keep values
	inst.signature = signature
	inst.version = version
	inst.count = count
	inst.pid = pid2

	return inst.Check(CheckSize | CheckHead)
}

func (inst *EntityFile) openRawReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {

	if pool == nil {
		pool = inst.file.Context.Parent.Pool
	}

	if pool == nil {
		return nil, nil, fmt.Errorf("no reader pool")
	}

	ehr := entityHeaderReader{}
	file := inst.file.Path

	// open
	in, err := pool.OpenReader(file, nil)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if in != nil {
			in.Close()
		}
	}()

	// seek
	err = ehr.seek(item, in)
	if err != nil {
		return nil, nil, err
	}

	// read header
	hx, err := ehr.readHeader(item, in)
	if err != nil {
		return nil, nil, err
	}

	stream := in
	in = nil
	return hx, stream, nil
}

// IndexToHeader ...
func (inst *EntityFile) IndexToHeader(src *git.PackIndexItem) *git.PackedObjectHeaderEx {
	dst := &git.PackedObjectHeaderEx{}
	dst.OID = src.OID
	dst.PID = src.PID
	dst.Offset = src.Offset
	dst.Type = src.PackedType
	dst.Size = src.Length
	return dst
}

// OpenSimpleObjectReader ...
func (inst *EntityFile) OpenSimpleObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	hx, in1, err := inst.openRawReader(item, pool)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if in1 != nil {
			in1.Close()
		}
	}()
	// unzip
	compre := inst.file.Context.Parent.Compression
	in2, err := compre.NewReader(in1)
	if err != nil {
		return nil, nil, err
	}
	in3 := &entityReader{
		in1: in1,
		in2: in2,
	}
	in1 = nil
	return hx, in3, nil
}

// ReadSimpleObjectHeader ...
func (inst *EntityFile) ReadSimpleObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error) {
	hx, in, err := inst.openRawReader(item, pool)
	if err != nil {
		return nil, err
	}
	defer func() {
		in.Close()
	}()

	// err = inst.tryReadBodyPlainData(in, hx)
	// if err != nil {
	// 	return nil, err
	// }

	return hx, nil
}

// Scan ...
func (inst *EntityFile) Scan() ([]*git.PackedObjectHeaderEx, error) {
	ctx := inst.file.Context.Parent
	file := inst.file.Path
	reader := &packObjectOffsetScanner{context: ctx}
	return reader.scan(file)
}

// OpenObjectReader ...
func (inst *EntityFile) OpenObjectReader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, io.ReadCloser, error) {
	return inst.OpenSimpleObjectReader(item, pool)
}

// ReadObjectHeader ...
func (inst *EntityFile) ReadObjectHeader(item *git.PackedObjectHeaderEx, pool afs.ReaderPool) (*git.PackedObjectHeaderEx, error) {
	return inst.ReadSimpleObjectHeader(item, pool)
}

////////////////////////////////////////////////////////////////////////////////

type entityReader struct {
	in1    io.Closer
	in2    io.ReadCloser
	closed bool
}

func (inst *entityReader) Read(b []byte) (int, error) {
	in := inst.in2
	if in == nil || inst.closed {
		return 0, fmt.Errorf("this stream is closed")
	}
	return in.Read(b)
}

func (inst *entityReader) Close() error {
	clist := make([]io.Closer, 0)
	clist = append(clist, inst.in1)
	clist = append(clist, inst.in2)
	for _, c := range clist {
		if c == nil {
			continue
		}
		err := c.Close()
		if err != nil {
			vlog.Warn(err)
		}
	}
	inst.in1 = nil
	inst.in2 = nil
	inst.closed = true
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type entityHeaderReader struct{}

func (inst *entityHeaderReader) seek(item *git.PackedObjectHeaderEx, in io.ReadSeeker) error {
	pos1 := item.Offset
	pos2, err := in.Seek(pos1, io.SeekStart)
	if err != nil {
		return err
	}
	if pos1 != pos2 {
		return fmt.Errorf("cannot seek to position, want:%v have:%v", pos1, pos2)
	}
	return nil
}

func (inst *entityHeaderReader) getPosition(in io.ReadSeeker) int64 {
	n, err := in.Seek(0, io.SeekCurrent)
	if err != nil {
		n = 0
	}
	return n
}

func (inst *entityHeaderReader) readHeader(item *git.PackedObjectHeaderEx, in io.Reader) (*git.PackedObjectHeaderEx, error) {

	hx := &git.PackedObjectHeaderEx{}
	hx.OID = item.OID
	hx.PID = item.PID

	err := inst.readHeaderBase(hx, in)
	if err != nil {
		return nil, err
	}

	if hx.Type == git.PackedDeltaOFS {
		err = inst.readHeaderDeltaOFS(hx, in)
	} else if hx.Type == git.PackedDeltaRef {
		err = inst.readHeaderDeltaRef(hx, in)
	}

	if err != nil {
		return nil, err
	}

	hx.Offset = item.Offset
	return hx, nil
}

func (inst *entityHeaderReader) readHeaderBase(hx *git.PackedObjectHeaderEx, in io.Reader) error {
	buffer := entity7bitsBuffer{}
	err := buffer.load(in)
	if err != nil {
		return err
	}
	hx.Type = buffer.parseType()
	hx.Size = buffer.parseSize()
	return nil
}

func (inst *entityHeaderReader) readHeaderDeltaOFS(hx *git.PackedObjectHeaderEx, in io.Reader) error {
	buffer := entity7bitsBuffer{}
	err := buffer.load(in)
	if err != nil {
		return err
	}
	ofs, err := buffer.parseDeltaOffset()
	if err != nil {
		return err
	}
	hx.DeltaOffset = ofs
	return nil
}

func (inst *entityHeaderReader) readHeaderDeltaRef(hx *git.PackedObjectHeaderEx, in io.Reader) error {
	pid := hx.PID
	idsize := pid.Size()
	cb1 := idsize.SizeInBytes()
	buf := make([]byte, cb1)
	cb2, err := in.Read(buf)
	if err != nil {
		return err
	}
	if cb1 != cb2 {
		return fmt.Errorf("read id, with bad size")
	}
	xid := pid.GetFactory().Create(buf)
	hx.DeltaRef = xid.(git.ObjectID)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type entity7bitsBuffer struct {
	data   [16]byte
	length int
}

func (inst *entity7bitsBuffer) load(in io.Reader) error {
	buf := inst.data[:]
	size := len(buf)
	count := 0
	for i := 0; i < size; i++ {
		b := buf[i : i+1]
		cb, err := in.Read(b)
		if err != nil {
			return err
		} else if cb != 1 {
			return fmt.Errorf("bad head in stream")
		}
		count++
		c := uint8(b[0])
		if (c & 0x80) == 0 {
			inst.length = count
			return nil
		}
	}
	return fmt.Errorf("out of buffer")
}

func (inst *entity7bitsBuffer) parseType() git.PackedObjectType {
	b := inst.data[0]
	t := (b >> 4) & 0x07
	return git.PackedObjectType(t)
}

func (inst *entity7bitsBuffer) parseSize() int64 {
	value := int64(0)
	for i := inst.length - 1; i >= 0; i-- {
		b := inst.data[i]
		if i == 0 {
			value = (value << 4) | int64(b&0x0f)
		} else {
			value = (value << 7) | int64(b&0x7f)
		}
	}
	return value
}

func (inst *entity7bitsBuffer) parseDeltaOffset() (int64, error) {
	return inst.parseGitVLQ()
}

func (inst *entity7bitsBuffer) parseSimple7bitsInt() (int64, error) {
	const (
		mask7b = 0x7f
	)
	num := int64(0)
	for i := inst.length - 1; i >= 0; i-- {
		n := inst.data[i]
		num = (num << 7) | (int64(n) & mask7b)
	}
	return num, nil
}

func (inst *entity7bitsBuffer) parseGitVLQ() (int64, error) {
	// as Git-VLQ formatted int
	const (
		maskContinue = uint8(128) // 1000 000
		maskLength   = uint8(127) // 0111 1111
		lengthBits   = uint8(7)   // subsequent bytes has 7 bits to store the length
	)
	i := 0
	data := inst.data[:]
	var c byte
	c = data[i]
	i++
	var v = int64(c & maskLength)
	for c&maskContinue > 0 {
		v++
		if i < inst.length {
			c = data[i]
			i++
		} else {
			return 0, fmt.Errorf("buffer overflow")
		}
		v = (v << lengthBits) + int64(c&maskLength)
	}
	return v, nil
}

// ReadGitVLQ ... 读取 git 格式的变长整数
func ReadGitVLQ(in io.Reader) (int64, error) {
	b := entity7bitsBuffer{}
	err := b.load(in)
	if err != nil {
		return 0, err
	}
	return b.parseGitVLQ()
}

// ReadSimple7bitsInt ... 读取 git 格式的变长整数
func ReadSimple7bitsInt(in io.Reader) (int64, error) {
	b := entity7bitsBuffer{}
	err := b.load(in)
	if err != nil {
		return 0, err
	}
	return b.parseSimple7bitsInt()
}

////////////////////////////////////////////////////////////////////////////////
