package testcmds

import (
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

	return nil
}
