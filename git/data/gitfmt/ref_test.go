package gitfmt

import (
	"testing"

	"github.com/bitwormhole/gitlib/git"
)

func TestRefLS(t *testing.T) {

	// ids := git.DefaultIdentityFactory()
	//               0123456789abcdef0123456789abcdef, 128-bits
	id, err := git.ParseObjectID("0123456789abcdef0123456789abcdef")
	if err != nil {
		t.Error(err)
		return
	}

	r0 := &git.Ref{ID: id}
	text, err := FormatRef(r0)
	if err != nil {
		t.Error(err)
		return
	}

	r1, err := ParseRef(text)
	if err != nil {
		t.Error(err)
		return
	}

	if !r1.ID.Equals(r0.ID) {
		t.Error("want:", r0.ID, " have:", r1.ID)
	}

	return
}
