package git

import "github.com/bitwormhole/gitlib/git/files"

// UserContext 表示用户上下文
type UserContext interface {
	GetUserConfig() UserConfig
}

// SystemContext 表示系统上下文
type SystemContext interface {
	GetSystemConfig() SystemConfig
}

// Context 表示git(仓库)上下文
type Context interface {
	UserContext() UserContext
	SystemConfig() SystemConfig
	RepositoryLocator() RepositoryLocator
	RepositoryFinder() RepositoryFinder
	RepositoryFactory() RepositoryFactory
	FS() files.FS
}
