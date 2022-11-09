package git

// Tree ...
type Tree struct {
	Items []*TreeItem
}

// TreeItem ...
type TreeItem struct {
	Name string
	ID   ObjectID // todo ...
	Mode string
}