// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibcfg

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	lib0x4595be "github.com/bitwormhole/gitlib/etc/lib"
	git0x229c8a "github.com/bitwormhole/gitlib/git"
	pktline0xd37953 "github.com/bitwormhole/gitlib/git/network/pktline"
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComConfigAlgorithms struct {
	instance *lib0x4595be.ConfigAlgorithms
	 markup0x23084a.Component `class:"git-context-configurer"`
	Algorithms []git0x229c8a.AlgorithmRegistry `inject:".git-algorithm-registry"`
}


type pComTheSHA1 struct {
	instance *lib0x4595be.TheSHA1
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComTheSHA256 struct {
	instance *lib0x4595be.TheSHA256
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComTheSHA512 struct {
	instance *lib0x4595be.TheSHA512
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComTheMD5 struct {
	instance *lib0x4595be.TheMD5
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComTheCRC32 struct {
	instance *lib0x4595be.TheCRC32
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComTheDeflate struct {
	instance *lib0x4595be.TheDeflate
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComThePlain struct {
	instance *lib0x4595be.ThePlain
	 markup0x23084a.Component `class:"git-algorithm-registry"`
}


type pComConfigCommands struct {
	instance *lib0x4595be.ConfigCommands
	 markup0x23084a.Component `class:"cli-handler-registry"`
}


type pComConfigConnectors struct {
	instance *lib0x4595be.ConfigConnectors
	 markup0x23084a.Component `class:"git-context-configurer"`
	Connectors []pktline0xd37953.ConnectorRegistry `inject:".pktline-connector-registry"`
}


type pComHTTPGitConnectorReg struct {
	instance *lib0x4595be.HTTPGitConnectorReg
	 markup0x23084a.Component `class:"pktline-connector-registry"`
}


type pComConfigContextBase struct {
	instance *lib0x4595be.ConfigContextBase
	 markup0x23084a.Component `class:"git-context-configurer"`
}


type pComConfigContextWithInstructions struct {
	instance *lib0x4595be.ConfigContextWithInstructions
	 markup0x23084a.Component `class:"git-context-configurer"`
	Instructions []store0x8467b3.ServiceRegistry `inject:".git-instruction-registry"`
}


type pComConfigCore struct {
	instance *lib0x4595be.ConfigCore
	 markup0x23084a.Component `class:"git-core-configurer"`
}


type pComConfigInstructions struct {
	instance *lib0x4595be.ConfigInstructions
	 markup0x23084a.Component `class:"git-instruction-registry"`
}


type pComGitlibAgent struct {
	instance *lib0x4595be.GitlibAgent
	 markup0x23084a.Component `id:"git-lib-agent" class:"life"`
	CLI cli0xf7c71e.CLI `inject:"#cli"`
	ContextConfigurers []store0x8467b3.ContextConfigurer `inject:".git-context-configurer"`
	CoreConfigurers []store0x8467b3.CoreConfigurer `inject:".git-core-configurer"`
}

