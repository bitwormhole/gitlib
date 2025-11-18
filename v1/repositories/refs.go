package repositories

// 表示一个引用名称, 例如: 'refs/heads/main'
type Refname string

type Refs interface {

	// as a dir
	DirNode
}

type Ref interface {

	// as a dir
	FileNode
}
