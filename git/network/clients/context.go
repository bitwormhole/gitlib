package clients

import (
	"context"
	"net/http"

	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// Context ...
type Context struct {
	Lib        repositories.Lib
	Repository repositories.Repository
	Session    repositories.Session
	Path       afs.Path
	Connection pktline.Connection

	// Actions Actions

	// collections
	Branches map[string]*gitconfig.Branch // all
	Remotes  map[string]*gitconfig.Remote // all
	RnB      []*gitconfig.RemoteAndBranch
	Intents  []*Intent

	// current
	Intent        Intent
	RawRemote     gitconfig.Remote
	RawBranch     gitconfig.Branch
	FetchTemplate gitconfig.FetchRefspecTemplate
	RemoteName    string
}

// Init ...
func (inst *Context) Init(ctx context.Context, wd afs.Path) error {

	lib, err := repositories.GetLib(ctx)
	if err != nil {
		return err
	}

	repo, err := inst.loadRepository(wd, lib)
	if err != nil {
		return err
	}

	inst.Lib = lib
	inst.Repository = repo
	inst.Path = wd
	return nil
}

func (inst *Context) loadRepository(wd afs.Path, lib repositories.Lib) (repositories.Repository, error) {

	if wd == nil {
		return nil, nil
	}

	layout, err := lib.Locator().Locate(wd)
	if err != nil {
		return nil, err
	}

	return lib.Loader().Load(layout)
}

// OpenSession ...
func (inst *Context) OpenSession() (repositories.Session, error) {
	repo := inst.Repository
	session, err := repo.OpenSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// OpenConnection ...
func (inst *Context) OpenConnection(params *pktline.ConnParams) (pktline.Connection, error) {

	conn := inst.Connection
	intent := &inst.Intent
	if conn != nil {
		return conn, nil
	}

	if params == nil {
		params = &pktline.ConnParams{}
	}
	if params.Service == "" {
		params.Service = pktline.ServiceGitUploadPack
	}
	if params.URL == "" {
		params.URL = intent.URL
	}
	if params.Method == "" {
		params.Method = http.MethodGet
	}

	lib := inst.Lib
	conn, err := lib.Connectors().Connect(params)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Clone ...
func (inst *Context) Clone() *Context {
	src := inst
	dst := &Context{}
	*dst = *src
	return dst
}
