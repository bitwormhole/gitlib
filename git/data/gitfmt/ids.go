package gitfmt

import "github.com/bitwormhole/gitlib/git"

// CreateObjectID ...
func CreateObjectID(b []byte) (git.ObjectID, error) {
	return git.CreateObjectID(b)
}

// ParseObjectID ...
func ParseObjectID(text string) (git.ObjectID, error) {
	return git.ParseObjectID(text)
}
