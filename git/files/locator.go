package files

import "github.com/bitwormhole/starter/io/fs"

// RepositoryLocator 仓库定位器
type RepositoryLocator interface {
	Locate(path fs.Path) (*RepositoryLocation, error)
}
