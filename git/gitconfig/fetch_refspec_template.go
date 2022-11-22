package gitconfig

import (
	"strings"

	"github.com/bitwormhole/gitlib/git"
)

// FetchRefspecTemplate 表示一个 fetch 模板,
// like '+refs/heads/*:refs/remotes/origin/*',
// default key='remote.<name>.fetch'
type FetchRefspecTemplate string

func (t FetchRefspecTemplate) String() string {
	str := string(t)
	return strings.TrimSpace(str)
}

// GetLocalTemplate ...
func (t FetchRefspecTemplate) GetLocalTemplate() string {
	str := string(t)
	i1 := strings.Index(str, "+")
	i2 := strings.Index(str, ":")
	if 0 <= i1 && i1 < i2 {
		s2 := str[i1+1 : i2]
		return strings.TrimSpace(s2)
	}
	return ""
}

// GetRemoteTemplate ...
func (t FetchRefspecTemplate) GetRemoteTemplate() string {
	str := string(t)
	// i1 := strings.Index(str, "+")
	i2 := strings.Index(str, ":")
	if 0 < i2 {
		s2 := str[i2+1:]
		return strings.TrimSpace(s2)
	}
	return ""
}

// MakeLocalRefspecWithName ...
func (t FetchRefspecTemplate) MakeLocalRefspecWithName(name string) git.ReferenceName {
	i := strings.LastIndex(name, "/")
	if 0 <= i {
		name = name[i+1:]
	}
	if name == "" {
		name = "default"
	}
	temp := t.GetLocalTemplate()
	result := strings.ReplaceAll(temp, "*", name)
	return git.ReferenceName(result).Normalize()
}

// MakeRemoteRefspecWithName ...
func (t FetchRefspecTemplate) MakeRemoteRefspecWithName(name string) git.ReferenceName {
	i := strings.LastIndex(name, "/")
	if 0 <= i {
		name = name[i+1:]
	}
	if name == "" {
		name = "default"
	}
	temp := t.GetRemoteTemplate()
	result := strings.ReplaceAll(temp, "*", name)
	return git.ReferenceName(result).Normalize()
}
