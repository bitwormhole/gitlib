// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibdemo

import (
	demo0x52dcb1 "github.com/bitwormhole/gitlib/etc/demo"
	testcmds0x82bca1 "github.com/bitwormhole/gitlib/etc/demo/testcmds"
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComTestListObjectsInPack struct {
	instance *testcmds0x82bca1.TestListObjectsInPack
	 markup0x23084a.Component `class:"cli-handler-registry"`
	WD string `inject:"${test.repo.path}"`
	LA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComTestPackDeltaObjects struct {
	instance *testcmds0x82bca1.TestPackDeltaObjects
	 markup0x23084a.Component `class:"cli-handler-registry"`
	WD string `inject:"${test.repo.path}"`
	LA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComTestReadObjects struct {
	instance *testcmds0x82bca1.TestReadObjects
	 markup0x23084a.Component `class:"cli-handler-registry"`
	WD string `inject:"${test.repo.path}"`
	LA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComTestReadPackIdx struct {
	instance *testcmds0x82bca1.TestReadPackIdx
	 markup0x23084a.Component `class:"cli-handler-registry"`
	WD string `inject:"${test.repo.path}"`
	LA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComTestPoint struct {
	instance *demo0x52dcb1.TestPoint
	 markup0x23084a.Component `class:"life"`
	Context application0x67f6c5.Context `inject:"context"`
	Agent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
	CmdKey string `inject:"${test.gitlib.command}"`
	WD string `inject:"${test.repo.path}"`
}

