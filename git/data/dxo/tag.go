package dxo

// Tag ...
type Tag struct {
	Target    ObjectID
	Author    *User
	Committor *User
	Content   string
}
