package gitfmt

import (
	"testing"

	"github.com/bitwormhole/gitlib/git"
)

func TestHeadLS(t *testing.T) {

	h0 := &git.HEAD{
		Name: "refs/remotes/o1/abcdefg",
	}

	text, err := FormatHEAD(h0)
	if err != nil {
		t.Error(err)
		return
	}

	h1, err := ParseHEAD(text)
	if err != nil {
		t.Error(err)
		return
	}

	if h1.Name != h0.Name {
		t.Error("want:", h0.Name, " have:", h1.Name)
	}
}
