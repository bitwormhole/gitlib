package testcmds

import (
	"fmt"
	"net/http"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/gitlib/git/network/servers"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
)

// TestServerAPI ...
type TestServerAPI struct {
	markup.Component `class:"cli-handler-registry"`

	WD string         `inject:"${test.repo.path}"`
	LA store.LibAgent `inject:"#git-lib-agent"`

	MainServer servers.MainServer `inject:"#git-main-server"`
}

func (inst *TestServerAPI) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *TestServerAPI) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "test-server-api",
		Handler: inst.run,
	}
	return []*cli.HandlerRegistration{hr}
}

func (inst *TestServerAPI) run(task *cli.Task) error {

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

	return inst.h(session)
}

func (inst *TestServerAPI) h(session store.Session) error {

	conn := &myTestServerConn{}
	repo := session.GetRepository()
	layout := repo.Layout()

	sc := &servers.Context{
		Protocol: "http",
		Method:   http.MethodGet,
		User:     "demo",
		Alias:    "demo",
		Service:  "git-upload-pack",

		Repository: repo,
		Session:    session,
		Layout:     layout,
		Connection: conn,
	}

	return inst.MainServer.Execute(sc)
}

////////////////////////////////////////////////////////////////////////////////

type myTestServerConn struct {
}

func (inst *myTestServerConn) _Impl() pktline.Connection {
	return inst
}

func (inst *myTestServerConn) GetGroup() pktline.ConnectionGroup {
	return nil
}

func (inst *myTestServerConn) GetAttribute(name string) any {
	return nil
}

func (inst *myTestServerConn) SetAttribute(name string, value any) {}

func (inst *myTestServerConn) GetParams() *pktline.ConnParams {
	return nil
}

func (inst *myTestServerConn) GetService() string {
	return ""
}

// @return (reader,contentType,error)
func (inst *myTestServerConn) OpenReader() (pktline.ReaderCloser, string, error) {
	return nil, "", fmt.Errorf("no impl")
}

func (inst *myTestServerConn) OpenWriter(contentType string) (pktline.WriterCloser, error) {
	return nil, fmt.Errorf("no impl")
}

// 创建新的附加连接
func (inst *myTestServerConn) NewConnection(p *pktline.ConnParams) (pktline.Connection, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *myTestServerConn) Close() error {
	return fmt.Errorf("no impl")
}
