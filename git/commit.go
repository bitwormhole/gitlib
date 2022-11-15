package git

// Commit ...
type Commit struct {
	HyperMessage

	Parents   []ObjectID
	Tree      ObjectID
	Author    *Operator
	Committer *Operator
}
