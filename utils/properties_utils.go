package utils

import (
	"sort"
	"strings"

	"github.com/starter-go/application/properties"
)

// PropertiesTools 提供一组操作 collection.Properties 的快捷工具
type PropertiesTools struct {
}

// ListNames 根据提供的前缀和后缀，找出相应的属性名称
// prefix like 'xxxxxxx.'
// suffix like '.xxxxxxx'
func (PropertiesTools) ListNames(p properties.Table, prefix string, suffix string) []string {

	dst := make([]string, 0)

	// check params
	if !strings.HasSuffix(prefix, ".") {
		return dst
	}
	if !strings.HasPrefix(suffix, ".") {
		return dst
	}
	if p == nil {
		return dst
	}

	table := p.Export(nil)
	for key := range table {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			name := key[len(prefix) : len(key)-len(suffix)]
			dst = append(dst, name)
		}
	}

	sort.Strings(dst)
	return dst
}
