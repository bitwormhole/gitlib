package repository

// Factory 是仓库的工厂
type Factory interface {

	// Open 打开位于 location 的仓库
	Open(location *Location) (Viewport, error)
}
