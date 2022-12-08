package testcmds

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// TestReadObjects ...
type TestReadObjects struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *TestReadObjects) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestReadObjects) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-read-objects",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestReadObjects) run(task *cli.Task) error {

	lib, err := inst.LA.GetLib()
	if err != nil {
		return err
	}

	wd := lib.FS().NewPath(inst.WD)
	repo, err := lib.RepositoryLoader().LoadWithPath(wd)
	if err != nil {
		return err
	}

	// session
	session, err := repo.OpenSession()
	if err != nil {
		return err
	}
	defer func() { session.Close() }()

	// head
	head, err := repo.HEAD().GetValue(session)
	if err != nil {
		return err
	}

	// ref
	ref := repo.Refs().GetRef(head)
	commitID, err := ref.GetValue(session)
	if err != nil {
		return err
	}

	const limit = 10
	return inst.scanCommits(commitID, session, limit)
}

func (inst *TestReadObjects) scanCommits(commitID git.ObjectID, session store.Session, limit int) error {
	id := commitID
	for i := 0; i < limit; i++ {
		if id == nil {
			break
		}
		commit, err := session.LoadCommit(id)
		if err != nil {
			return err
		}
		err = inst.scanCommit(commit, session)
		if err != nil {
			return err
		}
		id = inst.getParentCommit(commit)
	}
	return nil
}

func (inst *TestReadObjects) getParentCommit(commit *git.Commit) git.ObjectID {
	plist := commit.Parents
	for _, id := range plist {
		if id != nil {
			return id
		}
	}
	return nil
}

func (inst *TestReadObjects) scanCommit(commit *git.Commit, session store.Session) error {

	vlog.Warn("scan commit ", commit.ID.String())

	item := &git.TreeItem{
		ID:   commit.Tree,
		Name: "[ROOT]",
		Mode: git.TreeItemModeFolder,
	}

	err := inst.scanTree(item, session, 0)
	if err != nil {
		vlog.Error(err)
	}

	return nil
}

func (inst *TestReadObjects) scanTree(item *git.TreeItem, session store.Session, depth int) error {

	tab := inst.makeTabString(depth)
	treeid := item.ID
	vlog.Warn("scan tree ", treeid.String(), tab, item.Name)

	tree, err := session.LoadTree(treeid)
	if err != nil {
		return err
	}

	items := tree.Items
	for _, item := range items {
		if item.IsFolder() {
			err = inst.scanTree(item, session, depth+1)
			if err != nil {
				return err
			}
		} else if item.IsFile() {
			err = inst.scanBlob(item, session, depth+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (inst *TestReadObjects) scanBlob(item *git.TreeItem, session store.Session, depth int) error {

	oid := item.ID
	vlog.Debug("scan blob ", oid.String())

	info, in, err := session.ReadObject(item.ID)
	if err != nil {
		return err
	}
	defer func() {
		in.Close()
	}()

	t := info.Type
	l := info.Length
	h := fmt.Sprintf("%v %v.", t, l)
	head := []byte(h)
	head[len(head)-1] = 0

	hash := session.GetRepository().Digest().New()
	hash.Write(head)

	cb, err := io.Copy(hash, in)
	if err != nil {
		return err
	}
	if cb != l {
		return fmt.Errorf("bad length, want:%v have:%v", l, cb)
	}

	sum1 := oid.Bytes()
	sum2 := hash.Sum(nil)
	if !bytes.Equal(sum1, sum2) {
		return fmt.Errorf("bad sum, want:%v have:%v", sum1, sum2)
	}

	return nil
}

func (inst *TestReadObjects) makeTabString(depth int) string {
	builder := strings.Builder{}
	builder.WriteString(" ...")
	for i := 0; i < depth; i++ {
		builder.WriteString("....")
	}
	builder.WriteString(" ")
	return builder.String()
}
