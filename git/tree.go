package git

// Tree ...
type Tree struct {
	ID    ObjectID // 此 tree 对象的 ID
	Items []*TreeItem
}

////////////////////////////////////////////////////////////////////////////////

// TreeItem ...
type TreeItem struct {
	Name string
	ID   ObjectID // 该条目的 ID
	Mode TreeItemMode
}

// IsFolder ...
func (inst *TreeItem) IsFolder() bool {
	return (inst.Mode == TreeItemModeFolder)
}

// IsFile ...
func (inst *TreeItem) IsFile() bool {
	return (inst.Mode == TreeItemModeFile) || (inst.Mode == TreeItemModeExe)
}

// Executable ...
func (inst *TreeItem) Executable() bool {
	return (inst.Mode == TreeItemModeExe)
}
