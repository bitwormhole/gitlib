package gitfmt

import (
	"testing"

	"github.com/bitwormhole/gitlib/git"
)

func TestRefLS(t *testing.T) {

	ids := git.DefaultIdentityFactory()
	//               0123456789abcdef0123456789abcdef, 128-bits
	id := ids.Parse("0123456789abcdef0123456789abcdef")

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

	if !git.HashEqual(r1.ID, r0.ID) {
		t.Error("want:", r0.ID, " have:", r1.ID)
	}
}
