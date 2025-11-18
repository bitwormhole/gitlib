package repositories

type Repository interface {

	// as a dir
	DirNode

	// 取仓库的布局
	GetLayout() Layout

	// 取对象文件夹
	GetObjects() Objects

	// 取仓库的工作树集合
	GetWorktrees() Worktrees

	// 取仓库的主工作树, 如果是裸仓库,返回 nil
	GetWorktree() WorktreeFrontend
}
