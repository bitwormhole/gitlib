package repo

import (
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter/lang"
)

// DefaultRepositoryDriver 默认的仓库驱动
type DefaultRepositoryDriver struct {
	MyLocator repository.Locator
	MyFactory repository.Factory
}

func (inst *DefaultRepositoryDriver) _Impl() repository.Driver {
	return inst
}

func (inst *DefaultRepositoryDriver) Factory() repository.Factory {
	return inst.MyFactory
}

func (inst *DefaultRepositoryDriver) Locator() repository.Locator {
	return inst.MyLocator
}

func (inst *DefaultRepositoryDriver) Accept(uri lang.URI) bool {
	return inst.MyLocator.Accept(uri)
}
