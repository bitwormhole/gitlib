package algorithms

import (
	"testing"

	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/gitlib/git/data/dxo"
)

func TestPathMapDefault(t *testing.T) {

	path1 := files.FS().NewPath(t.TempDir())
	id := dxo.ObjectID("c6db8558854411a3631aa9a394f44b934de1c3c3")

	pm0 := &PathPatternMapping{}
	pm := pm0.WithPattern("xx/xxx/xxxx/xx")

	path2 := pm.Map(path1, id)
	t.Log("path1 = ", path1.String())
	t.Log("path2 = ", path2.String())
}
