package gitfmt

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gitlib/git"
)

// FormatRef ...
func FormatRef(ref *git.Ref) (string, error) {
	if ref == nil {
		return "", errors.New("the param 'ref' is nil")
	}
	str := ref.ID.String()
	str = strings.TrimSpace(str) + "\n"
	size := ref.ID.Size()
	if size < 100 {
		return "", errors.New("bad object-id: " + str)
	}
	return str, nil
}

// ParseRef ...
func ParseRef(text string) (*git.Ref, error) {
	oid, err := ParseObjectID(text)
	if err != nil {
		return nil, err
	}
	ref := &git.Ref{
		ID: oid,
	}
	return ref, nil
}
