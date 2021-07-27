package ids

import (
	"strings"
	"testing"
)

func TestSha256convert(t *testing.T) {

	str0 := "77df263f49123356d28a4a8715d25bf5b980beeeb503cab46ea61ac9f3320eda"
	t.Log(str0)

	id0, err := ParseSha256id(str0)
	if err != nil {
		t.Error(err)
	}

	str1 := id0.ToObjectId().String()
	byt1 := id0.ToObjectId().Bytes()
	t.Log(str1)

	id1, err := CreateSha256id(byt1)
	if err != nil {
		t.Error(err)
	}

	str2 := id1.String()
	t.Log(str2)

	if strings.Compare(str0, str1) != 0 {
		t.Error("str0!=str1")
	}
	if strings.Compare(str0, str2) != 0 {
		t.Error("str0!=str2")
	}
}
