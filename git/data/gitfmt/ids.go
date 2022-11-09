package gitfmt

import "github.com/bitwormhole/gitlib/git"

// CreateObjectID ...
func CreateObjectID(b []byte) (git.ObjectID, error) {
	return git.DefaultIdentityFactory().TryCreate(b)
}

// ParseObjectID ...
func ParseObjectID(text string) (git.ObjectID, error) {
	return git.DefaultIdentityFactory().TryParse(text)
}
