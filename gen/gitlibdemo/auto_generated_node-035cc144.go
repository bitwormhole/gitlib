// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibdemo

import (
	demo0x52dcb1 "github.com/bitwormhole/gitlib/etc/demo"
	testcmds0x82bca1 "github.com/bitwormhole/gitlib/etc/demo/testcmds"
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComTestReadPackIdx struct {
	instance *testcmds0x82bca1.TestReadPackIdx
	 markup0x23084a.Component `class:"cli-handler-registry"`
	WD string `inject:"${test.repo.path}"`
	LA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComTestPoint struct {
	instance *demo0x52dcb1.TestPoint
	 markup0x23084a.Component `class:"life"`
	Agent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
	Command string `inject:"${test.gitlib.command}"`
	WD string `inject:"${test.repo.path}"`
}

