package git

import (
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// Ref 是表示一个 .git/refs/[type]/[name] 的实体
type Ref struct {
	ID   ObjectID      // the Primary-Key
	Name ReferenceName // 不常用
}

////////////////////////////////////////////////////////////////////////////////

// ReferenceName is the name for .git/refs/*
type ReferenceName string

func (v ReferenceName) String() string {
	return string(v)
}

// NormalizeWithType ...
func (v ReferenceName) NormalizeWithType(aType string) ReferenceName {

	aType = strings.TrimSpace(aType)
	if aType == "" {
		aType = "heads"
	}

	// parse
	str := v.String()
	parts := strings.Split(str, "/")
	buffer := make([]string, 0)
	lastPart := ""

	for i := len(parts) - 1; i >= 0; i-- {
		p := parts[i]
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		buffer = append(buffer, p)
		lastPart = p
	}

	if lastPart == "" {
		buffer = append(buffer, "unnamed")
	}

	// append
	if lastPart != "refs" {
		if lastPart != aType {
			buffer = append(buffer, aType)
		}
		buffer = append(buffer, "refs")
	}

	// rebuild
	builder := strings.Builder{}
	sep := ""
	for i := len(buffer) - 1; i >= 0; i-- {
		p := buffer[i]
		builder.WriteString(sep)
		builder.WriteString(p)
		sep = "/"
	}
	str = builder.String()
	return ReferenceName(str)
}

// Normalize ...
func (v ReferenceName) Normalize() ReferenceName {
	return v.NormalizeWithType("heads")
}

// IsWildcard ...
func (v ReferenceName) IsWildcard() bool {
	str := string(v)
	return strings.ContainsRune(str, '*')
}

////////////////////////////////////////////////////////////////////////////////
