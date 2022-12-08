package unit

import (
	"testing"

	"github.com/bitwormhole/gitlib"
)

func TestGetLib(t *testing.T) {

	lib := gitlib.GetLib()

	dir := lib.FS().NewPath(t.TempDir())

	t.Log(dir.GetPath())
}
