package sessions

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/starter-go/afs"
)

type tempFileFactory struct {
	t0    time.Time
	mu    sync.Mutex
	index int
}

func (inst *tempFileFactory) init() {
	inst.t0 = time.Now()
}

func (inst *tempFileFactory) newBuilder() *tempFileBuilder {

	inst.mu.Lock()
	defer func() {
		inst.mu.Unlock()
	}()

	inst.index++

	return &tempFileBuilder{
		prefix: "~",
		suffix: ".tmp",
		t0:     inst.t0,
		index0: inst.index,
	}
}

////////////////////////////////////////////////////////////////////////////////

type tempFileBuilder struct {
	dir    afs.Path
	prefix string
	suffix string
	nonce  string
	index0 int
	index1 int
	t0     time.Time
	t1     time.Time
}

func (inst *tempFileBuilder) Create() afs.Path {

	nonce := rand.Int63()
	inst.nonce = strconv.FormatInt(nonce, 10)
	inst.index1++
	inst.t1 = time.Now()

	dir := inst.dir
	if dir == nil {
		path := os.TempDir()
		dir = files.FS().NewPath(path)
	}

	parts := make([]string, 0)
	parts = append(parts, inst.formatTime(inst.t0))
	parts = append(parts, inst.formatTime(inst.t1))
	parts = append(parts, inst.formatInt(inst.index0))
	parts = append(parts, inst.formatInt(inst.index1))
	parts = append(parts, inst.nonce)

	builder := strings.Builder{}
	builder.WriteString(inst.prefix)
	for i, p := range parts {
		if i > 0 {
			builder.WriteRune('-')
		}
		builder.WriteString(p)
	}
	builder.WriteString(inst.suffix)
	return dir.GetChild(builder.String())
}

func (inst *tempFileBuilder) formatTime(t time.Time) string {
	n := t.Unix()
	return strconv.FormatInt(n, 10)
}

func (inst *tempFileBuilder) formatInt(n int) string {
	return strconv.Itoa(n)
}

////////////////////////////////////////////////////////////////////////////////

type tempBuffer struct {
	tmpFile   afs.Path
	flushSize int
	closed    bool
	buffer    bytes.Buffer
	out       io.WriteCloser
}

func (inst *tempBuffer) _Impl() store.TemporaryBuffer {
	return inst
}

func (inst *tempBuffer) GetTemporaryFile() afs.Path {
	return inst.tmpFile
}

func (inst *tempBuffer) GetFlushSize() int {
	return inst.flushSize
}

func (inst *tempBuffer) SetFlushSize(size int) {
	inst.flushSize = size
}

func (inst *tempBuffer) getOut() (io.WriteCloser, error) {
	o := inst.out
	if o != nil {
		return o, nil
	}
	if inst.closed {
		return nil, fmt.Errorf("stream is closed")
	}
	o, err := inst.tmpFile.GetIO().OpenWriter(&afs.Options{
		Mkdirs: true, Create: true,
	})
	if err != nil {
		return nil, err
	}
	inst.out = o
	return o, nil
}

func (inst *tempBuffer) doFlush() error {
	data := inst.buffer.Bytes()
	inst.buffer.Reset()
	o, err := inst.getOut()
	if err != nil {
		return err
	}
	_, err = o.Write(data)
	return err
}

func (inst *tempBuffer) tryFlush() error {
	size := inst.buffer.Len()
	if size < inst.flushSize {
		return nil
	}
	return inst.doFlush()
}

func (inst *tempBuffer) SaveToFile(dst afs.Path) error {

	err := inst.doFlush()
	if err != nil {
		return err
	}

	err = inst.doClose()
	if err != nil {
		return err
	}

	dir := dst.GetParent()
	if !dir.Exists() {
		dir.Mkdirs(&afs.Options{Create: true, Mkdirs: true})
	}

	src := inst.tmpFile
	return src.MoveTo(dst)
}

func (inst *tempBuffer) Write(b []byte) (int, error) {
	if inst.closed {
		return 0, fmt.Errorf("stream is closed")
	}
	n, err := inst.buffer.Write(b)
	if err != nil {
		return n, err
	}
	err = inst.tryFlush()
	return n, err
}

func (inst *tempBuffer) doClose() error {
	o := inst.out
	inst.out = nil
	inst.closed = true
	if o == nil {
		return nil
	}
	return o.Close()
}

func (inst *tempBuffer) Close() error {

	err := inst.doClose()
	if err != nil {
		return err
	}

	tmp := inst.tmpFile
	if tmp != nil {
		if tmp.IsFile() {
			tmp.Delete()
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
