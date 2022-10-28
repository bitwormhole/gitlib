package dxo

// Tree ...
type Tree struct {
	Items []*TreeItem
}

// TreeItem ...
type TreeItem struct {
	Name string
	ID   int // todo ...
	Mode int
}
