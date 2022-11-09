package gitfmt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/git"
)

// FormatTree ...
func FormatTree(tree *git.Tree) ([]byte, error) {
	tf := treeFormatter{}
	return tf.format(tree)
}

// ParseTree ...
func ParseTree(data []byte) (*git.Tree, error) {
	tp := treeReader{}
	tp.init(data)
	return tp.readAll()
}

////////////////////////////////////////////////////////////////////////////////

type treeFormatter struct{}

func (inst *treeFormatter) format(tree *git.Tree) ([]byte, error) {
	return nil, errors.New("no impl")
}

////////////////////////////////////////////////////////////////////////////////

type treeReader struct {
	data   []byte
	pos    int
	end    int
	idSize int // in bytes
}

func (inst *treeReader) init(data []byte) {
	if data == nil {
		data = []byte{}
	}
	inst.data = data
	inst.pos = 0
	inst.end = len(data)
}

func (inst *treeReader) readAll() (*git.Tree, error) {
	tree := &git.Tree{}
	for !inst.isEOF() {
		item := &git.TreeItem{}
		err := inst.readItem(item)
		if err != nil {
			return nil, err
		}
		tree.Items = append(tree.Items, item)
	}
	return tree, nil
}

func (inst *treeReader) isEOF() bool {
	return inst.end <= inst.pos
}

func (inst *treeReader) readItem(dst *git.TreeItem) error {
	err := inst.readItemHead(dst)
	if err != nil {
		return err
	}
	return inst.readItemID(dst)
}

func (inst *treeReader) readItemHead(dst *git.TreeItem) error {

	i0 := inst.pos // the begin of item
	i1 := i0       // the SPACE gap between mode & name
	i2 := i0       // end of head, '\0'
	end := inst.end
	data := inst.data

	for i := i1; i < end; i++ {
		b := data[i]
		if b == 0 {
			i2 = i
			break
		} else if b == ' ' && i1 == i0 {
			i1 = i
		}
	}

	if i0 < i1 && i1 < i2 && i2 < end {
		mode := data[i0:i1]
		name := data[i1+1 : i2]
		dst.Mode = strings.TrimSpace(string(mode))
		dst.Name = string(name)
	} else {
		return fmt.Errorf("bad tree item head")
	}

	inst.pos = i2 + 1
	inst.idSize = 20 // 20(bytes) = 160(bits)
	return nil
}

func (inst *treeReader) readItemID(dst *git.TreeItem) error {
	i1 := inst.pos
	i2 := inst.pos + inst.idSize
	if 0 < i1 && i1 < i2 && i2 <= inst.end {
		// continue
	} else {
		return fmt.Errorf("EOF with exception")
	}
	idData := inst.data[i1:i2]
	id, err := CreateObjectID(idData)
	if err != nil {
		return err
	}
	dst.ID = id
	inst.pos = i2
	return nil
}

////////////////////////////////////////////////////////////////////////////////
