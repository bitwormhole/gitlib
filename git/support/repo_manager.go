package support

import (
	"errors"

	"github.com/bitwormhole/gitlib/git/repository"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
)

// GitRepoManager 仓库管理器的默认实现
type GitRepoManager struct {
	markup.Component `id:"git-repository-manager"`

	Drivers []repository.Driver `inject:".git-repository-driver"`
}

func (inst *GitRepoManager) _Impl() repository.Manager {
	return inst
}

// Open 打开仓库
func (inst *GitRepoManager) Open(uri lang.URI) (repository.View, error) {
	all := inst.Drivers
	for _, driver := range all {
		if driver.Supports(uri) {
			return driver.Open(uri)
		}
	}
	return nil, errors.New("unsupported repository URI: " + uri.String())
}

// OpenByPath 通过路径打开仓库
func (inst *GitRepoManager) OpenByPath(path fs.Path) (repository.View, error) {
	if path == nil {
		return nil, errors.New("path==nil")
	}
	uri := path.URI()
	if uri == nil {
		return nil, errors.New("uri==nil")
	}
	return inst.Open(uri)
}
