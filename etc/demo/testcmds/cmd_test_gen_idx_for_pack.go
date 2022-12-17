package testcmds

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/objects/pack"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// TestGenIdxForPack ...
type TestGenIdxForPack struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *TestGenIdxForPack) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestGenIdxForPack) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-gen-idx-for-pack",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestGenIdxForPack) run(task *cli.Task) error {

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

	// dao := session.GetPacks()
	plist := repo.Objects().ListPacks()
	for _, pid := range plist {
		err = inst.genIdx(session, pid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *TestGenIdxForPack) genIdx(session store.Session, pid git.PackID) error {

	repo := session.GetRepository()
	pack1 := repo.Objects().GetPack(pid)
	path := pack1.GetDotPack()

	d := repo.Digest()
	c := repo.Compression()
	p := session.GetReaderPool()
	pack2, err := pack.NewPack(&pack.File{
		Digest:      d,
		Compression: c,
		Pool:        p,
		Path:        path,
		Type:        pack.FileTypePack,
	})
	if err != nil {
		return err
	}

	_, err = pack2.Scan()
	if err != nil {
		return err
	}

	return nil
}
