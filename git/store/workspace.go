package store

import "bitwormhole.com/starter/afs"

// Workspace 表示工作区目录
type Workspace interface {
	Path() afs.Path
}
