package git

import (
	"strings"
	"testing"
)

func TestRefname(t *testing.T) {

	src := map[string]string{
		"tags:abcd-1":                    "refs/tags/abcd-1",
		"remotes:abc/def-2":              "refs/remotes/abc/def-2",
		"heads:main-3":                   "refs/heads/main-3",
		"heads:heads/abc-4":              "refs/heads/abc-4",
		"remotes:remotes/abc/def-5":      "refs/remotes/abc/def-5",
		"heads:refs/heads/abc-6":         "refs/heads/abc-6",
		"remotes:refs/remotes/abc/def-7": "refs/remotes/abc/def-7",
		":heads/xyz-8":                   "refs/heads/xyz-8",
		":heads/xyz/def-9":               "refs/heads/xyz/def-9",
		":refs/heads/xyz-10":             "refs/heads/xyz-10",
		":refs/remotes/xyz/def-11":       "refs/remotes/xyz/def-11",
		":":                              "refs/heads/unnamed",
	}

	for k, want := range src {
		kk := strings.Split(k, ":")
		aType := kk[0]
		from := kk[1]
		name1 := ReferenceName(from)
		name2 := name1.NormalizeWithType(aType)
		t.Log("want_type:", aType, " name1:", name1, " >>>> name2:", name2)
		if want != name2.String() {
			t.Error("want:", want, " have:", name2)
		}
	}

	t.Log("done")
}
