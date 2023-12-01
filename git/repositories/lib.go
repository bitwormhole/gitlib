package repositories

import (
	"context"

	"github.com/starter-go/afs"
	// "bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/network/pktline"
)

// Lib ...
type Lib interface {
	FS() afs.FS

	// GetCLI(required bool) cli.CLI

	Loader() Loader

	Finder() Finder

	Locator() Locator

	InstructionServiceManager() ServiceManager

	Connectors() pktline.ConnectorManager

	// 把这个 Lib 绑定到指定的 Context
	Bind(cc context.Context) context.Context
}

// LibAgent ... [inject:"#git-lib-agent"]
type LibAgent interface {
	GetLib() (Lib, error)
}

// GetLib 从给定的 Context 取与之绑定的 Lib 对象
func GetLib(cc context.Context) (Lib, error) {
	b, err := GetBinding(cc)
	if err != nil {
		return nil, err
	}
	return b.GetLib()
}
