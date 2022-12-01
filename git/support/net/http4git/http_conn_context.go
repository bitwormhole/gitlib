package http4git

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/git/network/pktline"
)

type connectionContext struct {
	rawURL  string // without 'info/refs'
	fullURL string // with 'info/refs'
}

func (inst *connectionContext) init(url string, useSSL bool) error {

	isHTTP := strings.HasPrefix(url, "http://")
	isHTTPS := strings.HasPrefix(url, "https://")

	if useSSL {
		if !isHTTPS {
			return fmt.Errorf("it's not a HTTPS URL [%v]", url)
		}
	} else {
		if !isHTTPS && !isHTTP {
			return fmt.Errorf("it's not a HTTP(or HTTPS) URL [%v]", url)
		}
	}

	inst.rawURL = url
	inst.fullURL = url + "/info/refs"
	return nil
}

func (inst *connectionContext) open(p *pktline.ConnParams) (pktline.Connection, error) {

	service := p.Service
	method := p.Method
	url := inst.fullURL + "?service=" + service

	conn := &httpConn{
		context: inst,
		url:     url,
		service: service,
		method:  method,
	}

	conn.init(method)
	return conn, nil
}
