package dxo

// Commit ...
type Commit struct {
	Parents   []ObjectID
	Author    *User
	Committor *User
	Content   string
}
