package servers

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
)

// MainServerImpl 实现 MainServer 这个接口
type MainServerImpl struct {
	markup.Component `id:"git-main-server"`

	ServerRegistryList []ServerRegistry `inject:".git-server-registry"`

	registrations []*ServerRegistration
}

func (inst *MainServerImpl) _Impl() MainServer {
	return inst
}

// Accept ...
func (inst *MainServerImpl) Accept(c *Context) bool {
	return true
}

// Execute ...
func (inst *MainServerImpl) Execute(c *Context) error {
	all := inst.getAll()
	for _, r2 := range all {
		ser := r2.Server
		if ser.Accept(c) {
			return ser.Execute(c)
		}
	}
	str := inst.stringify(c)
	return fmt.Errorf("no handler support the servers.Context: %v", str)
}

func (inst *MainServerImpl) getAll() []*ServerRegistration {
	dst := inst.registrations
	if dst != nil {
		return dst
	}
	src := inst.ServerRegistryList
	for _, r1 := range src {
		if r1 == nil {
			continue
		}
		r2 := r1.GetServerRegistration()
		if r2 == nil {
			continue
		}
		ser := r2.Server
		if ser == nil {
			continue
		}
		dst = append(dst, r2)
	}
	inst.registrations = dst
	return dst
}

func (inst *MainServerImpl) stringify(c *Context) string {
	b := strings.Builder{}
	b.WriteString(c.Protocol)
	b.WriteString("://")
	b.WriteString(c.User)
	b.WriteString("@0.0.0.0/")
	b.WriteString(c.Alias)
	b.WriteString("?service=")
	b.WriteString(c.Service)
	// b.WriteString(c.Method)
	return b.String()
}
