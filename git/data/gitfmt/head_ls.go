package gitfmt

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gitlib/git"
)

// FormatHEAD ...
func FormatHEAD(h *git.HEAD) (string, error) {
	str := h.Name.Normalize().String()
	return "ref: " + str, nil
}

// ParseHEAD ...
func ParseHEAD(text string) (*git.HEAD, error) {
	parts := strings.Split(text, ":")
	if len(parts) != 2 {
		return nil, errors.New("bad HEAD value: " + text)
	}
	p1 := strings.TrimSpace(parts[0])
	p2 := strings.TrimSpace(parts[1])
	if p1 != "ref" || !strings.HasPrefix(p2, "refs/") {
		return nil, errors.New("bad HEAD value: " + text)
	}
	name := git.ReferenceName(p2)
	h := &git.HEAD{}
	h.Name = name.Normalize()
	return h, nil
}
