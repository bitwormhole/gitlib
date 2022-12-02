package unit

import (
	"bytes"
	"crypto/rand"
	"testing"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

func TestRepoObjects(t *testing.T) {

	unit := initUnit(t)
	repoName := "foo"

	wd := unit.tmp.GetChild("a/b/c")
	wd.Mkdirs(nil)
	wd2 := wd.GetChild(repoName)

	// init repo
	client := unit.lib.GetCLI(true).GetClient()
	err := client.Run(&cli.Task{
		Context: unit.context,
		Command: "git init " + repoName,
		WD:      wd.GetPath(),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	////////////////////////////////
	// open repo
	repo, err := unit.lib.RepositoryLoader().LoadWithPath(wd2)
	if err != nil {
		t.Fatal(err)
		return
	}

	session, err := repo.OpenSession()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer func() { session.Close() }()

	// write obj
	err = doTestWriteSparseObject(session)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("OK")
}

func doTestWriteSparseObject(session store.Session) error {

	data := make([]byte, 1024)
	rand.Read(data)
	buffer := &bytes.Buffer{}
	buffer.Write(data)
	size := len(data)

	obj := &git.Object{
		Type:   git.ObjectTypeBLOB,
		Length: int64(size),
	}
	obj, err := session.GetSparseObjects().WriteSparseObject(obj, buffer)
	if err != nil {
		return err
	}
	return nil
}
