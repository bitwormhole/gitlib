package utils

import (
	"os"
	"strings"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
)

// ComputeAbsolutePath ... 把 href 解析为完整的绝对路径
func ComputeAbsolutePath(href string, base afs.Path) (afs.Path, error) {

	if base == nil {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		base = files.FS().NewPath(wd)
	}

	href = strings.TrimSpace(href)

	if strings.HasPrefix(href, "/") {
		// as a posix abs path
		abs := base.GetFS().NewPath(href)
		return abs, nil
	} else if isWindowsAbsolutePath(href) {
		// as a windows abs path
		abs := base.GetFS().NewPath(href)
		return abs, nil
	}

	abs := base.GetChild(href)
	return abs, nil
}

func isWindowsAbsolutePath(path string) bool {

	const (
		c1 = "\\"
		c2 = "/"
	)

	path2 := strings.ReplaceAll(path, c1, c2)
	parts := strings.Split(path2, c2)
	p0 := parts[0]
	if len(p0) != 2 {
		return false
	}

	r1 := rune(p0[1])
	if r1 != ':' {
		return false
	}

	r0 := rune(p0[0])
	abc := false
	if 'a' <= r0 && r0 <= 'z' {
		abc = true
	} else if 'a' <= r0 && r0 <= 'z' {
		abc = true
	}
	return abc
}
