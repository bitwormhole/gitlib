package http

import (
	"fmt"
	"io"
	"net/http"

	"bitwormhole.com/starter/vlog"
	"github.com/bitwormhole/gitlib/git/network/clients"
	"github.com/bitwormhole/gitlib/git/network/pktline"
)

type gitUploadPackAdvertisementLoader struct {
}

func (inst *gitUploadPackAdvertisementLoader) load(c *clients.Context) error {

	conn, err := c.Connection.NewConnection(&pktline.ConnParams{
		Method:  http.MethodGet,
		Service: pktline.ServiceGitUploadPack,
	})
	if err != nil {
		return err
	}

	r, t, err := conn.OpenReader()
	if err != nil {
		return err
	}
	defer func() {
		r.Close()
	}()

	if t != pktline.TypeGitUploadPackAdvertisement {
		return fmt.Errorf("unsupported content type:%v", t)
	}

	for {
		p, err := r.Read()
		if p != nil {
			err = inst.handlePacket(c, p)
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}

	return nil
}

func (inst *gitUploadPackAdvertisementLoader) handlePacket(c *clients.Context, p *pktline.Packet) error {

	h := p.Head
	b := p.Body
	bodySize := 0
	if b != nil {
		bodySize = len(b)
	}

	if h == "" && bodySize == 0 {
		h = "'0000'"
	}

	vlog.Info(h)
	return nil
}
