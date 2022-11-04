package lib

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/gitlib/git/commands"
	"github.com/bitwormhole/starter/markup"
)

// ConfigCommands ...
type ConfigCommands struct {
	markup.Component `class:"cli-handler-registry"`
}

func (inst *ConfigCommands) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *ConfigCommands) GetHandlers() []*cli.HandlerRegistration {

	list := []*cli.HandlerRegistration{}
	list = inst.fill(list, &commands.GitCmd{})

	list = inst.fill(list, &commands.GitAdd{})
	list = inst.fill(list, &commands.GitCommit{})
	list = inst.fill(list, &commands.GitInit{})
	list = inst.fill(list, &commands.GitPush{})
	list = inst.fill(list, &commands.GitStatus{})

	return list
}

func (inst *ConfigCommands) fill(dst []*cli.HandlerRegistration, src cli.HandlerRegistry) []*cli.HandlerRegistration {
	some := src.GetHandlers()
	dst = append(dst, some...)
	return dst
}
