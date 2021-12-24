// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	repository0x5aaf5a "github.com/bitwormhole/gitlib/git/repository"
	unit0xa250fd "github.com/bitwormhole/gitlib/src/test/golang/unit"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComUnitTest1 struct {
	instance *unit0xa250fd.UnitTest1
	 markup0x23084a.Component `initMethod:"Init"`
	RM repository0x5aaf5a.Manager `inject:"#git-repository-manager"`
	RepoPath string `inject:"${test.repo.path}"`
}

