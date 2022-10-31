package store

import (
	"context"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/instructions"
)

// Lib ...
type Lib interface {
	FS() afs.FS

	GetCLI(required bool) cli.CLI

	RepositoryLoader() RepositoryLoader

	RepositoryFinder() RepositoryFinder

	RepositoryLocator() RepositoryLocator

	InstructionServiceManager() instructions.ServiceManager

	// 把这个 Lib 绑定到指定的 Context
	Bind(cc context.Context) context.Context
}

// GetLib 从给定的 Context 取与之绑定的 Lib 对象
func GetLib(cc context.Context) (Lib, error) {
	b, err := GetBinding(cc)
	if err != nil {
		return nil, err
	}
	return b.GetLib()
}
