package git

// Tree ...
type Tree struct {
	ID    ObjectID // 此 tree 对象的 ID
	Items []*TreeItem
}

// TreeItem ...
type TreeItem struct {
	Name string
	ID   ObjectID // 该条目的 ID
	Mode string
}
