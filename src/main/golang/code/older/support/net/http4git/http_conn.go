package http4git

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bitwormhole/gitlib/git/network/pktline"
)

type httpConn struct {
	context *connectionContext
	url     string
	service string
	method  string

	// io
	requestFIFO        ByteFIFO
	requestContentType string
	requestTx          pktline.WriterCloser
	responseRx         pktline.ReaderCloser

	// status
	started  bool
	starting bool
	stopped  bool
	stopping bool
}

func (inst *httpConn) _Impl() pktline.Connection {
	return inst
}

func (inst *httpConn) init(method string) {
	method = strings.TrimSpace(method)
	method = strings.ToUpper(method)
	if method == "" {
		method = http.MethodGet
	}
	inst.method = method
}

func (inst *httpConn) Close() error {
	inst.stopping = true
	return nil
}

func (inst *httpConn) GetGroup() pktline.ConnectionGroup {
	return nil
}

func (inst *httpConn) GetAttribute(name string) any {
	return nil
}

func (inst *httpConn) SetAttribute(name string, value any) {

}

func (inst *httpConn) GetParams() *pktline.ConnParams {
	return nil
}

func (inst *httpConn) GetService() string {
	return inst.service
}

// 创建新的附加连接
func (inst *httpConn) NewConnection(p *pktline.ConnParams) (pktline.Connection, error) {
	return inst.context.open(p)
}

func (inst *httpConn) OpenWriter(contentType string) (pktline.WriterCloser, error) {
	tx := inst.requestTx
	if tx != nil {
		return tx, nil
	}
	fifo, err := OpenByteFIFO()
	if err != nil {
		return nil, err
	}
	tx = pktline.NewWriterCloser(fifo, true)
	inst.requestFIFO = fifo
	inst.requestTx = tx
	inst.requestContentType = contentType
	return tx, nil
}

func (inst *httpConn) OpenReader() (pktline.ReaderCloser, string, error) {

	url := inst.url
	method := inst.method
	body1 := inst.requestFIFO
	client := http.DefaultClient

	req, err := http.NewRequest(method, url, body1)
	if err != nil {
		return nil, "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	code := resp.StatusCode
	body2 := resp.Body
	ctype := resp.Header.Get("Content-Type")

	defer func() {
		if body2 != nil {
			body2.Close()
		}
	}()

	if code != http.StatusOK {
		return nil, "", fmt.Errorf("HTTP %v", resp.Status)
	}

	rx := pktline.NewReaderCloser(body2, true)
	inst.responseRx = rx
	inst.started = true
	body2 = nil
	return rx, ctype, nil
}
