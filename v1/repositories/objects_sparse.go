package repositories

// 表示一个稀疏对象
type SparseObject interface {
	FileNode

	GetID() ObjectID

	GetObjects() Objects
}
