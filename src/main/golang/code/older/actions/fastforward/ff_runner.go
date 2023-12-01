package fastforward

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

type fastForwardRunner struct {
	parent  *FastForward
	session store.Session
	ref     store.Ref
}

func (inst *fastForwardRunner) run() error {

	repo := inst.parent.theRepo
	session, err := repo.OpenSession()
	if err != nil {
		return err
	}
	defer func() {
		session.Close()
	}()
	inst.session = session

	name := inst.parent.theRef
	ref := repo.Refs().GetRef(name)
	inst.ref = ref

	return inst.apply()
}

func (inst *fastForwardRunner) apply() error {

	id1 := inst.parent.theOld
	id2 := inst.parent.theNew
	if id1 == nil || id2 == nil {
		return fmt.Errorf("param is nil")
	}

	zero := id1.GetFactory().Zero()
	if git.HashEqual(zero, id1) {
		id1 = nil
	}
	if git.HashEqual(zero, id2) {
		id2 = nil
	}

	err := inst.checkOldID(id1)
	if err != nil {
		return err
	}

	err = inst.checkSteps(id1, id2)
	if err != nil {
		return err
	}

	if id1 == nil && id2 == nil {
		return fmt.Errorf("bad request param")
	} else if id1 == nil {
		return inst.doCreate()
	} else if id2 == nil {
		return inst.doDelete()
	}
	return inst.doUpdate()
}

func (inst *fastForwardRunner) checkSteps(from, to git.ObjectID) error {

	return nil // todo ...
}

func (inst *fastForwardRunner) checkOldID(id1 git.ObjectID) error {

	session := inst.session
	ref := inst.ref
	isZero := (id1 == nil)

	if isZero {
		if ref.Exists() {
			return fmt.Errorf("the ref is exists")
		}
		return nil
	}

	id2, err := ref.GetValue(session)
	if err != nil {
		return err
	}

	eq := git.HashEqual(id1, id2)
	if !eq {
		s1 := id1.String()
		s2 := id2.String()
		return fmt.Errorf("bad old commit, want:%v have:%v", s1, s2)
	}

	return nil
}

func (inst *fastForwardRunner) doCreate() error {
	name := inst.ref.Name()
	value := inst.parent.theNew
	return inst.session.SaveRef(&git.Ref{
		Name: name,
		ID:   value,
	})
}

func (inst *fastForwardRunner) doUpdate() error {
	name := inst.ref.Name()
	value := inst.parent.theNew
	return inst.session.SaveRef(&git.Ref{
		Name: name,
		ID:   value,
	})
}

func (inst *fastForwardRunner) doDelete() error {
	file := inst.ref.Path()
	return file.Delete()
}
