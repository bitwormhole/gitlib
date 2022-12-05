package sessions

import (
	"fmt"
	"io"
	"strconv"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

type plainSparseObjectSaver struct {
	session store.Session
}

func (inst *plainSparseObjectSaver) readObjectMeta(so store.SparseObject) (*git.Object, error) {
	spo := inst.session.GetSparseObjects()
	obj, in, err := spo.ReadSparseObject(so)
	if err != nil {
		return nil, err
	}
	defer func() {
		if in != nil {
			in.Close()
		}
	}()
	return obj, nil
}

func (inst *plainSparseObjectSaver) Save(o *git.Object, data io.Reader) (*git.Object, error) {

	if o == nil || data == nil {
		return nil, fmt.Errorf("param is nil")
	}
	if o.Type == "" {
		return nil, fmt.Errorf("bad param:object.type")
	}
	if o.Length < 0 {
		return nil, fmt.Errorf("bad param:object.length")
	}

	// check old object
	repo := inst.session.GetRepository()
	wantID := o.ID
	if wantID != nil {
		old := repo.Objects().GetSparseObject(wantID)
		if old.Exists() {
			return inst.readObjectMeta(old)
		}
	}

	compress := repo.Compression()
	digest := repo.Digest().New()
	total := int64(0)

	// buffer
	buffer := make([]byte, 1024*4)
	temp := inst.session.NewTemporaryBuffer(nil)
	defer func() { temp.Close() }()
	out, err := compress.NewWriter(temp)
	if err != nil {
		return nil, err
	}
	defer func() { out.Close() }()

	// make head
	lenStr := strconv.FormatInt(o.Length, 10)
	headStr := o.Type.String() + string(' ') + lenStr + "."
	headBin := []byte(headStr)
	headBin[len(headBin)-1] = 0

	// write head
	out.Write(headBin)
	digest.Write(headBin)

	// for data
	for {
		n, err := data.Read(buffer)
		if n > 0 {
			data2 := buffer[0:n]
			digest.Write(data2)
			out.Write(data2)
			total += int64(n)
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	// check length
	if total != o.Length {
		return nil, fmt.Errorf("bad object length, want:%v, have:%v", o.Length, total)
	}

	// check id
	sum := digest.Sum([]byte{})
	haveID, err := git.CreateObjectID(sum[:])
	if err != nil {
		return nil, err
	}
	if wantID != nil {
		if !git.HashEqual(wantID, haveID) {
			const f = "bad object hash-id, want:%v, have:%v"
			return nil, fmt.Errorf(f, wantID.String(), haveID.String())
		}
	}

	// save
	obj := repo.Objects().GetSparseObject(haveID)
	err = temp.SaveToFile(obj.Path())
	if err != nil {
		return nil, err
	}

	return &git.Object{
		ID:     haveID,
		Type:   o.Type,
		Length: total,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////

type rawSparseObjectSaver struct {
	session store.Session
}

func (inst *rawSparseObjectSaver) Save(o *git.Object, data io.Reader) (*git.Object, error) {
	return nil, fmt.Errorf("no impl: rawSparseObjectSaver.save")
}

////////////////////////////////////////////////////////////////////////////////
