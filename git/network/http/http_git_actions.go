package http

import (
	"github.com/bitwormhole/gitlib/git/network/clients"
)

type actions struct {
}

func (inst *actions) doGetGitUploadPackAdvertisement(c *clients.Context) error {
	loader := &gitUploadPackAdvertisementLoader{}
	return loader.load(c)
}

func (inst *actions) doPostGitUploadPack(c *clients.Context) error {
	return nil
}

func (inst *actions) doPostGitReceivePack(c *clients.Context) error {

	return nil
}
