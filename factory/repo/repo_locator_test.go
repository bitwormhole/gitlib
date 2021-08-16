package repo

import (
	"os"
	"strings"
	"testing"

	srctest "github.com/bitwormhole/gitlib/src/test"
	"github.com/bitwormhole/starter/io/fs"
)

func TestLocator(t *testing.T) {

	res := srctest.ExportResources()
	all := res.All()
	for _, x := range all {
		t.Log(x.Name)
	}

	const keyBegin = "test."
	const keyEnd = ".repo"
	plist := make([]string, 0)
	env := os.Environ()
	for _, kv := range env {
		i := strings.Index(kv, "=")
		if i < 0 {
			continue
		}
		key := strings.TrimSpace(kv[0:i])
		val := strings.TrimSpace(kv[i+1:])
		if strings.HasPrefix(key, keyBegin) && strings.HasSuffix(key, keyEnd) {
			plist = append(plist, val)
		}
	}

	for _, p := range plist {
		path := fs.Default().GetPath(p)
		locator := (&DefaultRepositoryLocator{})._Impl()
		location, err := locator.Locate(path)
		if err == nil {
			t.Log(location.CoreDirectory.Path())
		} else {
			t.Error(err)
		}
	}

}
