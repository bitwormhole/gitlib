package http

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/git/network/clients"
)

// GitClient ...
type GitClient struct {
	// markup.Component `class:"git-client-registry"`
}

func (inst *GitClient) _Impl() (clients.ClientRegistry, clients.Client) {
	return inst, inst
}

// GetClientRegistration ...
func (inst *GitClient) GetClientRegistration() *clients.ClientRegistration {
	return &clients.ClientRegistration{
		Client: inst,
	}
}

// Accept ...
func (inst *GitClient) Accept(c *clients.Context) bool {
	i := &c.Intent
	url := i.URL
	b1 := strings.HasPrefix(url, "http://")
	b2 := strings.HasPrefix(url, "https://")
	return b1 || b2
}

// Execute ...
func (inst *GitClient) Execute(c *clients.Context) error {
	i := &c.Intent
	action := i.Action
	switch action {
	case clients.ActionFetch:
		return inst.doFetch(c)
	case clients.ActionPush:
		return inst.doPush(c)
	default:
		break
	}
	return fmt.Errorf("unsupported action:%v", action.String())
}

func (inst *GitClient) doPush(c *clients.Context) error {
	actions := &actions{}
	err := actions.doGetGitUploadPackAdvertisement(c)
	if err != nil {
		return err
	}

	err = actions.doPostGitReceivePack(c)
	if err != nil {
		return err
	}

	return nil
}

func (inst *GitClient) doFetch(c *clients.Context) error {
	actions := &actions{}
	err := actions.doGetGitUploadPackAdvertisement(c)
	if err != nil {
		return err
	}

	err = actions.doPostGitUploadPack(c)
	if err != nil {
		return err
	}

	return nil
}
