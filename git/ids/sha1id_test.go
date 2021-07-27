package ids

import (
	"strings"
	"testing"
)

func TestSha1convert(t *testing.T) {

	str0 := "06b9d97e9e7617173d4873fb479f321429ea5282"
	t.Log(str0)

	id0, err := ParseSha1id(str0)
	if err != nil {
		t.Error(err)
	}

	str1 := id0.ToObjectId().String()
	byt1 := id0.ToObjectId().Bytes()
	t.Log(str1)

	id1, err := CreateSha1id(byt1)
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
