package testcmds

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// TestReadPackIdx ...
type TestReadPackIdx struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *TestReadPackIdx) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestReadPackIdx) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-read-pack-idx",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestReadPackIdx) run(task *cli.Task) error {

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

	dao := session.GetPacks()
	plist := repo.Objects().ListPacks()
	for _, pid := range plist {
		err := dao.CheckPack(pid, pack.CheckAll)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *TestReadPackIdx) h(*git.PackIndex) {
}
