package git

// Tag ...
type Tag struct {
	HyperMessage

	Target    ObjectID
	Author    *Operator
	Committor *Operator
}
