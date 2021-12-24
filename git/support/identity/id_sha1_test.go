package identity

import (
	"strings"
	"testing"
)

func TestSha1convert(t *testing.T) {

	str0 := "06b9d97e9e7617173d4873fb479f321429ea5282"
	t.Log(str0)

	factory := GetSha1IDFactory()

	id0, err := factory.Parse(str0)
	if err != nil {
		t.Error(err)
	}

	str1 := id0.String()
	byt1 := id0.Bytes()
	t.Log(str1)

	id1, err := factory.Create(byt1)
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
