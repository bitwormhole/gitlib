package git

// Commit ...
type Commit struct {
	Parents   []ObjectID
	Tree      ObjectID
	Author    *Operator
	Committer *Operator

	Ext map[string]*HyperHeader

	Content string
}
