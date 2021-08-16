package repo

import (
	"errors"

	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/lang"
)

// DefaultRepositoryManager 默认的仓库管理器
type DefaultRepositoryManager struct {
	Drivers []repository.Driver
}

func (inst *DefaultRepositoryManager) _Impl() repository.Manager {
	return inst
}

// FindDriver 根据URI查找驱动
func (inst *DefaultRepositoryManager) FindDriver(uri lang.URI) (repository.Driver, error) {
	list := inst.Drivers
	if list == nil {
		return nil, errors.New("drivers.list==nil")
	}
	for _, driver := range list {
		if driver == nil {
			continue
		}
		if driver.Accept(uri) {
			return driver, nil
		}
	}
	return nil, errors.New("no driver supports the uri=" + uri.String())
}

// GetAllDrivers 获取驱动列表
func (inst *DefaultRepositoryManager) GetAllDrivers() []repository.Driver {
	src := inst.Drivers
	if src == nil {
		empty := make([]repository.Driver, 0)
		return empty
	}
	dst := make([]repository.Driver, 0, len(src))
	copy(dst, src)
	return dst
}

// Open 打开仓库
func (inst *DefaultRepositoryManager) Open(uri lang.URI) (repository.Viewport, error) {

	driver, err := inst.FindDriver(uri)
	if err != nil {
		return nil, err
	}

	location, err := driver.Locator().Locate(uri)
	if err != nil {
		return nil, err
	}

	vpt, err := driver.Factory().Open(location)
	if err != nil {
		return nil, err
	}

	return vpt, nil
}
