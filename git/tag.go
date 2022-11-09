package git

// Tag ...
type Tag struct {
	Target    ObjectID
	Author    *Operator
	Committor *Operator
	Content   string
}
