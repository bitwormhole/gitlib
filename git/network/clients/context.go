package clients

import (
	"context"
	"net/http"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/store"
)

////////////////////////////////////////////////////////////////////////////////

// Local ...
type Local struct {
	Lib        store.Lib
	Repository store.Repository
	Session    store.Session
	Path       afs.Path
}

// Remote ...
type Remote struct {
	Raw gitconfig.Remote

	RemoteName    string
	Connection    pktline.Connection
	URL           string
	FetchTemplate gitconfig.FetchRefspecTemplate
}

// Branch ...
type Branch struct {
	Raw gitconfig.Branch

	Merge     git.ReferenceName // like 'refs/heads/main'
	FetchRef  git.ReferenceName // like 'refs/remotes/origin/main'
	RemoteRef git.ReferenceName // like 'refs/heads/main'
}

////////////////////////////////////////////////////////////////////////////////

// Context ...
type Context struct {
	Local  Local
	Remote Remote
	Branch Branch
}

// Init ...
func (inst *Context) Init(ctx context.Context, wd afs.Path) error {

	lib, err := store.GetLib(ctx)
	if err != nil {
		return err
	}

	repo, err := inst.loadRepository(wd, lib)
	if err != nil {
		return err
	}

	local := &inst.Local
	local.Lib = lib
	local.Repository = repo
	local.Path = wd
	return nil
}

func (inst *Context) loadRepository(wd afs.Path, lib store.Lib) (store.Repository, error) {

	if wd == nil {
		return nil, nil
	}

	layout, err := lib.RepositoryLocator().Locate(wd)
	if err != nil {
		return nil, err
	}

	return lib.RepositoryLoader().Load(layout)
}

// OpenSession ...
func (inst *Context) OpenSession() (store.Session, error) {
	repo := inst.Local.Repository
	session, err := repo.OpenSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// OpenConnection ...
func (inst *Context) OpenConnection(service string) (pktline.Connection, error) {

	remote := &inst.Remote
	conn := remote.Connection
	if conn != nil {
		return conn, nil
	}

	params := &pktline.ConnParams{
		Service: service,
		URL:     remote.URL,
		Method:  http.MethodGet,
	}

	lib := inst.Local.Lib
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
