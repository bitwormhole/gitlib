package repositories

type WorktreeName string

// 工作树 (@用户端)
type WorktreeFrontend interface {

	// as a dir
	DirNode

	GetBackend() WorktreeBackend
}

// 工作树 (@仓库端)
type WorktreeBackend interface {

	// as a dir
	DirNode

	GetFrontend() WorktreeFrontend
}

// 工作树集合 (@仓库端)
type Worktrees interface {

	// as a dir
	DirNode

	GetBackend(name WorktreeName) WorktreeBackend

	GetFrontend(name WorktreeName) WorktreeFrontend

	ListNames() []WorktreeName
}
